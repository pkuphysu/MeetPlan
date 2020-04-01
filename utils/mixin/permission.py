from django.contrib.auth.decorators import login_required
from django.contrib.admin.views.decorators import staff_member_required
from django.conf import settings
from django.core.exceptions import PermissionDenied
from django.contrib.auth import REDIRECT_FIELD_NAME
from django.contrib.auth.decorators import user_passes_test
from functools import wraps


def user_admin_required(viewfunc):
    @wraps(viewfunc)
    def _wrapper(request, *args, **kwargs):
        if request.user.is_admin:
            return viewfunc(request, *args, **kwargs)
        else:
            raise PermissionDenied("您不是管理员！")

    return _wrapper


def user_teacher_required(viewfunc):
    @wraps(viewfunc)
    def _wrapper(request, *args, **kwargs):
        if request.user.is_teacher:
            return viewfunc(request, *args, **kwargs)
        else:
            raise PermissionDenied("您的身份不是教师！")

    return _wrapper


def user_student_required(viewfunc):
    @wraps(viewfunc)
    def _wrapper(request, *args, **kwargs):
        if not request.user.is_teacher:
            return viewfunc(request, *args, **kwargs)
        else:
            raise PermissionDenied("您的身份不是学生！")

    return _wrapper


def have_base_profile_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME,
                               redirect_filed_url='account_auth:baseprofile_add'):
    """
    Decorator for views that checks that the user is logged in and is a admin
    member, redirecting to the login page if necessary.
    """
    actual_decorator = user_passes_test(
        lambda u: hasattr(u, 'baseprofile'),
        login_url=redirect_filed_url,
        redirect_field_name=redirect_field_name
    )
    if view_func:
        return actual_decorator(view_func)
    return actual_decorator


def have_tea_profile_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME):
    actual_decorator = user_passes_test(
        lambda u:  hasattr(u, 'teacherprofile'),
        login_url='account_auth:teacher-profile-create',
        redirect_field_name=redirect_field_name
    )
    if view_func:
        return actual_decorator(view_func)
    return actual_decorator


def have_stu_profile_required(view_func=None, redirect_field_name=REDIRECT_FIELD_NAME):
    actual_decorator = user_passes_test(
        lambda u:  hasattr(u, 'studentprofile'),
        login_url='account_auth:student-profile-create',
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
        # return admin_required(view_func=view, login_url=settings.LOGIN_URL)
        return user_admin_required(viewfunc=view)


class TeacherRequiredMixin:
    """
    教师验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        view = super(TeacherRequiredMixin, cls).as_view(**initkwargs)
        # return teacher_required(view_func=view, login_url=settings.LOGIN_URL)
        return user_teacher_required(viewfunc=view)


class StudentRequiredMixin:
    """
    学生验证类
    """
    @classmethod
    def as_view(cls, **initkwargs):
        view = super(StudentRequiredMixin, cls).as_view(**initkwargs)
        # return teacher_required(view_func=view, login_url=settings.LOGIN_URL)
        return user_student_required(viewfunc=view)


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


class BaseProfileRequiredMixin:
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(BaseProfileRequiredMixin, cls).as_view(**initkwargs)
        return have_base_profile_required(view_func=view)


class StuProfileRequiredMixin:
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(StuProfileRequiredMixin, cls).as_view(**initkwargs)
        return have_stu_profile_required(view_func=view)


class TeaProfileRequiredMixin:
    @classmethod
    def as_view(cls, **initkwargs):
        # 调用父类的as_view
        view = super(TeaProfileRequiredMixin, cls).as_view(**initkwargs)
        return have_tea_profile_required(view_func=view)


class ViewMixin(LoginRequiredMixin, BaseProfileRequiredMixin):
    pass


class StuViewMixin(LoginRequiredMixin, BaseProfileRequiredMixin, StudentRequiredMixin, StuProfileRequiredMixin):
    pass


class TeaViewMixin(LoginRequiredMixin, BaseProfileRequiredMixin, TeacherRequiredMixin, TeaProfileRequiredMixin):
    pass

