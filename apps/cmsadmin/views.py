from django.http import HttpResponseRedirect
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic import ListView, CreateView, UpdateView, DeleteView
from django.views.generic.base import View

from utils.mixin.permission import AdminRequiredMixin
from .forms import UpdateRecordForm
from ..account_auth.models import User
from ..meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack
from ..meet_plan.utils import get_term_date
from ..portal.models import UpdateRecord


def noindex(request):
    return HttpResponseRedirect(reverse('cmsadmin:index'))


class IndexView(AdminRequiredMixin, View):
    def get(self, request):
        date_range = get_term_date()
        ctx = {
            'total_teacher_num': User.objects.filter(is_teacher=True).count(),
            'total_student_num': User.objects.filter(is_teacher=False).count(),
            'meetplan_total_num': MeetPlan.objects.count(),
            'meetplan_this_term_total_num': MeetPlan.objects.filter(start_time__gte=date_range[0],
                                                                    end_time__lte=date_range[1]).count(),
            'meetplanorder_this_term_total_num': MeetPlanOrder.objects.filter(meet_plan__start_time__gte=date_range[0],
                                                                              meet_plan__end_time__lte=date_range[
                                                                                  1]).count(),
            'meetplan_feedback_num': FeedBack.objects.count()
        }

        return TemplateResponse(request, template='cmsadmin/index.html', context=ctx)


class UpdateRecordListView(AdminRequiredMixin, ListView):
    model = UpdateRecord
    template_name = 'cmsadmin/site/updaterecord_all.html'
    context_object_name = 'record_list'


class UpdateRecordCreateView(AdminRequiredMixin, CreateView):
    model = UpdateRecord
    template_name = 'cmsadmin/site/updaterecord_create.html'
    form_class = UpdateRecordForm

    def get_success_url(self):
        return reverse('cmsadmin:site-updaterecord-all')


class UpdateRecordUpdateView(AdminRequiredMixin, UpdateView):
    model = UpdateRecord
    template_name = 'cmsadmin/site/updaterecord_update.html'
    form_class = UpdateRecordForm

    def get_success_url(self):
        return reverse('cmsadmin:site-updaterecord-all')


class UpdateRecordDeleteView(AdminRequiredMixin, DeleteView):
    model = UpdateRecord
    template_name = 'cmsadmin/site/updaterecord_delete.html'

    def get_success_url(self):
        return reverse('cmsadmin:site-updaterecord-all')
