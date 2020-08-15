from __future__ import absolute_import, unicode_literals

from celery import shared_task
from django.conf import settings
from django.contrib.auth import get_user_model
from django.template import loader

from MeetPlan.tools.celery import TransactionAwareTask, my_send_mail
from apps.meet_plan.models import MeetPlanOrder, FeedBack


@shared_task(base=TransactionAwareTask, bind=True)
def send_meetplan_order_create_email(self, meetplanorder_id):
    domain = settings.SITE_URL
    order = MeetPlanOrder.objects.get(id=meetplanorder_id)
    meetplan = order.meet_plan
    teacher = meetplan.teacher
    tea_email = [teacher.email]
    student = order.student
    stu_message = order.message
    stu_email = [student.email]

    # 构造邮件信息
    from_email = settings.EMAIL_FROM
    # 教师邮件
    tea_subject = '物理学院综合指导课新预约'
    tea_html = loader.render_to_string(
        'email/meetplan/tea_meetplan_order_create.html',
        {
            'domain': domain,
            'user_name': teacher.user_name,
            'start_time': meetplan.start_time,
            'end_time': meetplan.end_time,
            'place': meetplan.place,
            'tea_message': meetplan.message,

            'stu_name': student.user_name,
            'stu_email': student.email,
            'stu_message': stu_message,
        }
    )
    my_send_mail.delay(tea_subject, tea_html, from_email, tea_email)

    # 学生邮件
    stu_subject = '物理学院综合指导课新预约'
    stu_html = loader.render_to_string(
        'email/meetplan/stu_meetplan_order_create.html',
        {
            'domain': domain,
            'user_name': student.user_name,
            'stu_message': stu_message,

            'tea_name': teacher.user_name,
            'start_time': meetplan.start_time,
            'end_time': meetplan.end_time,
            'place': meetplan.place,
            'tea_email': teacher.email,
            'tea_message': meetplan.message,
        }
    )
    my_send_mail.delay(stu_subject, stu_html, from_email, stu_email)


@shared_task(base=TransactionAwareTask, bind=True)
def send_meetplan_order_update_email(self, meetplanorder_id, is_delete):
    if is_delete:
        order = MeetPlanOrder.objects.get_queryset(is_delete=True).filter(id=meetplanorder_id)[0]
    else:
        order = MeetPlanOrder.objects.get(id=meetplanorder_id)
    domain = settings.SITE_URL
    meetplan = order.meet_plan
    student = order.student
    stu_email = [student.email]
    teacher = meetplan.teacher
    # 构造邮件信息
    from_email = settings.EMAIL_FROM
    stu_subject = '物理学院综合指导课预约状态更新通知'
    stu_html = loader.render_to_string(
        'email/meetplan/stu_meetplan_order_update.html',
        {
            'domain': domain,
            'user_name': student.user_name,
            'tea_name': teacher.user_name,
            'start_time': meetplan.start_time,
            'end_time': meetplan.end_time,
            'place': meetplan.place,
            'tea_email': teacher.email,
            'tea_message': meetplan.message,
            'status': '删除' if order.is_delete else '已确认',
            'success': student.meetplanorder_set.filter(is_delete=False, completed=True).count()
        }
    )
    my_send_mail.delay(stu_subject, stu_html, from_email, stu_email)


@shared_task(base=TransactionAwareTask, bind=True)
def send_meetplan_feedback_create_email(self, feedback_id):
    domain = settings.SITE_URL
    feedback = FeedBack.objects.get(id=feedback_id)
    teacher = feedback.teacher
    message = feedback.message

    user_model = get_user_model()
    admins = user_model.objects.filter(is_active=True, is_admin=True)
    admin_email = list(admins.values_list('email', flat=True))

    # 构造邮件信息
    from_email = settings.EMAIL_FROM
    feedback_subject = '物理学院综合指导课新教师反馈'
    feedback_html = loader.render_to_string(
        'email/meetplan/tea_meetplan_feedback_create.html',
        {
            'domain': domain,
            'user_name': '管理员',
            'tea_name': teacher.user_name,
            'tea_email': teacher.email,
            'message': message,
        }
    )
    my_send_mail.delay(feedback_subject, feedback_html, from_email, admin_email)


@shared_task(base=TransactionAwareTask, bind=True)
def send_meetplan_feedback_update_email(self, feedback_id):
    domain = settings.SITE_URL
    feedback = FeedBack.objects.get(id=feedback_id)
    teacher = feedback.teacher
    message = feedback.message

    user_model = get_user_model()
    admins = user_model.objects.filter(is_active=True, is_admin=True)
    admins = list(admins.values_list('user_name', 'email'))

    # 构造邮件信息
    from_email = settings.EMAIL_FROM
    feedback_subject = '物理学院综合指导课教师反馈状态更新'
    feedback_html = loader.render_to_string(
        'email/meetplan/tea_meetplan_feedback_update.html',
        {
            'domain': domain,
            'user_name': teacher.user_name,
            'message': message,
            'status': '已回应' if feedback.have_checked else '待处理',
            'admins': admins
        }
    )
    my_send_mail.delay(feedback_subject, feedback_html, from_email, [teacher.email])


@shared_task
def send_meetplan_alert_everyday():
    domain = settings.SITE_URL
    from django.utils import timezone
    import datetime
    mto_qs = MeetPlanOrder.objects.filter(meet_plan__start_time__gte=timezone.now(),
                                          meet_plan__end_time__lte=timezone.now() + datetime.timedelta(days=1))
    subject = '物理学院综合指导课预约提醒'
    from_email = settings.EMAIL_FROM
    for mto in mto_qs:
        meetplan = mto.meet_plan
        teacher = meetplan.teacher
        student = mto.student

        html = loader.render_to_string(
            'email/meetplan/stu_meetplan_order_alert.html',
            {
                'domain': domain,
                'user_name': student.user_name,
                'stu_message': mto.message,

                'tea_name': teacher.user_name,
                'start_time': meetplan.start_time,
                'end_time': meetplan.end_time,
                'place': meetplan.place,
                'tea_email': teacher.email,
                'tea_message': meetplan.message,
            }
        )
        my_send_mail.delay(subject, html, from_email, [student.email])
