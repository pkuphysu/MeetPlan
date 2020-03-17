from django import forms
from django.contrib import admin
from django.contrib.auth.admin import UserAdmin
from django.contrib.auth.forms import ReadOnlyPasswordHashField
from django.urls import reverse
from django.utils.html import escape, mark_safe

from . import models

# Register your models here.


class UserCreationForm(forms.ModelForm):
    password1 = forms.CharField(label='Password', widget=forms.PasswordInput)
    password2 = forms.CharField(label='Password confirmation', widget=forms.PasswordInput)

    class Meta:
        model = models.User
        fields = ('identity_id', 'user_name', 'email')

    def clean_password2(self):
        # Check that the two password entries match
        password1 = self.cleaned_data.get("password1")
        password2 = self.cleaned_data.get("password2")
        if password1 and password2 and password1 != password2:
            raise forms.ValidationError("Passwords don't match")
        return password2

    def save(self, commit=True):
        user = super(UserCreationForm, self).save(commit=False)
        user.set_password(self.cleaned_data["password1"])
        if commit:
            user.save()
        return user


class UserChangeForm(forms.ModelForm):
    password = ReadOnlyPasswordHashField()

    class Meta:
        model = models.User
        fields = ('identity_id', 'user_name', 'email', 'is_active', 'is_teacher', 'is_superuser', 'is_delete')

    def clean_password(self):
        # Regardless of what the user provides, return the initial value.
        # This is done here, rather than on the field, because the
        # field does not have access to the initial value
        return self.initial["password"]


class PHYUserAdmin(UserAdmin):
    form = UserChangeForm
    add_form = UserCreationForm

    list_display = [
        'identity_id',
        'user_name',
        'email',
        'is_active',
        'is_teacher'
    ]
    list_filter = ('is_active', 'is_teacher', 'is_superuser')
    search_fields = ('identity_id', 'user_name')
    fieldsets = (
        (None, {'fields': ('identity_id', 'user_name', 'email', 'password')}),
        ('Status', {'fields': ('is_active', 'is_delete')}),
        ('Permissions', {'fields': ('is_teacher', 'is_superuser')}),
    )
    add_fieldsets = (
        (None, {
            'classes': ('wide',),
            'fields': ('identity_id', 'user_name',
                       'email', 'password1', 'password2',
                       'is_teacher', 'is_superuser')}
         ),
    )
    list_per_page = 20
    list_select_related = True
    ordering = ('identity_id',)


admin.site.register(models.User, PHYUserAdmin)


class BaseProfileAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'gender',
        'birth',
        'link_to_head_picture'
    ]

    def link_to_head_picture(self, obj):
        link = reverse("admin:filemanager_myimg_change", args=[obj.head_picture_id])
        return mark_safe(f'<a href="{link}">{escape(obj.head_picture.__str__())}</a>')

    link_to_head_picture.short_description = '头像'
    list_filter = ['gender']
    search_fields = ['user', 'telephone']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.BaseProfile, BaseProfileAdmin)


class GradeAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'grade',
    ]
    search_fields = ['grade']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.Grade, GradeAdmin)


class DepartmentAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'department',
    ]
    search_fields = ['department']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.Department, DepartmentAdmin)


class MajorAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'major',
        'link_to_department',
    ]

    def link_to_department(self, obj):
        link = reverse("admin:account_auth_department_change", args=[obj.department_id])
        return mark_safe(f'<a href="{link}">{escape(obj.department.__str__())}</a>')

    link_to_department.short_description = '系所'

    search_fields = ['major', 'link_to_department']
    list_filter = ['department']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.Major, MajorAdmin)


class StudentProfileAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'is_graduate',
        'grade',
        'phone_number',
        'department',
        'major',
        'dorm'
    ]

    list_filter = ['is_graduate', 'department', 'major', 'grade']
    search_fields = ['user', 'phone_number']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.StudentProfile, StudentProfileAdmin)


class TeacherProfileAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'phone_number',
        'department',
        'office',
        'introduce'
    ]

    list_filter = ['department']
    search_fields = ['user', 'office', 'introduce']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.TeacherProfile, TeacherProfileAdmin)