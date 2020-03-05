from django.core.exceptions import PermissionDenied
from django.http import HttpResponseRedirect
from django.template.response import TemplateResponse
from django.urls import reverse
from django.contrib.auth import logout
from django.views.generic.base import View
from django.views.generic.edit import CreateView, UpdateView
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer, SignatureExpired, BadData
from django.conf import settings

from utils.mixin.permission import LoginRequiredMixin, UserProfileRequiredMixin
from utils.mixin.view import ImgUploadViewMixin

from .models import User, UserProfile
from .forms import UserEmailUpdateForm, UserProfileUpdateForm, UserProfileCreateForm


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
            return TemplateResponse(request, template='account_auth/login/active.html')
        except SignatureExpired:
            ctx = {
                'error_message': '激活链接已过期！ 请登录或联系管理员获取新的激活链接。'
            }
            return TemplateResponse(request, template='404.html', context=ctx)
        except BadData:
            ctx = {
                'error_message': '激活链接错误！ 请复制粘贴完整的激活链接。'
            }
            return TemplateResponse(request, template='404.html', context=ctx)


class UserProfileAddView(LoginRequiredMixin, CreateView):
    # model = UserProfile
    form_class = UserProfileCreateForm
    template_name = 'account_auth/userprofile_create.html'
    success_url = '/index/'

    def form_valid(self, form):
        form.instance.user = self.request.user
        return super().form_valid(form)


class UserProfileUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, UpdateView):
    model = UserProfile
    template_name = 'account_auth/userprofile_update.html'
    form_class = UserProfileUpdateForm
    # fields = ['gender', 'telephone', 'birth', 'user_img']
    success_url = '/index/'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.user != self.request.user:
            raise PermissionDenied
        return obj


class UserEmailUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, UpdateView):
    model = User
    form_class = UserEmailUpdateForm
    template_name = 'account_auth/useremail_update.html'
    success_url = '/index/'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj != self.request.user:
            raise PermissionDenied
        return obj


class UserProfileImgUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, ImgUploadViewMixin):
    template_name = 'account_auth/userprofileimg_upload.html'

    def get_success_url(self):
        return reverse('portal:index')

    def form_valid(self, form):
        response = super().form_valid(form)
        self.request.user.userprofile.head_picture = self.object
        self.request.user.userprofile.save()
        return response
