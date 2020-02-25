from django import forms
from django.contrib import admin
from django.contrib.auth.admin import UserAdmin
from django.contrib.auth.models import Group
from django.contrib.auth.forms import ReadOnlyPasswordHashField

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


class UserProfileAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'gender',
        'birth',
        'telephone',
        'user_img'
    ]
    list_filter = ['gender']
    search_fields = ['user', 'telephone']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.UserProfile, UserProfileAdmin)
