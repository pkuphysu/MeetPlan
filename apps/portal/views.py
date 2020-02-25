from django.shortcuts import render
from django.http import HttpResponseRedirect, HttpResponse
from django.urls import reverse
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic.base import View

from utils.mixin import LoginRequiredMixin, UserProfileRequiredMixin


# Create your views here.


# http://hostname 或 http://hostname/ 重定向至 http://hostname/index/
def noindex(request):
    if not request.user.is_authenticated:
        if request.GET.get('token'):
            return HttpResponseRedirect(reverse('account_auth:iaaa_auth')
                                        + '?rand=%s&token=%s' % (request.GET.get('rand'),
                                                                 request.GET.get('token')))
        else:
            return HttpResponseRedirect(reverse('account_auth:iaaa_login'))
    else:
        return HttpResponseRedirect(reverse('portal:index'))


class IndexView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'index.html')


class AboutView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'portal/about.html')


class ContactView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'portal/contact.html')


class RecruitmentView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'portal/recruitment.html')


class FriendLinkView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'portal/friendlink.html')
