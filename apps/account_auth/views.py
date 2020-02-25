from django.core.exceptions import PermissionDenied
from django.shortcuts import render
from django.http import HttpResponseRedirect, HttpResponse
from django.urls import reverse
from django.contrib.auth import logout
from django.views.generic.edit import CreateView, UpdateView

from utils.mixin import LoginRequiredMixin, UserProfileRequiredMixin
from .models import User, UserProfile
from .forms import UserEmailUpdateForm, UserProfileUpdateForm, UserProfileCreateForm, UserProfileImgUpdateForm


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


class UserProfileImgUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, UpdateView):
    model = UserProfile
    template_name = 'account_auth/userprofileimg_update.html'
    form_class = UserProfileImgUpdateForm
    success_url = '/index/'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.user != self.request.user:
            raise PermissionDenied
        return obj
