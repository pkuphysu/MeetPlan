from django.http import HttpResponseRedirect

from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic.base import View

from utils.mixin.permission import AdminRequiredMixin
from ..account_auth.models import User
from ..meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack
from ..meet_plan.utils import get_term_date


def noindex(request):
    return HttpResponseRedirect(reverse('cmsadmin:index'))


class IndexView(AdminRequiredMixin, View):
    def get(self, request):
        date_range = get_term_date()
        ctx = {
            'total_teacher_num': User.objects.filter(is_teacher=True).count(),
            'total_student_num': User.objects.filter(is_teacher=False).count(),
            'term_start_date': date_range[0].strftime("%Y-%m-%d"),
            'term_end_date': date_range[1].strftime("%Y-%m-%d"),
            'meetplan_total_num': MeetPlan.objects.count(),
            'meetplan_this_term_total_num': MeetPlan.objects.filter(start_time__gte=date_range[0],
                                                                    end_time__lte=date_range[1]).count(),
            'meetplanorder_this_term_total_num': MeetPlanOrder.objects.filter(meet_plan__start_time__gte=date_range[0],
                                                                              meet_plan__end_time__lte=date_range[1]).count(),
            'meetplan_feedback_num': FeedBack.objects.count()
        }

        return TemplateResponse(request, template='cmsadmin/index.html', context=ctx)

