from django.db import models
from db.base_model import BaseModel
from django.conf import settings
# Create your models here.


class File(BaseModel):
    """文件管理类"""
    UPLOAD_OR_GENERATE_CHOICES = (
        (True, '上传'),
        (False, '生成')
    )
    user = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, blank=True, null=True,
                             verbose_name='所有者')
    file = models.FileField(verbose_name='文件')
    upload_or_generate = models.BooleanField(verbose_name='上传还是生成', choices=UPLOAD_OR_GENERATE_CHOICES,
                                             default=True)

    class Meta:
        verbose_name = "文件管理"
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}-{}'.format(self.user, self.file.url)


class Img(BaseModel):
    """图片管理类"""
    UPLOAD_OR_GENERATE_CHOICES = (
        (True, '上传'),
        (False, '生成')
    )
    user = models.ForeignKey(to=settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, blank=True, null=True,
                             verbose_name='所有者')
    img = models.ImageField(verbose_name='图片')
    upload_or_generate = models.BooleanField(verbose_name='上传还是生成', choices=UPLOAD_OR_GENERATE_CHOICES,
                                             default=True)

    class Meta:
        verbose_name = "图片管理"
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}-{}'.format(self.user, self.img.url)