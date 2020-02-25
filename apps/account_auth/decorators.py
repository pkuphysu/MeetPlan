from django.shortcuts import redirect
from functools import wraps
from django.contrib.auth.models import Permission, ContentType
from django.http import Http404

from utils import rest


def user_login_required(func):
    """判断是否登录"""

    def wapper(request, *args, **kwargs):
        if request.user.is_authenticated:
            return func(request, *args, **kwargs)
        else:
            if request.is_ajax():
                return rest.un_auth(message="请先登录!")
            else:
                return redirect('/')

    return wapper


def user_permission_required(model):
    def decorator(viewfunc):
        @wraps(viewfunc)
        def _wrapper(request, *args, **kwargs):
            content_type = ContentType.objects.get_for_model(model)
            permissions = Permission.objects.filter(content_type=content_type)

            codenames = [content_type.app_label + "." +
                         permission.codename for permission in permissions]
            result = request.user.has_perms(codenames)
            if result:
                return viewfunc(request, *args, **kwargs)
            else:
                print('=' * 20)
                raise Http404()

        return _wrapper

    return decorator


def user_superuser_required(viewfunc):
    @wraps(viewfunc)
    def _wrapper(request, *args, **kwargs):
        if request.user.is_superuser:
            return viewfunc(request, *args, **kwargs)
        else:
            raise Http404()

    return _wrapper


def user_teacher_required(viewfunc):
    @wraps(viewfunc)
    def _wrapper(request, *args, **kwargs):
        if request.user.is_teacher:
            return viewfunc(request, *args, **kwargs)
        else:
            raise Http404()

    return _wrapper
