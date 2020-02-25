from celery import Celery
from django.core.mail import send_mail
from django.conf import settings
import os

# 为celery设置环境变量
os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'PKU_PHY_SU.settings')


app = Celery('pku_phy_celery.tasks')

app.conf.update(
    # 配置broker, 这里我们用redis作为broker
    BROKER_URL='redis://172.16.179.142:6379/pku_phy_celery',
    # 使用项目数据库存储任务执行结果
    CELERY_RESULT_BACKEND='django-db',
    # 配置定时器模块，定时器信息存储在数据库中
    CELERYBEAT_SCHEDULER='django_celery_beat.schedulers.DatabaseScheduler',
)

# 设置app自动加载任务
# 从已经安装的app中查找任务
app.autodiscover_tasks(settings.INSTALLED_APPS)

@app.tasks
def send_register_active_email(to_email, username, token):
    '''发送激活邮件'''
    # 组织邮件信息
    subject = '北京大学物理学院学生会欢迎信息'
    message = ''
    sender = settings.EMAIL_FROM
    receiver = [to_email]
    html_message = '<h1>%s, 欢迎您激活北京大学物理学院学生会服务</h1>请点击下面链接激活您的账户<br/><a href="http://127.0.0.1:8000/user/active/%s">http://127.0.0.1:8000/user/active/%s</a>' % (username, token, token)

    send_mail(subject, message, sender, receiver, html_message=html_message)