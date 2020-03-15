from django.db import models

# Create your models here.
from db.base_model import BaseModel
from apps.filemanager.models import MyImg


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

