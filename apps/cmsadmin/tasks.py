from __future__ import absolute_import, unicode_literals
from celery import shared_task

from PKU_PHY_SU.tools.celery import TransactionAwareTask
from apps.account_auth.tasks import send_account_active_email


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






