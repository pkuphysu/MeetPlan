from __future__ import absolute_import, unicode_literals

import chardet
from celery import shared_task

from MeetPlan.tools.celery import TransactionAwareTask
from apps.account_auth.tasks import send_account_register_email


def get_encoding(file):
    with open(file, 'rb') as f:
        return chardet.detect(f.readline())['encoding']


@shared_task(base=TransactionAwareTask)
def account_create_many_user(file_id):
    from apps.filemanager.models import MyFile
    file = MyFile.objects.get(id=file_id).file

    import os
    ext = os.path.splitext(file.name)[1]
    if ext == '.csv':
        import csv
        try:
            with open(file.path, 'r', encoding=get_encoding(file.path)) as f:
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
                        send_account_register_email.delay(user_id=identity)

        except Exception as e:
            print(e)
    elif ext == '.xlsx':
        # TODO： 补充其他拓展格式
        pass
    elif ext == '.xls':
        pass
    else:
        pass


@shared_task(base=TransactionAwareTask)
def meetplanorder_create_many(file_id):
    from apps.filemanager.models import MyFile
    file = MyFile.objects.get(id=file_id).file

    import os
    ext = os.path.splitext(file.name)[1]
    if ext == '.csv':
        import csv
        from apps.meet_plan.models import MeetPlan, MeetPlanOrder
        from django.utils import timezone
        try:
            with open(file.path, 'r', encoding=get_encoding(file.path)) as f:
                reader = csv.reader(f)
                from apps.account_auth.models import User
                for [identity, name] in reader:
                    if reader.line_num == 1:
                        # 跳过首行
                        continue
                    student = User.objects.get(identity_id=identity)
                    teacher = User.objects.get(identity_id='0000000000')
                    mt = MeetPlan.objects.create(teacher=teacher,
                                                 place='补录',
                                                 start_time=timezone.now(),
                                                 end_time=timezone.now(),
                                                 message='本科生科研替代',
                                                 available_choice=3,
                                                 )
                    mt.save()
                    mto = MeetPlanOrder.objects.create(meet_plan=mt,
                                                       student=student,
                                                       completed=True,
                                                       message='本科生科研替代')
                    mto.save()
                    mto = MeetPlanOrder.objects.create(meet_plan=mt,
                                                       student=student,
                                                       completed=True,
                                                       message='本科生科研替代')
                    mto.save()
                    mto = MeetPlanOrder.objects.create(meet_plan=mt,
                                                       student=student,
                                                       completed=True,
                                                       message='本科生科研替代')
                    mto.save()

        except Exception as e:
            print(e)
    elif ext == '.xlsx':
        # TODO： 补充其他拓展格式
        pass
    elif ext == '.xls':
        pass
    else:
        pass


@shared_task(base=TransactionAwareTask)
def meetplan_create_teacher_report(user_id, app_name, start_date, end_date):
    import csv
    import os
    import time
    import uuid
    import dateutil.parser
    from django.conf import settings
    from apps.meet_plan.models import MeetPlan, MeetPlanOrder
    from apps.account_auth.models import User
    from apps.filemanager.models import MyFile

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
    file.file.name = 'generate/' + fn + '.csv'
    file.save()


@shared_task(base=TransactionAwareTask)
def meetplan_create_student_report(user_id, app_name, start_date, end_date, grades, date_or_grade, detail):
    import csv
    import os
    import time
    import uuid
    import dateutil.parser
    from django.conf import settings
    from apps.meet_plan.models import MeetPlan, MeetPlanOrder
    from apps.account_auth.models import User, Grade
    from apps.filemanager.models import MyFile

    if not os.path.exists(settings.MEDIA_ROOT):
        os.mkdir(settings.MEDIA_ROOT)

    d = os.path.join(settings.MEDIA_ROOT, 'generate')
    if not os.path.exists(d):
        os.mkdir(d)

    fn = time.strftime('%Y-%m-%d-%H-%M-%S')
    fn = fn + '_%s' % str(uuid.uuid4())
    file_path = os.path.join(d, fn + '.csv')

    start_date = dateutil.parser.parse('{}{}'.format(start_date, 'T00:00:00+08:00'))
    end_date = dateutil.parser.parse('{}{}'.format(end_date, 'T00:00:00+08:00'))

    if date_or_grade == 'True':
        with open(file_path, 'w', encoding='UTF-8') as fp:
            writer = csv.writer(fp)
            writer.writerow(['学号', '姓名', '邮件', '系所', '专业', '教师', '职工号', '开始时间', '结束时间', '是否完成'])
            mt_qs = MeetPlan.objects.filter(start_time__gte=start_date,
                                            end_time__lte=end_date)
            mto_qs = MeetPlanOrder.objects.filter(meet_plan__in=mt_qs)

            for mto in mto_qs:
                student = mto.student
                teacher = mto.meet_plan.teacher

                department = student.studentprofile.department.department if hasattr(student,
                                                                                     'studentprofile') else '未定义'
                major = student.studentprofile.major.major if hasattr(student, 'studentprofile') else '未定义'

                writer.writerow([student.identity_id, student.user_name, student.email, department, major,
                                 teacher.user_name, teacher.identity_id,
                                 mto.meet_plan.start_time, mto.meet_plan.end_time,
                                 '已完成' if mto.completed else '未完成'])

        file = MyFile.objects.create(upload_or_generate=False, app=app_name, user_id=user_id,
                                     remark='学生报表，时间从{}到{}'.format(start_date.strftime('%Y-%m-%d'),
                                                                   end_date.strftime('%Y-%m-%d')))
    else:
        with open(file_path, 'w', encoding='UTF-8') as fp:
            writer = csv.writer(fp)
            if detail == 'False':
                writer.writerow(['学号', '姓名', '邮件', '系所', '专业', '总预约次数', '总完成次数', '是否达到毕业要求', '备注'])
            else:
                writer.writerow(['学号', '姓名', '邮件', '系所', '专业', '教师', '职工号', '开始时间', '结束时间', '是否完成', '备注'])
            user_qs = User.objects.filter(is_teacher=False)
            user_qs1 = user_qs.filter(studentprofile__isnull=False).filter(studentprofile__grade_id__in=grades)
            user_qs2 = user_qs.filter(studentprofile__isnull=True)

            for user in user_qs1:
                department = user.studentprofile.department.department
                major = user.studentprofile.major.major

                mto = MeetPlanOrder.objects.filter(student=user)
                complete_num = mto.filter(completed=True).count()
                if detail == 'False':
                    writer.writerow([user.identity_id, user.user_name, user.email, department, major,
                                     mto.count(), complete_num,
                                     '已达到' if complete_num >= 8 else '未达到', '正常'])
                else:
                    mto_qs = mto.filter(meet_plan__start_time__gte=start_date,
                                        meet_plan__end_time__lte=end_date)
                    for mto in mto_qs:
                        writer.writerow([user.identity_id, user.user_name, user.email, department, major,
                                         mto.meet_plan.teacher.user_name, mto.meet_plan.teacher.identity_id,
                                         mto.meet_plan.start_time, mto.meet_plan.end_time,
                                         '已完成' if mto.completed else '未完成', '正常'])

            for user in user_qs2:
                department = '未定义'
                major = '未定义'
                mto = MeetPlanOrder.objects.filter(student=user)
                complete_num = mto.filter(completed=True).count()

                if detail == 'False':
                    writer.writerow([user.identity_id, user.user_name, user.email, department, major,
                                     mto.count(), complete_num,
                                     '已达到' if complete_num >= 8 else '未达到', '未设置学生信息无法判断年级'])
                else:
                    mto_qs = mto.filter(meet_plan__start_time__gte=start_date,
                                        meet_plan__end_time__lte=end_date)
                    for mto in mto_qs:
                        writer.writerow([user.identity_id, user.user_name, user.email, department, major,
                                         mto.meet_plan.teacher.user_name, mto.meet_plan.teacher.identity_id,
                                         mto.meet_plan.start_time, mto.meet_plan.end_time,
                                         '已完成' if mto.completed else '未完成', '未设置学生信息无法判断年级'])

        file = MyFile.objects.create(upload_or_generate=False, app=app_name, user_id=user_id,
                                     remark='{}级学生报表'.format(
                                         list(Grade.objects.filter(id__in=grades).values_list('grade', flat=True))))
    file.file.name = 'generate/' + fn + '.csv'
    file.save()
