from django.core.exceptions import PermissionDenied
from django.db.models import When, Case, BooleanField
from django.shortcuts import get_object_or_404
from django.urls import reverse
from django.utils import timezone
from django.views.generic import ListView
from django.views.generic.edit import CreateView, FormView

from utils.mixin.permission import StuViewMixin
from .forms import MeetPlanOrderCreateForm, StudentAddMeetPlanOrderForm
from .models import MeetPlan, MeetPlanOrder
from .utils import get_term_date
from ..account_auth.models import User


class TeacherListView(StuViewMixin, ListView):
    template_name = 'meet_plan/student/teacher_all.html'
    context_object_name = 'teacher_list'
    paginate_by = 50

    def get_queryset(self):
        return User.objects.filter(is_teacher=True, is_active=True).order_by('identity_id')


class TeacherPlanListView(StuViewMixin, ListView):
    template_name = 'meet_plan/student/teacher_plan_all.html'
    context_object_name = 'plan_list'
    paginate_by = 50
    teacher = ''

    def get_queryset(self):
        date_range = get_term_date()
        self.teacher = get_object_or_404(User, pk=self.kwargs['pk'])
        if not self.teacher.is_teacher:
            raise PermissionDenied('您将要查看的用户身份为学生，这是不合理的，请向管理员反馈！')
        return MeetPlan.objects.filter(teacher=self.teacher,
                                       start_time__gt=date_range[0],
                                       end_time__lt=date_range[1]) \
            .annotate(available=Case(When(start_time__lt=timezone.now(), then=False),
                                     default=True,
                                     output_field=BooleanField())) \
            .order_by('start_time')

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['teacher_id'] = self.teacher.pk
        context['teacher_name'] = self.teacher.user_name
        return context


class MeetPlanOrderCreateView(StuViewMixin, CreateView):
    model = MeetPlanOrder
    template_name = 'meet_plan/student/order_create.html'
    form_class = MeetPlanOrderCreateForm
    meet_plan = ''

    def form_valid(self, form):
        form.instance.meet_plan = self.meet_plan
        form.instance.student = self.request.user
        response = super().form_valid(form)

        from .tasks import send_meetplan_order_create_email
        send_meetplan_order_create_email.delay(self.object.id)

        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.meet_plan.teacher_id])
        cache.delete(key)

        return response

    def post(self, request, *args, **kwargs):
        self.meet_plan = get_object_or_404(MeetPlan, pk=self.kwargs['pk'])
        if not self.meet_plan.available_choice or self.meet_plan.start_time < timezone.now():
            raise PermissionDenied('该安排已满或该安排已过期！')

        if MeetPlanOrder.objects.filter(student=self.request.user,
                                        meet_plan=self.meet_plan).count() > 0:
            raise PermissionDenied('您不能多次选同一个综合指导课安排！')

        return super().post(request, *args, **kwargs)

    def get_success_url(self):
        return reverse('meet_plan:stu-index')


class MeetPlanOrderAddView(StuViewMixin, FormView):
    form_class = StudentAddMeetPlanOrderForm
    template_name = 'meet_plan/student/planorder_add.html'

    def get_success_url(self):
        return reverse('meet_plan:stu-index')

    def form_valid(self, form):
        import dateutil.parser, datetime
        from apps.account_auth.models import User
        data = form.cleaned_data
        end_time = start_time = dateutil.parser.parse('{}T{}+08:00'.format(data['date'], data['time']))
        long = int(data['long'])
        place = data['place']
        message = data['message']
        teacher = data['teacher']
        duration = datetime.timedelta(hours=0.5)
        for i in range(long):
            end_time += duration

        meetplan = MeetPlan.objects.create(teacher=teacher,
                                           place=place,
                                           start_time=start_time,
                                           end_time=end_time,
                                           allow_other=False,
                                           message=message)
        meetplan.save()
        print(meetplan.id)

        meetplan_order = MeetPlanOrder.objects.create(meet_plan=meetplan,
                                                      student=self.request.user,
                                                      completed=False,
                                                      message=message)
        meetplan_order.save()
        from .tasks import send_apply_for_add_meetplan_order_email
        send_apply_for_add_meetplan_order_email.delay(meetplan_order.id)

        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_total_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_avail_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.request.user.id])
        cache.delete(key)

        return super().form_valid(form)
