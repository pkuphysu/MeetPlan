from django import forms
from utils.mixin.form import FormMixin
from .models import User, BaseProfile, StudentProfile, TeacherProfile, Department, Major


class UserEmailForm(forms.ModelForm, FormMixin):
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


class UserProfileForm(forms.ModelForm, FormMixin):
    class Meta:
        model = BaseProfile
        fields = ['birth', 'gender']
        labels = {
            'gender': '性别',
            'birth': '生日',
        }
        widgets = {
            'gender': forms.Select(attrs={'class': 'form-control'},
                                   choices=BaseProfile.GenderChoices),
            'birth': forms.TextInput(attrs={'class': 'form-control',
                                            'id': 'datepicker'}),
        }


class StudentProfileForm(forms.ModelForm, FormMixin):
    field_order = ['is_graduate', 'grade', 'phone_number', 'department', 'major', 'dorm']

    class Meta:
        model = StudentProfile
        fields = [
            'is_graduate',
            'phone_number',
            'department',
            'major',
            'dorm',
            'grade'
        ]
        labels = {
            'is_graduate': '身份',
            'phone_number': '手机',
            'department': '系所',
            'major': '专业',
            'dorm': '宿舍',
            'grade': '年级',
        }
        widgets = {
            'department': forms.Select(attrs={'class':'form-control'}),
            'major': forms.Select(attrs={'class': 'form-control'}),
            'is_graduate': forms.Select(attrs={'class': 'form-control'},
                                        choices=StudentProfile.GRADUATE_CHOICES),
            'phone_number': forms.TextInput(attrs={'class': 'form-control'}),
            'dorm': forms.TextInput(attrs={'class': 'form-control'}),
            'grade': forms.Select(attrs={'class': 'form-control'})
        }

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.fields['major'].queryset = Major.objects.none()
        if 'department' in self.data:
            try:
                department_id = int(self.data.get('department'))
                self.fields['major'].queryset = Major.objects.filter(department_id=department_id)
            except (ValueError, TypeError):
                pass  # invalid input from the client; ignore and fallback to empty City queryset
        elif self.instance.pk:
            self.fields['major'].queryset = self.instance.department.major_set


class TeacherProfileForm(forms.ModelForm, FormMixin):
    field_order = ['department', 'office', 'phone_number', 'introduce']

    class Meta:
        model = TeacherProfile
        fields = ['phone_number', 'department', 'office', 'introduce']
        labels = {
            'phone_number': '办公室电话',
            'department': '系所',
            'office': '办公室',
            'introduce': '个人简介',
        }
        widgets = {
            'phone_number': forms.TextInput(attrs={'class': 'form-control'}),
            'department': forms.Select(attrs={'class': 'form-control'}),
            'office': forms.TextInput(attrs={'class': 'form-control'}),
            'introduce': forms.Textarea(attrs={'class': 'form-control',
                                               'row': '5',
                                               'placeholder': 'Enter...'}
                                        )
        }

