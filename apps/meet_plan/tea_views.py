from django.core.exceptions import PermissionDenied
from django.urls import reverse
from django.views.generic import ListView, DetailView
from django.views.generic.edit import CreateView, UpdateView, DeleteView, FormView

from utils.mixin.permission import TeaViewMixin
from .models import MeetPlan, MeetPlanOrder, FeedBack
from .forms import MeetPlanForm, MeetPlanOrderUpdateForm, FeedBackCreateForm, MeetPlanFastCreateForm
from .utils import get_term_date


# Create your views here.


class MeetPlanCreateView(TeaViewMixin, CreateView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_create.html'
    form_class = MeetPlanForm

    def get_success_url(self):
        return reverse('meet_plan:tea-plan-detail', kwargs={'pk': self.object.pk})

    def form_valid(self, form):

        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_total_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_avail_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.request.user.id])
        cache.delete(key)

        form.instance.teacher = self.request.user
        return super().form_valid(form)

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        date_range = get_term_date()
        context['term_start_date'] = date_range[0].strftime("%Y-%m-%d")
        context['term_end_date'] = date_range[1].strftime("%Y-%m-%d")
        return context


class MeetPlanFastCreateView(TeaViewMixin, FormView):
    template_name = 'meet_plan/teacher/plan_create_fast.html'
    form_class = MeetPlanFastCreateForm

    def get_success_url(self):
        return reverse('meet_plan:tea-index')

    def form_valid(self, form):
        import dateutil.parser, datetime

        data = form.cleaned_data
        start_time = dateutil.parser.parse('{}T{}+08:00'.format(data['date'], data['time']))
        long = int(data['long'])
        place = data['place']
        allow_other = True if data['allow_other'] == '1' else False
        message = data['message']
        term_range = get_term_date()[1] if data['every_week'] == '2' else start_time

        meetplan_obj_list = []
        duration = datetime.timedelta(hours=0.5)
        duration_week = datetime.timedelta(weeks=1)
        while start_time <= term_range:
            s_time = start_time
            for i in range(long):
                meetplan_obj_list.append(
                    MeetPlan(
                        teacher=self.request.user,
                        place=place,
                        start_time=s_time,
                        end_time=s_time+duration,
                        allow_other=allow_other,
                        message=message,
                        available_choice=2 if allow_other else 1
                    )
                )
                s_time += duration
            start_time += duration_week
        MeetPlan.objects.bulk_create(meetplan_obj_list)

        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_total_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_avail_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.request.user.id])
        cache.delete(key)

        return super().form_valid(form)

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        date_range = get_term_date()
        context['term_start_date'] = date_range[0].strftime("%Y-%m-%d")
        context['term_end_date'] = date_range[1].strftime("%Y-%m-%d")
        return context


class MeetPlanUpdateView(TeaViewMixin, UpdateView):
    model = MeetPlan
    form_class = MeetPlanForm
    template_name = 'meet_plan/teacher/plan_update.html'

    def get_success_url(self):
        return reverse('meet_plan:tea-plan-detail', kwargs={'pk': self.object.pk})

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.teacher != self.request.user:
            raise PermissionDenied('您只能更改您创建的综合指导课安排！')
        return obj

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        date_range = get_term_date()
        context['semesterstartdate'] = date_range[0].strftime("%Y-%m-%d")
        context['semesterenddate'] = date_range[1].strftime("%Y-%m-%d")
        return context

    def form_valid(self, form):
        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_plan_detail', [self.object.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_total_num', [self.object.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_avail_num', [self.object.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.object.id])
        cache.delete(key)

        return super().form_valid(form)


class MeetPlanListView(TeaViewMixin, ListView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_all.html'
    paginate_by = 50
    context_object_name = 'meetplan_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(teacher=self.request.user).order_by('-create_time')


class MeetPlanDetailView(TeaViewMixin, DetailView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_detail.html'
    context_object_name = 'plan'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.teacher != self.request.user:
            raise PermissionDenied('您只能查看您创建的综合指导课安排！')
        return obj


class MeetPlanDeleteView(TeaViewMixin, DeleteView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_confirm_delete.html'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.teacher != self.request.user:
            raise PermissionDenied('您只能删除您创建的综合指导课安排！')
        return obj

    def get_success_url(self):
        return reverse('meet_plan:tea-index')

    def delete(self, request, *args, **kwargs):
        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_total_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_avail_num', [self.request.user.id])
        cache.delete(key)
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.request.user.id])
        cache.delete(key)

        return super().delete(request, *args, **kwargs)


class MeetPlanOrderUpdateView(TeaViewMixin, UpdateView):
    model = MeetPlanOrder
    template_name = 'meet_plan/teacher/planorder_update.html'
    form_class = MeetPlanOrderUpdateForm

    def get_success_url(self):
        return reverse('meet_plan:tea-index')

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.meet_plan.teacher != self.request.user:
            raise PermissionDenied('您只能修改属于您创建的综合指导课的预约！')
        return obj

    def form_valid(self, form):
        from .tasks import send_meetplan_order_update_email
        domain = self.request.get_host()
        if form.has_changed():
            send_meetplan_order_update_email.delay(self.object.id, domain, False)

        response = super().form_valid(form)
        return response


class MeetPlanOrderDeleteView(TeaViewMixin, DeleteView):
    model = MeetPlanOrder
    template_name = 'meet_plan/teacher/planorder_confirm_delete.html'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.meet_plan.teacher != self.request.user:
            raise PermissionDenied('您只能删除您创建的综合指导课的预约！')
        return obj

    def delete(self, request, *args, **kwargs):
        from django.core.cache import cache
        from django.core.cache.utils import make_template_fragment_key
        key = make_template_fragment_key('meetplan_meetplan_order_avail_num', [self.request.user.id])
        cache.delete(key)

        response = super().delete(request, args, kwargs)
        from .tasks import send_meetplan_order_update_email
        domain = self.request.get_host()
        send_meetplan_order_update_email.delay(self.object.id, domain, True)
        return response

    def get_success_url(self):
        return reverse('meet_plan:tea-index')


class FeedBackCreateView(TeaViewMixin, CreateView):
    model = FeedBack
    template_name = 'meet_plan/teacher/feedback_create.html'
    form_class = FeedBackCreateForm

    def form_valid(self, form):
        form.instance.teacher = self.request.user
        response = super().form_valid(form)

        from .tasks import send_meetplan_feedback_create_email
        domain = self.request.get_host()
        send_meetplan_feedback_create_email.delay(self.object.id, domain)

        return response

    def get_success_url(self):
        return reverse('meet_plan:tea-feedback-list')


class FeedBackListView(TeaViewMixin, ListView):
    model = FeedBack
    template_name = 'meet_plan/teacher/feedback_all.html'
    paginate_by = 10
    context_object_name = 'feedback_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(teacher=self.request.user).order_by('-create_time')
