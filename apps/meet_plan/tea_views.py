from django.core.exceptions import PermissionDenied
from django.views.generic import ListView, DetailView
from django.views.generic.edit import CreateView, UpdateView

from utils.mixin.permission import LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin
from .models import MeetPlan, MeetPlanOrder, FeedBack
from .forms import MeetPlanCreateForm, MeetPlanUpdateForm, MeetPlanOrderUpdateForm, FeedBackCreateForm
from .utils import get_semester_date_range


# Create your views here.


class MeetPlanCreateView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, CreateView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_create.html'
    form_class = MeetPlanCreateForm

    def form_valid(self, form):
        form.instance.teacher = self.request.user
        return super().form_valid(form)

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        date_range = get_semester_date_range()
        context['semesterstartdate'] = date_range[0].strftime("%Y-%m-%d")
        context['semesterenddate'] = date_range[1].strftime("%Y-%m-%d")
        return context


class MeetPlanUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, UpdateView):
    model = MeetPlan
    form_class = MeetPlanUpdateForm
    template_name = 'meet_plan/teacher/plan_update.html'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.teacher != self.request.user:
            raise PermissionDenied
        return obj

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        date_range = get_semester_date_range()
        context['semesterstartdate'] = date_range[0].strftime("%Y-%m-%d")
        context['semesterenddate'] = date_range[1].strftime("%Y-%m-%d")
        return context


class MeetPlanListView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, ListView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_all.html'
    paginate_by = 50
    context_object_name = 'meetplan_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False, teacher=self.request.user).order_by('-create_time')


class MeetPlanDetailView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, DetailView):
    model = MeetPlan
    template_name = 'meet_plan/teacher/plan_detail.html'
    context_object_name = 'plan'

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        if obj.teacher != self.request.user:
            raise PermissionDenied
        return obj

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        order_list = MeetPlanOrder.objects.filter(meet_plan=self.get_object(), is_delete=False)
        context['ordernum'] = order_list.count()
        for i in range(0, order_list.count()):
            context['order%did' % i] = order_list[i].id
            context['order%dstuid' % i] = order_list[i].student.identity_id
            context['order%dstu' % i] = order_list[i].student.user_name
            context['order%dfinish' % i] = '已确认' if order_list[i].completed else '待确认'
            context['order%dmessage' % i] = order_list[i].message
        return context


class MeetPlanOrderUpdateView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, UpdateView):
    model = MeetPlanOrder
    template_name = 'meet_plan/teacher/planorder_update.html'
    form_class = MeetPlanOrderUpdateForm

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        self.ori_obj = super().get_object(queryset=queryset)
        if obj.meet_plan.teacher != self.request.user:
            raise PermissionDenied
        return obj

    def form_valid(self, form):
        from .tasks import send_meetplan_order_update_email
        domain = self.request.get_host()
        if self.object.completed != self.ori_obj.completed or self.object.is_delete != self.ori_obj.is_delete:
            send_meetplan_order_update_email.delay(self.object.id, domain)

        response = super().form_valid(form)
        return response


class FeedBackCreateView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, CreateView):
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


class FeedBackListView(LoginRequiredMixin, UserProfileRequiredMixin, TeacherRequiredMixin, ListView):
    model = FeedBack
    template_name = 'meet_plan/teacher/feedback_all.html'
    paginate_by = 10
    context_object_name = 'feedback_list'

    def get_queryset(self):
        qs = super().get_queryset()
        return qs.filter(is_delete=False, teacher=self.request.user).order_by('-create_time')
