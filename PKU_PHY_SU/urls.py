"""PKU_PHY_SU URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.2/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    # django 默认后台, 仅限技术管理人员登录
    path('superadmin/', admin.site.urls),
    # 管理员内容管理页面
    path('cmsadmin/', include('apps.cmsadmin.urls')),
    # 综合指导课相关页面
    path('meetplan/', include('apps.meet_plan.urls')),
    # 账户管理相关页面
    path('account_auth/', include('apps.account_auth.urls')),
    # 网站首页页面
    path('', include('apps.portal.urls')),
]
