from __future__ import absolute_import, unicode_literals
from celery import shared_task
from django.contrib.auth import get_user_model
from django.template import loader
from django.conf import settings
from django.urls import reverse
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer

from PKU_PHY_SU.tools.celery import TransactionAwareTask, my_send_mail


@shared_task(base=TransactionAwareTask, bind=True)
def send_account_active_email(self, user_id, domain):
    """发送激活邮件"""
    # 加密用户信息
    serializer = Serializer(settings.SECRET_KEY, expires_in=60 * 60 * 24 * 7)
    user_model = get_user_model()
    user = user_model.objects.get(identity_id=user_id)
    user_name, to_email = user.user_name, user.email

    info = {'active': user_id}
    token = serializer.dumps(info).decode()
    active_path = reverse('account_auth:active-account', kwargs={'token': token})
    active_url = 'http://{}{}'.format(domain, active_path)

    print('-----尝试发送邮件-------')
    subject, from_email, to = '物理学院账户激活', settings.EMAIL_FROM, [to_email]
    html_content = loader.render_to_string(
        'email/account_auth/account_active_email.html',               #需要渲染的html模板
                        {
                            'user_name': user_name,
                            'domain': domain,
                            'active_url': active_url
                        }
                   )

    my_send_mail.delay(subject, html_content, from_email, to)


@shared_task(base=TransactionAwareTask, bind=True)
def create_many_user(self, file_id, domain):
    from apps.filemanager.models import File
    file = File.objects.get(id=file_id).file

    import os
    ext = os.path.splitext(file.name)[1]
    if ext == '.csv':
        import csv
        try:
            with open(file.path, 'r', encoding='utf-8') as f:
                reader = csv.reader(f)
                from apps.account_auth.models import User
                for [identity, name, is_teacher] in reader:
                    if reader.line_num == 1:
                        # 跳过首行
                        continue
                    user = User.objects.get_or_create(identity_id=identity)
                    # 如果不存在则创建
                    if user[1]:
                        user[0].user_name = name
                        user[0].is_teacher = True if is_teacher == '教职工' else False
                        user[0].save()
                        # 发送激活邮件
                        send_account_active_email.delay(user_id=identity, domain=domain)

        except Exception as e:
            print(e)
    elif ext == '.xlsx':
        # TODO： 补充其他拓展格式
        pass
    elif ext == '.xls':
        pass
    else:
        pass






