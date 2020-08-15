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
    user_name, to_email = user.user_name, user.email

    info = {'active': user_id}
    token = serializer.dumps(info).decode()
    active_path = reverse('account_auth:active-account', kwargs={'token': token})

    print('-----尝试发送邮件-------')
    subject, from_email, to = '物理学院账户激活', settings.EMAIL_FROM, [to_email]
    html_content = loader.render_to_string(
        'email/account_auth/account_active_email.html',  # 需要渲染的html模板
        {
            'user_name': user_name,
            'domain': domain,
            'active_url': active_path
        }
    )

    my_send_mail.delay(subject, html_content, from_email, to)
