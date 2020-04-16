from django import forms
from utils.mixin.form import FormMixin
from ..account_auth.models import User, Grade, Department, Major
from ..meet_plan.models import MeetPlan, MeetPlanOrder, FeedBack
from ..meet_plan.utils import get_term_date


class UserForm(forms.ModelForm, FormMixin):
    field_order = ['identity_id', 'user_name', 'email', 'is_teacher', 'is_admin']

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


class MeetPlanForm(forms.ModelForm, FormMixin):
    field_order = ['teacher', 'place', 'start_time', 'end_time', 'allow_other', 'message']

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
            'teacher': forms.Select(attrs={'class': 'form-control'}),
            'place': forms.TextInput(attrs={'class': 'form-control'}),
            'start_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                     'id': 'starttimepicker',
                                                     'placeholder': 'yyyy/MM/dd HH:mm',
                                                     'readonly': 'readonly'}),
            'end_time': forms.DateTimeInput(attrs={'class': 'form-control',
                                                   'id': 'endtimepicker',
                                                   'placeholder': 'yyyy/MM/dd HH:mm',
                                                   'readonly': 'readonly'}),
            'allow_other': forms.Select(attrs={'class': 'form-control'},
                                        choices=MeetPlan.AllowOtherChoices),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }

    def __init__(self, *args, **kwargs):
        if 'teacher' in kwargs:
            flag = True
            teacher = kwargs.pop('teacher')
        else:
            flag = False

        super().__init__(*args, **kwargs)
        from db.base_model import Convert
        self.fields['teacher'].queryset = User.objects.filter(is_teacher=True).order_by(
            Convert('user_name', 'gbk').asc())

        if flag:
            self.fields['teacher'].initial = teacher


class MeetPlanOrderForm(forms.ModelForm, FormMixin):
    student = forms.ModelChoiceField(
        queryset=User.objects.filter(is_teacher=False).order_by('-identity_id'),
        widget=forms.Select(attrs={'class': 'form-control'}))

    field_order = ['meet_plan', 'student', 'completed', 'message']

    class Meta:
        model = MeetPlanOrder
        fields = [
            'meet_plan',
            'student',
            'completed',
            'message',
        ]
        labels = {
            'message': '问题'
        }
        help_texts = {
            'message': '填写预计谈话内容，让老师有所准备：（不要超过100字）'
        }
        widgets = {
            'meet_plan': forms.Select(attrs={'class': 'form-control'}),
            'completed': forms.Select(attrs={'class': 'form-control'},
                                      choices=((True, '已完成'), (False, '未完成'))),
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }

    def __init__(self, *args, **kwargs):
        if 'student' in kwargs:
            flag = True
            student = kwargs.pop('student')
        else:
            flag = False

        super().__init__(*args, **kwargs)
        date_range = get_term_date()
        self.fields['meet_plan'].queryset = MeetPlan.objects.filter(create_time__gt=date_range[0]).order_by('-id')

        if flag:
            self.fields['student'].initial = student


class FeedBackForm(forms.ModelForm, FormMixin):
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


class GradeForm(forms.ModelForm, FormMixin):
    class Meta:
        model = Grade
        fields = [
            'grade'
        ]
        labels = {
            'grade': '入学年份'
        }
        widgets = {
            'grade': forms.TextInput(attrs={'class': 'form-control'})
        }


class DepartmentForm(forms.ModelForm, FormMixin):
    class Meta:
        model = Department
        fields = [
            'department'
        ]
        labels = {
            'department': '系所/办公室'
        }
        widgets = {
            'department': forms.TextInput(attrs={'class': 'form-control'})
        }


class MajorForm(forms.ModelForm, FormMixin):
    class Meta:
        model = Major
        fields = [
            'major',
            'department'
        ]
        labels = {
            'major': '专业',
            'department': '所属系所'
        }
        widgets = {
            'major': forms.TextInput(attrs={'class': 'form-control'}),
            'department': forms.Select(attrs={'class': 'form-control'})
        }


class OptionForm(forms.Form, FormMixin):
    start = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                          'id': 'start_date',
                                                          'placeholder': 'yyyy-MM-dd',
                                                          'readonly': 'readonly'}),
                            label='学期开始日期')
    end = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                        'id': 'end_date',
                                                        'placeholder': 'yyyy-MM-dd',
                                                        'readonly': 'readonly'}),
                          label='学期结束日期')
    field_order = ['start', 'end']


class MeetPlanReportTeacherForm(forms.Form, FormMixin):
    start_date = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                               'id': 'start_date',
                                                               'placeholder': 'yyyy-MM-dd',
                                                               'readonly': 'readonly'}),
                                 label='统计开始日期')
    end_date = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                             'id': 'end_date',
                                                             'placeholder': 'yyyy-MM-dd',
                                                             'readonly': 'readonly'}),
                               label='统计结束日期', )
    field_order = ['start_date', 'end_date']

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        date_range = get_term_date()
        self.fields['start_date'].initial = date_range[0].strftime('%Y-%m-%d')
        self.fields['end_date'].initial = date_range[1].strftime('%Y-%m-%d')


class MeetPlanReportStudentForm(forms.Form, FormMixin):
    USE_CHOICES = (
        (True, '按照时间'),
        (False, '按照年级')
    )

    DETAIL_CHOICES = (
        (True, '统计'),
        (False, '不统计')
    )

    start_date = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                               'id': 'start_date',
                                                               'placeholder': 'yyyy-MM-dd',
                                                               'readonly': 'readonly'}),
                                 label='统计开始日期')
    end_date = forms.CharField(widget=forms.TextInput(attrs={'class': 'form-control',
                                                             'id': 'end_date',
                                                             'placeholder': 'yyyy-MM-dd',
                                                             'readonly': 'readonly'}),
                               label='统计结束日期', )

    use = forms.ChoiceField(widget=forms.Select(attrs={'class': 'form-control'}),
                            label='统计方式',
                            help_text='当选择按照时间时，只有统计开始时间和结束时间是有用的，年级选项可以随便选会被忽略。'
                                      '当选择按照年级时，只有年级选项是有用的，开始时间和结束时间会被忽略但必须填写。',
                            choices=USE_CHOICES
                            )

    grade = forms.ModelMultipleChoiceField(queryset=Grade.objects.all().order_by('-id'),
                                           widget=forms.SelectMultiple(attrs={'class': 'form-control'}),
                                           label='年级(可多选)')

    detail = forms.ChoiceField(widget=forms.Select(attrs={'class': 'form-control'}),
                               label='学生详细选课情况',
                               help_text='此项只在选择“按照年级”时有用，当选择统计时，将会在输出文件中输出同学们的具体选课情况。'
                                         '当选择不统计时，只会输出同学们的已完成总数。',
                               choices=DETAIL_CHOICES, initial=False)
    field_order = ['start_date', 'end_date', 'use', 'grade', 'detail']

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        date_range = get_term_date()
        self.fields['start_date'].initial = date_range[0].strftime('%Y-%m-%d')
        self.fields['end_date'].initial = date_range[1].strftime('%Y-%m-%d')
