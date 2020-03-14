from __future__ import absolute_import, unicode_literals
from celery import shared_task

from PKU_PHY_SU.tools.celery import TransactionAwareTask
from apps.account_auth.tasks import send_account_active_email


@shared_task(base=TransactionAwareTask, bind=True)
def account_create_many_user(self, file_id, domain):
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


@shared_task(base=TransactionAwareTask, bind=True)
def meetplan_create_teacher_report(self, user_id, app_name, start_date, end_date):
    import csv
    import os
    import time
    import uuid
    import dateutil.parser
    from django.conf import settings
    from apps.meet_plan.models import MeetPlan, MeetPlanOrder
    from apps.account_auth.models import User
    from apps.filemanager.models import File as MyFile
    from django.core.files import File

    d = os.path.join(settings.MEDIA_ROOT, 'generate')
    fn = time.strftime('%Y-%m-%d-%H-%M-%S')
    fn = fn + '_%s' % str(uuid.uuid4())
    file_path = os.path.join(d, fn + '.csv')

    if not os.path.exists(os.path.dirname(file_path)):
        os.mkdir(os.path.dirname(file_path))

    start_date = dateutil.parser.parse('{}{}'.format(start_date, 'T00:00:00+08:00'))
    end_date = dateutil.parser.parse('{}{}'.format(end_date, 'T00:00:00+08:00'))

    with open(file_path, 'w', encoding='UTF-8') as fp:
        writer = csv.writer(fp)
        writer.writerow(['职工号', '姓名', '邮件', '系所', '总安排数', '总预约人数', '总完成数'])
        queryset = User.objects.filter(is_teacher=True)
        for user in queryset:
            department = user.teacherprofile.department.department if hasattr(user, 'teacherprofile') else '未定义'
            mt_queryset = MeetPlan.objects.filter(teacher=user,
                                                  start_time__gte=start_date,
                                                  end_time__lte=end_date)
            mto_queryset = MeetPlanOrder.objects.filter(meet_plan__in=mt_queryset)

            writer.writerow([user.identity_id, user.user_name, user.email, department,
                             mt_queryset.count(), mto_queryset.count(),
                             mto_queryset.filter(completed=True).count()])

    file = MyFile.objects.create(upload_or_generate=False, app=app_name, user_id=user_id,
                                 remark='教师报表，时间从{}到{}'.format(start_date.strftime('%Y-%m-%d'),
                                                               end_date.strftime('%Y-%m-%d')))
    file.file.name = fn + '.csv'
    file.save()


@shared_task(base=TransactionAwareTask, bind=True)
def meetplan_create_student_report(self, user_id, app_name, start_date, end_date):
    pass