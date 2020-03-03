from django.views.generic.list import ListView
from django.views.generic.edit import CreateView, UpdateView
from django.views.generic.base import View

from utils.mixin import AdminRequiredMixin
from apps.account_auth.models import User
from .forms import UserCreateForm, UserUpdateForm


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
        return super().form_valid(form)


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
