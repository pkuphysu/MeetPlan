from django.http import HttpResponseRedirect
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic import ListView
from django.views.generic.base import View

from utils.mixin.permission import ViewMixin
from .models import FriendLink
from ..account_auth.models import User


# Create your views here.


# http://hostname 或 http://hostname/ 重定向至 http://hostname/index/
def noindex(request):
    return HttpResponseRedirect(reverse('portal:index'))


class IndexView(ViewMixin, View):
    def get(self, request):
        return TemplateResponse(request, 'index.html')


class ContactView(ViewMixin, ListView):
    model = User
    template_name = 'portal/contact.html'
    context_object_name = 'admin_list'

    def get_queryset(self):
        return super().get_queryset().filter(is_admin=True)


class FriendLinkView(ViewMixin, ListView):
    model = FriendLink
    template_name = 'portal/friendlink.html'
    context_object_name = 'friendlink_list'
