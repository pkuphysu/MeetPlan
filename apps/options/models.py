from django.db import models

# Create your models here.
from db.base_model import BaseModel


class Option(BaseModel):
    app = models.CharField(verbose_name='应用', null=False, blank=False, max_length=32)
    name = models.CharField(verbose_name='属性', max_length=128, null=False, blank=False)
    value = models.TextField(verbose_name='值', null=False, blank=False)

    class Meta:
        unique_together = ('app', 'name')
        verbose_name = '设置'
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}-{}'.format(self.app, self.name)
