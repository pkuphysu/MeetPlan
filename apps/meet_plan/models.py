from django.conf import settings
from django.db import models

from db.base_model import BaseModel


# Create your models here.


class MeetPlan(BaseModel):
    """综合指导课模型类"""
    AllowOtherChoices = (
        (False, '不允许'),
        (True, '允许')
    )
    teacher = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING)
    place = models.CharField(verbose_name="谈话地点", max_length=128, help_text="谈话地点")
    start_time = models.DateTimeField(verbose_name='开始时间', help_text="谈话开始时间")
    end_time = models.DateTimeField(verbose_name='结束时间', help_text="谈话结束时间")
    allow_other = models.BooleanField(verbose_name='允许多人', max_length=1,
                                      choices=AllowOtherChoices,
                                      help_text="是否可两人同时进行综合指导课",
                                      default=True)
    message = models.TextField(verbose_name='备注', blank=True, null=True)
    available_choice = models.SmallIntegerField(verbose_name='还有多少名额', default=2)

    class Meta:
        verbose_name = '综合指导课'
        verbose_name_plural = verbose_name

    def __str__(self):
        return 'id:%d tea:%s' % (self.id, self.teacher.user_name)

    def save(self, *args, **kwargs):
        if self.id is None and not self.allow_other:
            self.available_choice = 1
        super(MeetPlan, self).save(*args, **kwargs)

    def delete(self, using=None, keep_parents=False, soft=True):
        self.meetplanorder_set.all().delete()
        return super().delete(using, keep_parents, soft)


class MeetPlanOrder(BaseModel):
    """综合指导课预约情况"""
    meet_plan = models.ForeignKey(to=MeetPlan, on_delete=models.DO_NOTHING)
    student = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING, db_index=True)
    completed = models.BooleanField(default=False)
    message = models.TextField(verbose_name='意向谈话内容', null=True, blank=True)

    class Meta:
        verbose_name = '综合指导课预约'
        verbose_name_plural = verbose_name

    def save(self, *args, **kwargs):
        if self.id is None:
            self.meet_plan.available_choice -= 1
            self.meet_plan.save()
        super(MeetPlanOrder, self).save(*args, **kwargs)

    def delete(self, using=None, keep_parents=False, soft=True):
        self.meet_plan.available_choice += 1
        self.meet_plan.save()
        super(MeetPlanOrder, self).delete(using, keep_parents, soft)

    def __str__(self):
        return 'id:%d tea:%s' % (self.id, self.student.user_name)


class FeedBack(BaseModel):
    """综合指导课反馈"""
    FeedBackChoices = (
        (False, '待回应'),
        (True, '已回应')
    )
    teacher = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING)
    message = models.TextField(verbose_name='反馈')
    have_checked = models.BooleanField(verbose_name='是否已经处理', max_length=1,
                                       choices=FeedBackChoices,
                                       default=False)

    class Meta:
        verbose_name = '教师反馈'
        verbose_name_plural = verbose_name
