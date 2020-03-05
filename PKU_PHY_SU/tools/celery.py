from __future__ import absolute_import, unicode_literals
import os
from abc import ABC

from celery import Celery, platforms, Task

from django.conf import settings
from django.db import transaction

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'PKU_PHY_SU.settings.base')  # 设置django环境

platforms.C_FORCE_ROOT = True   # 允许用root用户启动celery

app = Celery('PKU_PHY_SU')

app.config_from_object('django.conf:settings', namespace='CELERY')  # 使用CELERY_ 作为前缀，在settings中写配置

app.autodiscover_tasks(settings.INSTALLED_APPS)  # 发现任务文件每个app下的task.py


class TransactionAwareTask(Task, ABC):
    """
    Task class which is aware of django db transactions and only executes tasks
    after transaction has been committed
    """
    abstract = True

    def apply_async(self, *args, **kwargs):
        """
        Unlike the default task in celery, this task does not return an async result
        """
        transaction.on_commit(
            lambda: super(TransactionAwareTask, self).apply_async(
                *args, **kwargs))
