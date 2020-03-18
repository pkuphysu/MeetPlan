from django.db.models import Count
from django.http import HttpResponseRedirect
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic import View
from django.utils import timezone

from apps.meet_plan.utils import get_term_date
from utils.mixin.permission import ViewMixin, StuViewMixin, TeaViewMixin
# Create your views here.
from .models import MeetPlan, MeetPlanOrder


# /meetplan 或 /meetplan/ 重定向至 /meetplan/index/
class NoIndexView(ViewMixin, View):
    def get(self, request):
        if request.user.is_teacher:
            return HttpResponseRedirect(reverse('meet_plan:tea-index'))
        else:
            return HttpResponseRedirect(reverse('meet_plan:stu-index'))


class TeaIndexView(TeaViewMixin, View):
    def get(self, request):
        date_range = get_term_date()
        current_user = request.user
        queryset = MeetPlan.objects.filter(teacher=current_user)
        context = {
            'term_start_date': date_range[0].strftime("%Y-%m-%d"),
            'term_end_date': date_range[1].strftime("%Y-%m-%d"),
            'this_term_plan': queryset.filter(start_time__gt=date_range[0],
                                              end_time__lt=date_range[1]).order_by('start_time'),
            'history_plan_num': queryset.count(),
            'this_term_planorder_num': queryset.aggregate(num=Count('meetplanorder'))['num'],
            'this_term_plan_before_now': queryset.filter(start_time__gt=date_range[0],
                                                         end_time__lt=timezone.now()).order_by('-start_time'),
            'this_term_plan_after_now': queryset.filter(start_time__gte=timezone.now(),
                                                        end_time__lte=date_range[1]).order_by('start_time')

        }
        return TemplateResponse(request, 'meet_plan/teacher/index.html', context)


class StuIndexView(StuViewMixin, View):
    def get(self, request):
        date_range = get_term_date()
        current_user = request.user
        queryset = MeetPlanOrder.objects.filter(student=current_user)
        context = {
            'term_start_date': date_range[0].strftime("%Y-%m-%d"),
            'term_end_date': date_range[1].strftime("%Y-%m-%d"),
            'this_meetplanorder_list': queryset.filter(meet_plan__start_time__gt=date_range[0],
                                                       meet_plan__end_time__lt=date_range[1]).order_by(
                'meet_plan__start_time'),
            'meetplanorder_list': queryset.order_by('create_time'),
            'meetplanorder_ava_num': queryset.filter(completed=True).count()
        }

        return TemplateResponse(request, 'meet_plan/student/index.html', context=context)
