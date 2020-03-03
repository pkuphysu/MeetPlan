from __future__ import absolute_import, unicode_literals
import os
from celery import Celery

from django.conf import settings

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'PKU_PHY_SU.settings.base')  # 设置django环境

app = Celery('PKU_PHY_SU')

app.config_from_object('django.conf:settings', namespace='CELERY')  # 使用CELERY_ 作为前缀，在settings中写配置

app.autodiscover_tasks(settings.INSTALLED_APPS)  # 发现任务文件每个app下的task.py
