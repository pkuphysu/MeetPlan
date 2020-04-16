# Generated by Django 2.2.10 on 2020-04-01 11:37

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name='MeetPlan',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('place', models.CharField(help_text='谈话地点', max_length=128, verbose_name='谈话地点')),
                ('start_time', models.DateTimeField(help_text='谈话开始时间', verbose_name='开始时间')),
                ('end_time', models.DateTimeField(help_text='谈话结束时间', verbose_name='结束时间')),
                ('allow_other', models.BooleanField(choices=[(False, '不允许'), (True, '允许')], default=True, help_text='是否可两人同时进行综合指导课', max_length=1, verbose_name='允许多人')),
                ('message', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('available_choice', models.SmallIntegerField(default=2, verbose_name='还有多少名额')),
                ('teacher', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to=settings.AUTH_USER_MODEL)),
            ],
            options={
                'verbose_name': '综合指导课',
                'verbose_name_plural': '综合指导课',
            },
        ),
        migrations.CreateModel(
            name='MeetPlanOrder',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('completed', models.BooleanField(default=False)),
                ('message', models.TextField(blank=True, null=True, verbose_name='意向谈话内容')),
                ('meet_plan', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='meet_plan.MeetPlan')),
                ('student', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to=settings.AUTH_USER_MODEL)),
            ],
            options={
                'verbose_name': '综合指导课预约',
                'verbose_name_plural': '综合指导课预约',
            },
        ),
        migrations.CreateModel(
            name='FeedBack',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='更新时间')),
                ('is_delete', models.BooleanField(default=False, verbose_name='删除标记')),
                ('message', models.TextField(verbose_name='反馈')),
                ('have_checked', models.BooleanField(choices=[(False, '待回应'), (True, '已回应')], default=False, max_length=1, verbose_name='是否已经处理')),
                ('teacher', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to=settings.AUTH_USER_MODEL)),
            ],
            options={
                'verbose_name': '教师反馈',
                'verbose_name_plural': '教师反馈',
            },
        ),
    ]
