from django.urls import reverse
from django.views.generic.base import View
from django.views.generic.list import ListView
from django.views.generic.edit import CreateView, UpdateView, DeleteView, FormView

from utils.mixin.permission import AdminRequiredMixin
from apps.meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack
from .forms import MeetPlanForm, MeetPlanOrderForm, FeedBackForm, OptionForm
from apps.meet_plan.utils import get_term_date


class MeetPlanListView(AdminRequiredMixin, ListView):
    model = MeetPlan
    template_name = 'cmsadmin/meetplan/meetplan_all.html'
    paginate_by = 50
    context_object_name = 'meetplan_list'

    def get_queryset(self):
        return super().get_queryset().order_by('-create_time')

    def get_context_data(self, **kwargs):
        ctx = super().get_context_data(**kwargs)
        date_range = get_term_date()
        ctx['term_start_date'] = date_range[0]
        ctx['term_end_date'] = date_range[1]
        return ctx


class MeetPlanCreateView(AdminRequiredMixin, CreateView):
    model = MeetPlan
    template_name = 'cmsadmin/meetplan/meetplan_create.html'
    form_class = MeetPlanForm

    def get_success_url(self):
        return reverse('cmsadmin:meetplan_all')


class MeetPlanUpdateView(AdminRequiredMixin, UpdateView):
    model = MeetPlan
    form_class = MeetPlanForm
    template_name = 'cmsadmin/meetplan/meetplan_update.html'

    def get_success_url(self):
        return reverse('cmsadmin:meetplan_all')


class MeetPlanDeleteView(AdminRequiredMixin, DeleteView):
    model = MeetPlan
    template_name = 'cmsadmin/meetplan/meetplan_confirm_delete.html'

    def get_success_url(self):
        return reverse('cmsadmin:meetplan_all')


class MeetPlanOrderListView(AdminRequiredMixin, ListView):
    model = MeetPlanOrder
    template_name = 'cmsadmin/meetplan/meetplanorder_all.html'
    paginate_by = 50
    context_object_name = 'meetplanorder_list'

    def get_queryset(self):
        return super().get_queryset().order_by('-create_time')


class MeetPlanOrderCreateView(AdminRequiredMixin, CreateView):
    model = MeetPlanOrder
    template_name = 'cmsadmin/meetplan/meetplanorder_create.html'
    form_class = MeetPlanOrderForm

    def get_success_url(self):
        return reverse('cmsadmin:meetplanorder_all')


class MeetPlanOrderViewUpdate(AdminRequiredMixin, UpdateView):
    model = MeetPlanOrder
    form_class = MeetPlanOrderForm
    template_name = 'cmsadmin/meetplan/meetplanorder_update.html'

    def get_success_url(self):
        return reverse('cmsadmin:meetplanorder_all')


class MeetPlanOrderDeleteView(AdminRequiredMixin, DeleteView):
    model = MeetPlanOrder
    template_name = 'cmsadmin/meetplan/meetplanorder_confirm_delete.html'

    def get_success_url(self):
        return reverse('cmsadmin:meetplanorder_all')


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
    form_class = FeedBackForm
    template_name = 'cmsadmin/meetplan/meetplan_feedback_update.html'

    def get_success_url(self):
        return reverse('cmsadmin:feedback_all')

    def get_object(self, queryset=None):
        obj = super().get_object(queryset=queryset)
        self.ori_obj = super().get_object(queryset=queryset)
        return obj

    def form_valid(self, form):
        from apps.meet_plan.tasks import send_meetplan_feedback_update_email
        domain = self.request.get_host()
        if self.object.have_checked != self.ori_obj.have_checked:
            send_meetplan_feedback_update_email.delay(self.object.id, domain)

        response = super().form_valid(form)
        return response


class TermDateUpdateView(AdminRequiredMixin, FormView):
    template_name = 'cmsadmin/meetplan/term_date_update.html'
    form_class = OptionForm

    def get_success_url(self):
        return reverse('cmsadmin:meetplan_all')

    def form_valid(self, form):
        from apps.options.models import Option
        st = Option.objects.get(app='meet_plan', name='term_start_date')
        ed = Option.objects.get(app='meet_plan', name='term_end_date')
        print(type(form.cleaned_data['start']), form.cleaned_data['end'])
        st.value = '{}{}'.format(form.cleaned_data['start'], 'T00:00:00+08:00')
        ed.value = '{}{}'.format(form.cleaned_data['end'], 'T00:00:00+08:00')
        st.save()
        ed.save()
        return super().form_valid(form)
