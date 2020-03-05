from __future__ import absolute_import, unicode_literals
from celery import shared_task
from django.contrib.auth import get_user_model

from django.core.mail import EmailMessage
from django.template import loader
from django.conf import settings
from django.urls import reverse
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer

from PKU_PHY_SU.tools.celery import TransactionAwareTask


@shared_task(base=TransactionAwareTask, bind=True)
def send_account_active_email(self, user_id, domain):
    """发送激活邮件"""
    # 加密用户信息
    serializer = Serializer(settings.SECRET_KEY, expires_in=60 * 60 * 24 * 7)
    user_model = get_user_model()
    user = user_model.objects.get(identity_id=user_id)
    user_name, to_email = user.user_name, user.email

    to_email = '598049186@qq.com'

    info = {'active': user_id}
    token = serializer.dumps(info).decode()
    active_path = reverse('account_auth:active-account', kwargs={'token': token})
    active_url = 'http://{}{}'.format(domain, active_path)

    print('-----尝试发送邮件-------')
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




