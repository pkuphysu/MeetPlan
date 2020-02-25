from django.db import models


# Create your models here.


class FriendLink(models.Model):
    link = models.URLField(verbose_name='链接地址')
