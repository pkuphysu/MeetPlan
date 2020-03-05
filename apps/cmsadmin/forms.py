from django import forms
from utils.mixin.form import FormMixin
from apps.account_auth.models import User
from apps.meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack, SemesterDateRange


class UserCreateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = User
        fields = {
            'identity_id',
            'user_name',
            'email',
            'is_teacher',
            'is_admin'
        }
        labels = {
            'identity_id': '职工号\\学号',
            'user_name': '姓名',
            'email': '电子邮件',
            'is_teacher': '是否为教职工',
            'is_admin': '是否为管理员'
        }
        help_texts = {
            'is_admin': '管理员可登陆cmsadmin管理页面',
        }
        widgets = {
            'identity_id': forms.TextInput(attrs={'class': 'form-control'}),
            'user_name': forms.TextInput(attrs={'class': 'form-control'}),
            'email': forms.EmailInput(attrs={'class': 'form-control'}),
            'is_teacher': forms.Select(attrs={'class': 'form-control'},
                                       choices=((True, '是'), (False, '否'))),
            'is_admin': forms.Select(attrs={'class': 'form-control'},
                                     choices=((True, '是'), (False, '否')))
        }


class UserUpdateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = User
        fields = {
            'identity_id',
            'user_name',
            'email',
            'is_teacher',
            'is_admin',
            'is_delete'
        }
        labels = {
            'identity_id': '职工号\\学号',
            'user_name': '姓名',
            'email': '电子邮件',
            'is_teacher': '是否为教职工',
            'is_admin': '是否为管理员',
            'is_delete': '删除标记'
        }
        help_texts = {
            'is_admin': '管理员可登陆cmsadmin管理页面',
            'is_delete': '标记删除后,用户无法登录本网站,可从左侧"已删除用户管理"中查看'
        }
        widgets = {
            'identity_id': forms.TextInput(attrs={'class': 'form-control'}),
            'user_name': forms.TextInput(attrs={'class': 'form-control'}),
            'email': forms.EmailInput(attrs={'class': 'form-control'}),
            'is_teacher': forms.Select(attrs={'class': 'form-control'},
                                       choices=((True, '是'), (False, '否'))),
            'is_admin': forms.Select(attrs={'class': 'form-control'},
                                     choices=((True, '是'), (False, '否'))),
        }


class MeetPlanCreateForm(forms.ModelForm, FormMixin):
    teacher = forms.ModelChoiceField(
        queryset=User.objects.filter(is_teacher=True, is_delete=False).order_by('identity_id'),
        widget=forms.Select(attrs={'class': 'form-control'}))

    class Meta:
        model = MeetPlan
        fields = [
            'teacher', 'place', 'start_time', 'end_time', 'allow_other', 'message'
        ]
        labels = {
            'teacher': '老师',
            'place': '地点',
            'start_time': '开始时间',
            'end_time': '结束时间',
            'allow_other': '允许多人',
            'message': '备注',
        }
        help_texts = {}
        widgets = {
            'place': forms.TextInput(attrs={'class': 'form-control'}),
            'start_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                     'id': 'starttimepicker',
                                                     'placeholder': 'yyyy/M/d H:m'}),
            'end_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                   'id': 'endtimepicker',
                                                   'placeholder': 'yyyy/M/d H:m'}),
            'allow_other': forms.Select(attrs={'class': 'form-control'},
                                        choices=MeetPlan.AllowOtherChoices),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }


class MeetPlanUpdateForm(forms.ModelForm, FormMixin):
    teacher = forms.ModelChoiceField(
        queryset=User.objects.filter(is_teacher=True, is_delete=False).order_by('identity_id'),
        widget=forms.Select(attrs={'class': 'form-control'}))

    class Meta:
        model = MeetPlan
        fields = [
            'teacher', 'place', 'start_time', 'end_time', 'allow_other', 'message', 'is_delete'
        ]
        labels = {
            'teacher': '老师',
            'place': '地点',
            'start_time': '开始时间',
            'end_time': '结束时间',
            'allow_other': '允许多人',
            'message': '备注',
            'is_delete': '删除标记'
        }
        help_texts = {
            'is_delete': '勾选后提交表示删除, 不会再显示'
        }
        widgets = {
            'place': forms.TextInput(attrs={'class': 'form-control'}),
            'start_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                     'id': 'starttimepicker'
                                                     }),
            'end_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                   'id': 'endtimepicker'
                                                   }),
            'allow_other': forms.Select(attrs={'class': 'form-control'},
                                        choices=MeetPlan.AllowOtherChoices),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }


class MeetPlanOrderCreateForm(forms.ModelForm, FormMixin):
    meet_plan = forms.ModelChoiceField(queryset=MeetPlan.objects.filter(is_delete=False),
                                       widget=forms.Select(attrs={'class': 'form-control'}))
    student = forms.ModelChoiceField(
        queryset=User.objects.filter(is_teacher=False, is_delete=False).order_by('identity_id'),
        widget=forms.Select(attrs={'class': 'form-control'}))

    class Meta:
        model = MeetPlanOrder
        fields = [
            'meet_plan',
            'completed',
            'student',
            'message',
        ]
        labels = {
            'message': '问题'
        }
        help_texts = {
            'message': '填写预计谈话内容，让老师有所准备：（不要超过100字）'
        }
        widgets = {
            'completed': forms.Select(attrs={'class': 'form-control'},
                                      choices=((True,'已完成'),(False, '未完成'))),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }


class MeetPlanOrderUpdateForm(forms.ModelForm, FormMixin):
    meet_plan = forms.ModelChoiceField(queryset=MeetPlan.objects.filter(is_delete=False),
                                       widget=forms.Select(attrs={'class': 'form-control'}))
    student = forms.ModelChoiceField(
        queryset=User.objects.filter(is_teacher=False, is_delete=False).order_by('identity_id'),
        widget=forms.Select(attrs={'class': 'form-control'}))

    class Meta:
        model = MeetPlanOrder
        fields = [
            'meet_plan',
            'completed',
            'student',
            'message',
            'is_delete',
        ]
        labels = {
            'completed': '是否已经完成交流',
            'is_delete': '删除标记'
        }
        help_texts = {
            'is_delete': '勾选后提交表示删除, 不会再显示, 即此次预约失效, 即便其已经完成, 也不算数'
        }
        widgets = {
            'completed': forms.Select(attrs={'class': 'form-control'},
                                      choices=((True, '已完成'), (False, '未完成'))),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }


class FeedBackUpdateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = FeedBack
        fields = [
            'have_checked'
        ]
        labels = {
            'have_checked': '状态'
        }
        widgets = {
            'have_checked': forms.Select(attrs={'class': 'form-control'},
                                         choices=FeedBack.FeedBackChoices)
        }


class SemesterDateRangeCreateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = SemesterDateRange
        fields = {
            'start_date',
            'end_date'
        }
        widgets = {
            'start_date': forms.DateInput(attrs={'class': 'form-control',
                                                     'id': 'startdatepicker'
                                                     }),
            'end_date': forms.DateInput(attrs={'class': 'form-control',
                                                   'id': 'enddatepicker'
                                                   }),
        }