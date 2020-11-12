from __future__ import absolute_import, unicode_literals

from celery import shared_task
from django.conf import settings
from django.contrib.auth import get_user_model
from django.template import loader
from django.urls import reverse
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer

from MeetPlan.tools.celery import TransactionAwareTask, my_send_mail


@shared_task(base=TransactionAwareTask, bind=True)
def send_account_active_email(self, user_id):
    """发送激活邮件"""
    # 加密用户信息
    domain = settings.SITE_URL
    serializer = Serializer(settings.SECRET_KEY, expires_in=60 * 60 * 24 * 7)
    user_model = get_user_model()
    user = user_model.objects.get(identity_id=user_id)
    user_name, to_email = user.user_name, user.get_email()

    info = {'active': user_id}
    token = serializer.dumps(info).decode()
    active_path = reverse('account_auth:active-account', kwargs={'token': token})

    subject, from_email = '物理学院综合指导课账户激活', settings.EMAIL_FROM
    html_content = loader.render_to_string(
        'email/account_auth/account_active_email.html',  # 需要渲染的html模板
        {
            'user_name': user_name,
            'domain': domain,
            'active_url': active_path
        }
    )

    my_send_mail.delay(subject, html_content, from_email, to_email)


@shared_task(base=TransactionAwareTask, bind=True)
def send_account_register_email(self, user_id):
    """发送注册邮件"""
    # 加密用户信息
    domain = settings.SITE_URL

    user_model = get_user_model()
    user = user_model.objects.get(identity_id=user_id)
    user_name, to_email = user.user_name, user.get_email()

    subject, from_email = '物理学院综合指导课账户注册', settings.EMAIL_FROM
    html_content = loader.render_to_string(
        'email/account_auth/account_register_email.html',  # 需要渲染的html模板
        {
            'user_name': user_name,
            'domain': domain,
        }
    )

    my_send_mail.delay(subject, html_content, from_email, to_email)


@shared_task
def deactivate_user_every_eight_weeks():
    domain = settings.SITE_URL
    from_email = settings.EMAIL_FROM
    from django.utils import timezone
    import datetime
    from .models import User
    users = User.objects.filter(is_active=True)
    duration = datetime.timedelta(weeks=8)
    for user in users:
        if user.last_login is None or user.last_login + duration < timezone.now():
            user.is_active = False
            user.save()
            subject = '物理学院综合指导课账户锁定通知'
            html = loader.render_to_string(
                'email/account_auth/account_deactive_email.html',
                {
                    'domain': domain,
                    'user_name': user.user_name,
                    'last_login': user.last_login,
                }
            )
            to_email = user.get_email()

            my_send_mail.delay(subject, html, from_email, to_email)
