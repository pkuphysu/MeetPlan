from django.db.models import Count, Q
from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic import View

from apps.meet_plan.utils import get_semester_date_range
from utils.mixin import LoginRequiredMixin, UserProfileRequiredMixin
# Create your views here.
from apps.account_auth.models import User, UserProfile
from .models import MeetPlan, MeetPlanOrder


# /meetplan 或 /meetplan/ 重定向至 /meetplan/index/
def noindex(request):
    return HttpResponseRedirect(reverse('meet_plan:index'))


class IndexView(LoginRequiredMixin, UserProfileRequiredMixin, View):
    date_range = get_semester_date_range()

    def get(self, request):
        current_user = request.user
        if current_user.is_teacher:
            # TODO: 完善教师综合指导课概览页面
            queryset = MeetPlan.objects.filter(teacher=current_user, is_delete=False)
            context = {
                'semesterstartdate': self.date_range[0].strftime("%Y-%m-%d"),
                'semesterenddate': self.date_range[1].strftime("%Y-%m-%d"),
                'this_semester_plan': queryset.filter(start_time__gt=self.date_range[0],
                                                      end_time__lt=self.date_range[1]).order_by('start_time'),
                'history_plan_num': queryset.count(),
                'this_semeseter_planorder_num': queryset.aggregate(num=Count('meetplanorder',
                                                                             filter=Q(meetplanorder__is_delete=False)))[
                    'num'],
                'this_semester_planorder': queryset.filter(start_time__gt=self.date_range[0],
                                                           end_time__lt=self.date_range[1],
                                                           meetplanorder__isnull=False).distinct().order_by(
                    '-start_time')

            }
            return TemplateResponse(request, 'meet_plan/teacher/index.html', context)
        else:
            # TODO: 完善学生综合指导课概览页面
            queryset = MeetPlanOrder.objects.filter(student=current_user, is_delete=False)
            context = {
                'semesterstartdate': self.date_range[0].strftime("%Y-%m-%d"),
                'semesterenddate': self.date_range[1].strftime("%Y-%m-%d"),
                'this_meetplanorder_list': queryset.filter(meet_plan__start_time__gt=self.date_range[0],
                                                           meet_plan__end_time__lt=self.date_range[1]).order_by(
                    'meet_plan__start_time'),
                'meetplanorder_list': queryset.order_by('create_time'),
                'meetplanorder_ava_num':queryset.filter(completed=True).count()
            }

            return TemplateResponse(request, 'meet_plan/student/index.html', context=context)
