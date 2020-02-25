from django.http import HttpResponseRedirect
from django.shortcuts import render

from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic.base import View

from utils.mixin import AdminRequiredMixin, LoginRequiredMixin


def noindex(request):
    return HttpResponseRedirect(reverse('cmsadmin:index'))


class IndexView(AdminRequiredMixin, View):
    @staticmethod
    def get(request):
        return TemplateResponse(request, template='cmsadmin/index.html')

