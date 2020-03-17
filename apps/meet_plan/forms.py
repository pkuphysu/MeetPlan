from django import forms

from utils.mixin.form import FormMixin
from .models import MeetPlan, MeetPlanOrder, FeedBack


class MeetPlanForm(forms.ModelForm, FormMixin):
    field_order = ['place', 'start_time', 'end_time', 'allow_other', 'message']

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
            'completed'
        ]
        labels = {
            'completed': '是否已经完成交流'
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
