from __future__ import absolute_import, unicode_literals
from celery import shared_task

from django.core.mail import EmailMultiAlternatives, EmailMessage
from django.template import loader

from django.conf import settings

import time


@shared_task
def send_account_active_email(to_email, user_name, active_url, domain):
    '''发送激活邮件'''
    # 组织邮件信息
    subject, from_email, to = '物理学院账户激活', settings.EMAIL_FROM, [to_email]
    html_content = loader.render_to_string(
                        'email/account_active_email.html',               #需要渲染的html模板
                        {
                            'user_name': user_name,
                            'domain': domain,
                            'active_url': active_url
                        }
                   )
    msg = EmailMessage(subject=subject, body=html_content, from_email=from_email, to=to)
    msg.content_subtype = "html"
    msg.send()




