from django.core.exceptions import PermissionDenied
from django.shortcuts import get_object_or_404

from django.views.generic import ListView
from django.db.models import Count, Sum, When, Case, BooleanField
from django.db.models import Q
from django.utils import timezone
from django.views.generic.edit import CreateView

from utils.mixin import LoginRequiredMixin, UserProfileRequiredMixin
from apps.account_auth.models import User
from .models import MeetPlan, MeetPlanOrder
from .forms import MeetPlanOrderCreateForm
from .utils import get_semester_date_range


class TeacherListView(LoginRequiredMixin, UserProfileRequiredMixin, ListView):
    date_range = get_semester_date_range()

    queryset = User.objects.filter(is_delete=False, is_teacher=True) \
        .annotate(meetplan_num=Count('meetplan',
                                     filter=Q(meetplan__start_time__gt=date_range[0],
                                              meetplan__end_time__lt=date_range[1],
                                              meetplan__is_delete=False),
                                     distinct=True
                                     ),
                  meetplan_available_num=Count('meetplan',
                                               filter=Q(meetplan__start_time__gt=timezone.now(),
                                                        meetplan__end_time__lt=date_range[1],
                                                        meetplan__available_choice__gt=0,
                                                        meetplan__is_delete=False),
                                               distinct=True
                                               ),
                  meetplanorder_available_num=Sum('meetplan__available_choice',
                                                  filter=Q(meetplan__start_time__gt=timezone.now(),
                                                           meetplan__end_time__lt=date_range[1],
                                                           meetplan__is_delete=False)
                                                  )
                  ) \
        .order_by('user_name')
    template_name = 'meet_plan/student/teacher_all.html'
    context_object_name = 'teacher_list'


class TeacherPlanListView(LoginRequiredMixin, UserProfileRequiredMixin, ListView):
    date_range = get_semester_date_range()
    template_name = 'meet_plan/student/teacher_plan_all.html'
    context_object_name = 'plan_list'
    paginate_by = 50
    teacher = ''

    def get_queryset(self):
        self.teacher = get_object_or_404(User, identity_id=self.kwargs['tea_id'])
        if not self.teacher.is_teacher:
            raise PermissionDenied
        return MeetPlan.objects.filter(teacher=self.teacher,
                                       start_time__gt=self.date_range[0],
                                       end_time__lt=self.date_range[1],
                                       is_delete=False) \
            .annotate(available=Case(When(start_time__lt=timezone.now(), then=False),
                                     default=True,
                                     output_field=BooleanField())) \
            .order_by('start_time')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['teacher_id'] = self.teacher.identity_id
        context['teacher_name'] = self.teacher.user_name
        return context


class MeetPlanOrderCreateView(LoginRequiredMixin, UserProfileRequiredMixin, CreateView):
    date_range = get_semester_date_range()
    model = MeetPlanOrder
    template_name = 'meet_plan/student/order_create.html'
    form_class = MeetPlanOrderCreateForm
    meet_plan = ''

    def form_valid(self, form):
        form.instance.meet_plan = self.meet_plan
        form.instance.student = self.request.user
        # TODO: 发送邮件
        return super().form_valid(form)

    # def is_limit_reached(self):
    #     return Contact.objects.filter(contact_owner=self.request.user).count() >= 100

    def post(self, request, *args, **kwargs):
        self.meet_plan = get_object_or_404(MeetPlan, id=self.kwargs['mp_id'])
        if MeetPlanOrder.objects.filter(student=self.request.user,
                                        meet_plan__start_time__gt=self.date_range[0],
                                        meet_plan__end_time__lt=self.date_range[1],
                                        is_delete=False
                                        ).count() >= 2:
            raise PermissionDenied
        if not self.meet_plan.available_choice or self.meet_plan.start_time < timezone.now():
            raise PermissionDenied
        else:
            return super().post(request, *args, **kwargs)
