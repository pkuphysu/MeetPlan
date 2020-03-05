from django.contrib.auth.decorators import login_required
from django.contrib.admin.views.decorators import staff_member_required
from django.conf import settings

from django.contrib.auth import REDIRECT_FIELD_NAME
from django.contrib.auth.decorators import user_passes_test


def admin_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME, login_url='admin:login'):
    """
    Decorator for views that checks that the user is logged in and is a admin
    member, redirecting to the login page if necessary.
    """
    actual_decorator = user_passes_test(
        lambda u: u.is_active and u.is_admin and not u.is_delete,
        login_url=login_url,
        redirect_field_name=redirect_field_name
    )
    if view_func:
        return actual_decorator(view_func)
    return actual_decorator


def teacher_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME, login_url='admin:login'):
    """
    Decorator for views that checks that the user is logged in and is a admin
    member, redirecting to the login page if necessary.
    """
    actual_decorator = user_passes_test(
        lambda u: u.is_active and u.is_teacher and not u.is_delete,
        login_url=login_url,
        redirect_field_name=redirect_field_name
    )
    if view_func:
        return actual_decorator(view_func)
    return actual_decorator


def have_profile_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME,
                          redirect_filed_url='account_auth:userprofile_add'):
    """
    Decorator for views that checks that the user is logged in and is a admin
    member, redirecting to the login page if necessary.
    """
    actual_decorator = user_passes_test(
        lambda u: u.is_active and hasattr(u, 'userprofile') and not u.is_delete,
        login_url=redirect_filed_url,
        redirect_field_name=redirect_field_name
    )
    if view_func:
        return actual_decorator(view_func)
    return actual_decorator


class AdminRequiredMixin:
    """
    普通管理员验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        view = super(AdminRequiredMixin, cls).as_view(**initkwargs)
        return admin_required(view_func=view, login_url=settings.LOGIN_URL)


class TeacherRequiredMixin:
    """
    教师验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        view = super(TeacherRequiredMixin, cls).as_view(**initkwargs)
        return teacher_required(view_func=view, login_url=settings.LOGIN_URL)


class StaffRequiredMixin:
    """
    超级管理员验证类, 此类用户可访问 django 默认 admin 后台, 本项目暂时用不到
    """
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(StaffRequiredMixin, cls).as_view(**initkwargs)
        return staff_member_required(view_func=view, login_url=settings.LOGIN_URL)


class LoginRequiredMixin:
    """
    登录验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(LoginRequiredMixin, cls).as_view(**initkwargs)
        return login_required(function=view, login_url=settings.LOGIN_URL)


class UserProfileRequiredMixin:
    """
    登录验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(UserProfileRequiredMixin, cls).as_view(**initkwargs)
        return have_profile_required(view_func=view)


