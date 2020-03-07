from django.db import models
from django.urls import reverse

from db.base_model import BaseModel
from django.conf import settings


# Create your models here.


class MeetPlan(BaseModel):
    """综合指导课模型类"""
    AllowOtherChoices = (
        (False, '不允许'),
        (True, '允许')
    )
    teacher = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, blank=True, null=True)
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

    def get_absolute_url(self):
        # return reverse('meet_plan:tea-add-plan', kwargs={'pk': self.pk})
        return reverse('meet_plan:tea-plan-detail', kwargs={'pk': self.id})

    def save(self, *args, **kwargs):
        if not self.allow_other:
            self.available_choice = 1
        super(MeetPlan, self).save(*args, **kwargs)


class MeetPlanOrder(BaseModel):
    """综合指导课预约情况"""
    meet_plan = models.ForeignKey(to=MeetPlan, on_delete=models.SET_NULL, blank=True, null=True)
    student = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, db_index=True, blank=True,
                                null=True)
    completed = models.BooleanField(default=False)
    message = models.TextField(verbose_name='意向谈话内容', null=True, blank=True)

    class Meta:
        verbose_name = '综合指导课预约'
        verbose_name_plural = verbose_name

    def save(self, *args, **kwargs):
        if self.id is not None:
            orig = MeetPlanOrder.objects.get(id=self.id)
            # TODO: 完善逻辑
            if orig.is_delete != self.is_delete:
                if self.is_delete:
                    self.meet_plan.available_choice += 1
                else:
                    self.meet_plan.available_choice -= 1
                self.meet_plan.save()
        else:
            self.meet_plan.available_choice -= 1
            self.meet_plan.save()
        super(MeetPlanOrder, self).save(*args, **kwargs)

    def get_absolute_url(self):
        # return reverse('meet_plan:tea-add-plan', kwargs={'pk': self.pk})
        return reverse('meet_plan:index')

    def __str__(self):
        return 'id:%d tea:%s' % (self.id, self.student.user_name)


class SemesterDateRange(BaseModel):
    """综合指导课学期起止时间"""
    start_date = models.DateField(verbose_name='学期开始时间')
    end_date = models.DateField(verbose_name='学期结束时间')

    class Meta:
        verbose_name = '学期起止时间'
        verbose_name_plural = verbose_name


class FeedBack(BaseModel):
    """综合指导课反馈"""
    FeedBackChoices = (
        (False, '待回应'),
        (True, '已回应')
    )
    teacher = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, blank=True, null=True)
    message = models.TextField(verbose_name='反馈')
    have_checked = models.BooleanField(verbose_name='是否已经处理', max_length=1,
                                       choices=FeedBackChoices,
                                       default=False)

    class Meta:
        verbose_name = '教师反馈'
        verbose_name_plural = verbose_name

    def get_absolute_url(self):
        # return reverse('meet_plan:tea-add-plan', kwargs={'pk': self.pk})
        return reverse('meet_plan:tea-feedback-list')
