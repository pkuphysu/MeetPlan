from django.conf import settings
from django.db import models

from db.base_model import BaseModel


# Create your models here.


class MyFile(BaseModel):
    """文件管理类"""
    UPLOAD_OR_GENERATE_CHOICES = (
        (True, '上传'),
        (False, '生成')
    )
    user = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING, verbose_name='所有者')
    file = models.FileField(verbose_name='文件')
    app = models.CharField(verbose_name='应用名', max_length=32, db_index=True)
    upload_or_generate = models.BooleanField(verbose_name='上传还是生成', choices=UPLOAD_OR_GENERATE_CHOICES,
                                             default=True)
    remark = models.TextField(verbose_name='备注', null=True, blank=True)

    class Meta:
        verbose_name = "文件管理"
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}-{}'.format(self.user, self.file.url)


class MyImg(BaseModel):
    """图片管理类"""
    UPLOAD_OR_GENERATE_CHOICES = (
        (True, '上传'),
        (False, '生成')
    )
    user = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING, verbose_name='所有者')
    img = models.ImageField(verbose_name='图片')
    app = models.CharField(verbose_name='应用名', max_length=32, db_index=True)
    upload_or_generate = models.BooleanField(verbose_name='上传还是生成', choices=UPLOAD_OR_GENERATE_CHOICES,
                                             default=True)
    remark = models.TextField(verbose_name='备注', null=True, blank=True)

    class Meta:
        verbose_name = "图片管理"
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}-{}'.format(self.user, self.img.url)
