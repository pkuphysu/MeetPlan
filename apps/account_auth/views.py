from django.core.exceptions import PermissionDenied
from django.http import HttpResponseRedirect
from django.shortcuts import get_object_or_404
from django.template.response import TemplateResponse
from django.urls import reverse
from django.contrib.auth import logout
from django.views.generic.base import View
from django.views.generic.edit import CreateView, UpdateView
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer, SignatureExpired, BadData
from django.conf import settings

from utils.mixin.permission import TeaViewMixin, StuViewMixin, ViewMixin, LoginRequiredMixin, TeacherRequiredMixin
from utils.mixin.view import ImgUploadViewMixin

from .models import User, BaseProfile, StudentProfile, Major, TeacherProfile
from .forms import UserEmailForm, UserProfileForm, StudentProfileForm, TeacherProfileForm
from . import urls

# Create your views here.


def logout_view(request):
    """退出登录"""
    logout(request)
    return HttpResponseRedirect(reverse('account_auth:iaaa_login'))


# /user 或 /user/ 重定向至 /user/index/
def noindex(request):
    return HttpResponseRedirect(reverse('user:index'))


def index(request):
    return HttpResponseRedirect(reverse('portal:index'))


class ActiveView(View):
    """用户激活"""

    def get(self, request, token):
        """进行用户激活"""
        # 进行解密，获取要激活的用户信息
        serializer = Serializer(settings.SECRET_KEY, 60 * 60 * 24 * 7)
        try:
            info = serializer.loads(token)
            # 获取待激活用户的id
            user_id = info['active']

            # 根据id获取用户信息
            user = User.objects.get(identity_id=user_id)
            user.is_active = True
            user.save()
            # 跳转到登录页面
            ctx = {
                'app_id': settings.APPID,
                'redirect_url': settings.APPREDIRECTURL
            }
            return TemplateResponse(request, template='account_auth/login/active.html', context=ctx)
        except SignatureExpired:
            raise PermissionDenied('激活链接已过期！ 请登录或联系管理员获取新的激活链接')
        except BadData:
            raise PermissionDenied('激活链接错误！ 请复制粘贴完整的激活链接')


class UserEmailUpdateView(ViewMixin, UpdateView):

    model = User
    form_class = UserEmailForm
    template_name = 'account_auth/user_email_update.html'

    def get_success_url(self):
        return reverse('portal:index')

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj != self.request.user:
            raise PermissionDenied('你只能更改自己的邮件！')
        return obj


class BaseProfileAddView(LoginRequiredMixin, CreateView):
    model = BaseProfile
    form_class = UserProfileForm
    template_name = 'account_auth/base_profile_create.html'

    def get_success_url(self):
        if self.request.user.is_teacher:
            return reverse('account_auth:teacher-profile-create')
        else:
            return reverse('account_auth:student-profile-create')

    def form_valid(self, form):
        form.instance.user = self.request.user
        return super().form_valid(form)

    def get(self, request, *args, **kwargs):
        if not hasattr(request.user, 'baseprofile'):
            return super().get(request, *args, **kwargs)
        else:
            if request.user.is_teacher:
                return HttpResponseRedirect(reverse('account_auth:teacher-profile-create'))
            else:
                return HttpResponseRedirect(reverse('account_auth:student-profile-create'))


class BaseProfileImgUpdateView(ViewMixin, ImgUploadViewMixin):
    template_name = 'account_auth/base_profile_img_upload.html'

    def get_success_url(self):
        return reverse('portal:index')

    def form_valid(self, form):
        form.instance.app = urls.app_name
        response = super().form_valid(form)
        self.request.user.baseprofile.head_picture = self.object
        self.request.user.baseprofile.save()
        return response


class StudentProfileCreateView(ViewMixin, CreateView):
    model = StudentProfile
    template_name = 'account_auth/student_profile_create.html'
    form_class = StudentProfileForm

    def get_success_url(self):
        return reverse('portal:index')

    def get(self, request, *args, **kwargs):
        try:
            id = request.user.studentprofile.id
            return HttpResponseRedirect(reverse('account_auth:student-profile-update',
                                                kwargs={'pk': id}))
        except StudentProfile.DoesNotExist:
            return super().get(request, *args, **kwargs)

    def post(self, request, *args, **kwargs):
        if request.user.is_teacher:
            raise PermissionDenied('您的身份是老师，请转向老师补充资料界面。')
        else:
            return super().post(request, *args, **kwargs)

    def form_valid(self, form):
        form.instance.user = self.request.user
        return super().form_valid(form)


class StudentProfileUpdateView(StuViewMixin, UpdateView):
    model = StudentProfile

    form_class = StudentProfileForm
    template_name = 'account_auth/student_profile_update.html'

    def get_success_url(self):
        return reverse('portal:index')

    def get_object(self, queryset=None):
        obj = get_object_or_404(StudentProfile, id=self.kwargs['pk'])
        if obj != self.request.user.studentprofile:
            raise PermissionDenied('你只能更改自己的个人资料！')
        return obj


class LoadMajorView(ViewMixin, View):
    def get(self, request):
        if request.is_ajax:
            department_id = request.GET.get('department')
            majors = Major.objects.filter(department_id=department_id)
            return TemplateResponse(request, 'account_auth/ajax/major_dropdown_list_options.html', {'majors': majors})
        else:
            raise PermissionDenied('本接口只允许ajax请求')


class TeacherProfileCreateView(ViewMixin, TeacherRequiredMixin, CreateView):
    model = TeacherProfile
    template_name = 'account_auth/teacher_profile_create.html'
    form_class = TeacherProfileForm

    def get_success_url(self):
        return reverse('portal:index')

    def get(self, request, *args, **kwargs):
        try:
            id = request.user.teacherprofile.id
            return HttpResponseRedirect(reverse('account_auth:teacher-profile-update',
                                            kwargs={'pk': id}))
        except TeacherProfile.DoesNotExist:
            return super().get(request, *args, **kwargs)

    def post(self, request, *args, **kwargs):
        if not request.user.is_teacher:
            raise PermissionDenied('您的身份是学生，请转向学生补充资料界面。')
        else:
            return super().post(request, *args, **kwargs)

    def form_valid(self, form):
        form.instance.user = self.request.user
        return super().form_valid(form)


class TeacherProfileUpdateView(TeaViewMixin, UpdateView):
    model = TeacherProfile
    form_class = TeacherProfileForm
    template_name = 'account_auth/teacher_profile_update.html'

    def get_success_url(self):
        return reverse('portal:index')

    def get_object(self, queryset=None):
        obj = get_object_or_404(TeacherProfile, id=self.kwargs['pk'])
        if obj != self.request.user.teacherprofile:
            raise PermissionDenied('你只能更改自己的个人资料！')
        return obj
