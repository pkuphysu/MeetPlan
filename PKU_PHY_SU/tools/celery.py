from __future__ import absolute_import, unicode_literals
import os
from abc import ABC

from celery import Celery, platforms, Task, shared_task

from django.conf import settings
from django.core.mail import EmailMessage
from django.db import transaction

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'PKU_PHY_SU.settings.base')  # 设置django环境

platforms.C_FORCE_ROOT = True   # 允许用root用户启动celery

app = Celery('PKU_PHY_SU')

app.config_from_object('django.conf:settings', namespace='CELERY')  # 使用CELERY_ 作为前缀，在settings中写配置

app.autodiscover_tasks()  # 发现任务文件每个app下的task.py


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


@shared_task(base=TransactionAwareTask, bind=True)
def my_send_mail(self, subject, html_content, from_email, to):
    if settings.DEBUG:
        to = ['598049186@qq.com']

    msg = EmailMessage(subject=subject, body=html_content, from_email=from_email, to=to)
    msg.content_subtype = "html"
    try:
        msg.send()
    except Exception as e:
        """
                邮件发送失败，使用retry进行重试

                retry的参数可以有：
                    exc：指定抛出的异常
                    throw：重试时是否通知worker是重试任务
                    eta：指定重试的时间／日期
                    countdown：在多久之后重试（每多少秒重试一次）
                    max_retries：最大重试次数
                """
        raise self.retry(exc=e, countdown=60*5, max_retries=5)