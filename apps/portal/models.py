from django.db import models

from apps.filemanager.models import MyImg
# Create your models here.
from db.base_model import BaseModel


class FriendLink(BaseModel):
    name = models.CharField(verbose_name='名称', max_length=32)
    url = models.URLField(verbose_name='链接地址')
    image = models.ForeignKey(to=MyImg, on_delete=models.DO_NOTHING)
    description = models.TextField(verbose_name='描述', null=True, blank=True)

    class Meta:
        verbose_name = '友情链接'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.name


class UpdateRecord(BaseModel):
    time = models.DateField(verbose_name='时间', blank=False, null=False)
    author = models.CharField(verbose_name='作者', blank=False, null=False, max_length=512)
    url = models.URLField(verbose_name='链接地址', blank=False, null=False, max_length=512)
    info = models.TextField(verbose_name='更新内容', blank=False, null=False)

    class Meta:
        verbose_name = '更新列表'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.time) + self.author
