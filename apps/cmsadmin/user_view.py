from django.urls import reverse
from django.views.generic.list import ListView
from django.views.generic.edit import CreateView, UpdateView
from django.views.generic.base import View
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer

from django.conf import settings
from utils.mixin import AdminRequiredMixin
from apps.account_auth.models import User
from .forms import UserCreateForm, UserUpdateForm
from .tasks import send_account_active_email

class UserView(AdminRequiredMixin, ListView):
    model = User
    template_name = 'cmsadmin/user/user_all.html'
    paginate_by = 50
    context_object_name = 'user_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False).order_by('-identity_id')


class UserCreateView(AdminRequiredMixin, CreateView):
    model = User
    template_name = 'cmsadmin/user/user_create.html'
    form_class = UserCreateForm
    success_url = '/cmsadmin/user_all/'

    def form_valid(self, form):
        # TODO: 如果添加成功, 发送邮件通知
        # return super().form_valid(form)
        response = super().form_valid(form)
        # 加密用户信息
        serializer = Serializer(settings.SECRET_KEY, expires_in=60*60*24*7)
        info = {'active': form.cleaned_data['identity_id']}
        token = serializer.dumps(info)  # bytes
        token = token.decode()
        email = '{}@pku.edu.cn'.format(form.cleaned_data['identity_id'])
        email = '598049186@qq.com'
        user_name = form.cleaned_data['user_name']
        domain = self.request.META['HTTP_HOST']
        active_path = reverse('account_auth:active-account', kwargs={'token': token})
        active_url = 'http://{}{}'.format(domain, active_path)
        # 发邮件
        send_account_active_email.delay(email, user_name, active_url, domain)
        return response


class CreateManyUserView(AdminRequiredMixin, View):
    model = User
    # TODO: 增加批量添加用户功能


class UpdateUserView(AdminRequiredMixin, UpdateView):
    model = User
    form_class = UserUpdateForm
    template_name = 'cmsadmin/user/user_update.html'
    success_url = '/cmsadmin/user_all/'


class DeletedUserListView(AdminRequiredMixin, ListView):
    model = User
    template_name = 'cmsadmin/user/user_deletelist.html'
    paginate_by = 50
    context_object_name = 'user_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=True).order_by('-update_time')
