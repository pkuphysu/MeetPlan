"""
Django settings for PKU_PHY_SU project.

Generated by 'django-admin startproject' using Django 2.2.5.

For more information on this file, see
https://docs.djangoproject.com/en/2.2/topics/settings/

For the full list of settings and their values, see
https://docs.djangoproject.com/en/2.2/ref/settings/
"""

import os

# 根据环境变量导入不同设置文件
# "export PHY_ENV='DEVELOP'"
if not os.environ.get('PHY_ENV'):
    from .pro import *
elif os.environ.get('PHY_ENV') == 'TEST':
    from .test import *
else:
    from .dev import *

# Build paths inside the project like this: os.path.join(BASE_DIR, ...)
BASE_DIR = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


# Application definition
INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',

    'django_celery_results',
    'django_celery_beat',

    'rest_framework',
    'phonenumber_field',

    'apps.account_auth.apps.AccountAuthConfig',
    'apps.filemanager.apps.FilemanagerConfig',
    'apps.cmsadmin.apps.CmsadminConfig',
    'apps.portal.apps.PortalConfig',
    'apps.meet_plan.apps.MeetPlanConfig',
]

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]

ROOT_URLCONF = 'PKU_PHY_SU.urls'

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, 'templates')]
        ,
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

WSGI_APPLICATION = 'PKU_PHY_SU.wsgi.application'

# Password validation
# https://docs.djangoproject.com/en/2.2/ref/settings/#auth-password-validators

AUTH_PASSWORD_VALIDATORS = [
    {
        'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator',
    },
    {
        'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator',
    },
]

# Internationalization
# https://docs.djangoproject.com/en/2.2/topics/i18n/

LANGUAGE_CODE = 'zh-hans'

TIME_ZONE = 'Asia/Shanghai'

USE_I18N = True

USE_L10N = True

USE_TZ = True


# Static files (CSS, JavaScript, Images)
# https://docs.djangoproject.com/en/2.2/howto/static-files/


MEDIA_URL = '/media/uploads/'
MEDIA_ROOT = os.path.join(BASE_DIR, "media/uploads/") # 项目目录下的media目录 需要在项目目录下创建media目录

STATIC_URL = '/static/'
# 开发阶段放置项目自己的静态文件
STATICFILES_DIRS = [os.path.join(BASE_DIR, "staticfiles"), ]

# 执行collectstatic命令后会将项目中的静态文件收集到该目录下面来（所以不应该在该目录下面放置自己的一些静态文件，因为会覆盖掉）, 生产环境用 Nginx 处理
STATIC_ROOT = os.path.join(BASE_DIR, 'static')

# Django 认证系统使用的模型类
AUTH_USER_MODEL = 'account_auth.User'

# 配置登录url地址
# LOGIN_URL = '/account/login/iaaa'  # /user/login/iaaa?next=
LOGIN_URL = '/account_auth/login/iaaa'

# Django Session 使用 Redis 缓存
SESSION_ENGINE = "django.contrib.sessions.backends.cache"
SESSION_CACHE_ALIAS = "default"

# Celery 配置
CELERY_WORKER_MAX_TASKS_PER_CHILD = 100000  # 每个worker执行10w个任务就会被销毁，可防止内存泄露
CELERY_RESULT_SERIALIZER = 'json'  # 结果序列化方案
CELERY_RESULT_BACKEND = 'django-db'  # BACKEND配置，这里使用redis
CELERY_TIMEZONE = 'Asia/Shanghai'

DEFAULT_FILE_STORAGE = 'PKU_PHY_SU.tools.storage.FileStorage'
