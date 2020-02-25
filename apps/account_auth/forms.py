from django import forms
from phonenumber_field.widgets import PhoneNumberPrefixWidget
from utils.mixin import FormMixin
from .models import User, UserProfile


class UserEmailUpdateForm(forms.ModelForm, FormMixin):
    # email = forms.EmailField(widget=forms.TextInput(attrs={'class': 'form-control'}))
    class Meta:
        model = User
        fields = ['email']
        labels = {
            'email': '电子邮件',
        }
        help_texts = {
            'email': '用户电子邮件.',
        }
        widgets = {
            'email': forms.EmailInput(attrs={'class': 'form-control'}),
        }
        # error_messages = {
        #     'email': {
        #         'max_length': ("This writer's name is too long."),
        #     },
        # }


class UserProfileCreateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = UserProfile
        fields = ['birth', 'gender', 'telephone']
        labels = {
            'gender': '性别',
            'telephone': '联系方式',
            'birth': '生日',
        }
        widgets = {
            'gender': forms.Select(attrs={'class': 'form-control'},
                                   choices=UserProfile.GenderChoices),
            'telephone': PhoneNumberPrefixWidget(attrs={'class': 'form-control'}, initial='CN'),
            'birth': forms.TextInput(attrs={'class': 'form-control',
                                            'id': 'datepicker',
                                            }),
        }


class UserProfileUpdateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = UserProfile
        fields = ['birth', 'gender', 'telephone']
        labels = {
            'gender': '性别',
            'telephone': '联系方式',
            'birth': '生日',
        }
        widgets = {
            'gender': forms.Select(attrs={'class': 'form-control'},
                                   choices=UserProfile.GenderChoices),
            'telephone': PhoneNumberPrefixWidget(attrs={'class': 'form-control'}, initial='CN'),
            'birth': forms.TextInput(attrs={'class': 'form-control',
                                            'id': 'datepicker',
                                            }),
        }


class UserProfileImgUpdateForm(forms.ModelForm, FormMixin):
    class Meta:
        model = UserProfile
        fields = ['user_img']
        labels = {
            'user_img': '头像',
        }
        # widgets = {
        #     'user_img': forms.FileInput()
        # }
