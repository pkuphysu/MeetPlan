from django import forms

from utils.mixin.form import FormMixin
from .models import MeetPlan, MeetPlanOrder, FeedBack


class MeetPlanCreateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = MeetPlan
        fields = [
            'place', 'start_time', 'end_time', 'allow_other', 'message'
        ]
        labels = {
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

    class Meta:
        model = MeetPlan
        fields = [
            'place', 'start_time', 'end_time', 'allow_other', 'message', 'is_delete'
        ]
        labels = {
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
    class Meta:
        model = MeetPlanOrder
        fields = [
            'message'
        ]
        labels = {
            'message': ''
        }
        help_texts = {
            'message': '填写预计谈话内容，让老师有所准备：（不要超过100字）'
        }
        widgets = {
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'})
        }


class MeetPlanOrderUpdateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = MeetPlanOrder
        fields = [
            'completed', 'is_delete'
        ]
        labels = {
            'completed': '是否已经完成交流',
            'is_delete': '删除标记'
        }
        help_texts = {
            'is_delete': '勾选后提交表示删除, 不会再显示, 即此次预约失效, 即便其已经完成, 也不算数'
        }


class FeedBackCreateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = FeedBack
        fields = [
            'message'
        ]
        labels = {
            'message': '反馈'
        }
        help_texts = {
            'message': '使用过程中的反馈'
        }

        widgets = {
            'message': forms.Textarea(attrs={'class': 'form-control',
                                             'row': '5',
                                             'placeholder': 'Enter...'
                                             }),
        }
