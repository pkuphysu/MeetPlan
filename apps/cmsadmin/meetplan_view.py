from django.http import HttpResponseRedirect
from django.shortcuts import render

from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic.list import ListView
from django.views.generic.edit import CreateView, UpdateView
from django.views.generic.base import View

from utils.mixin import AdminRequiredMixin
from apps.meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack
from .forms import MeetPlanCreateForm, MeetPlanUpdateForm, MeetPlanOrderCreateForm, MeetPlanOrderUpdateForm, FeedBackUpdateForm


class MeetPlanView(AdminRequiredMixin, ListView):
    model = MeetPlan
    template_name = 'cmsadmin/meetplan/meetplan_all.html'
    paginate_by = 50
    context_object_name = 'meetplan_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False).order_by('-create_time')


class MeetPlanCreateView(AdminRequiredMixin, CreateView):
    model = MeetPlan
    template_name = 'cmsadmin/meetplan/meetplan_create.html'
    form_class = MeetPlanCreateForm
    success_url = '/cmsadmin/meetplan_all/'


class MeetPlanUpdateView(AdminRequiredMixin, UpdateView):
    model = MeetPlan
    form_class = MeetPlanUpdateForm
    template_name = 'cmsadmin/meetplan/meetplan_update.html'
    success_url = '/cmsadmin/meetplan_all/'


class MeetPlanOrderView(AdminRequiredMixin, ListView):
    model = MeetPlanOrder
    template_name = 'cmsadmin/meetplan/meetplanorder_all.html'
    paginate_by = 50
    context_object_name = 'meetplanorder_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.order_by('-create_time')


class MeetPlanOrderCreateView(AdminRequiredMixin, CreateView):
    model = MeetPlanOrder
    template_name = 'cmsadmin/meetplan/meetplanorder_create.html'
    form_class = MeetPlanOrderCreateForm
    success_url = '/cmsadmin/meetplanorder_all/'


class UpdateMeetPlanOrderView(AdminRequiredMixin, UpdateView):
    model = MeetPlanOrder
    form_class = MeetPlanOrderUpdateForm
    template_name = 'cmsadmin/meetplan/meetplanorder_update.html'
    success_url = '/cmsadmin/meetplanorder_all/'


class FeedBackListView(AdminRequiredMixin, ListView):
    model = FeedBack
    template_name = 'cmsadmin/meetplan/meetplan_feedback_all.html'
    paginate_by = 20
    context_object_name = 'feedback_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False, teacher=self.request.user).order_by('-create_time')


class FeedBackUpdateView(AdminRequiredMixin, UpdateView):
    model = FeedBack
    form_class = FeedBackUpdateForm
    template_name = 'cmsadmin/meetplan/meetplan_feedback_update.html'
    success_url = '/cmsadmin/meetplanfeedback_all/'
