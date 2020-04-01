from abc import ABC

from django.db import models
from django.db.models.query import QuerySet
from django.db.models import Func


class Convert(Func, ABC):
    def __init__(self, expression, transcoding_name, **extra):
        super(Convert, self).__init__(
            expression, transcoding_name=transcoding_name, **extra)

    def as_mysql(self, compiler, connection):
        self.function = 'CONVERT'
        self.template = '%(function)s(%(expressions)s USING %(transcoding_name)s)'
        return super(Convert, self).as_sql(compiler, connection)


class SoftDeletableQuerySet(QuerySet):
    def delete(self):
        # 1.如果 obj = Student.objects.get(id=6).delete() 删除, 则是走不到该函数的
        # 2.如果 obj = Student.objects.filter(id=6).delete() 删除, 则才会走到该函数
        # print('In the delete of the MySoftDeletableQuerySet, self = {0}'.format(self))
        self.update(is_delete=True)


class BaseModelManager(models.Manager):
    _queryset_class = SoftDeletableQuerySet

    def get_queryset(self, is_delete=False):
        return super(BaseModelManager, self).get_queryset().filter(is_delete=is_delete)


class BaseModel(models.Model):
    """模型抽象基类"""
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')
    update_time = models.DateTimeField(auto_now=True, verbose_name='更新时间')
    is_delete = models.BooleanField(default=False, verbose_name='删除标记')

    objects = BaseModelManager()

    class Meta:
        abstract = True

    def delete(self, using=None, keep_parents=False, soft=True):
        """实现软删除"""
        # 1.如果 obj = Student.objects.get(id=6).delete() 删除, 则才会走到该函数
        # 2.如果 obj = Student.objects.filter(id=6).delete() 删除, 则是走不到该函数的
        if soft:
            self.is_delete = True
            self.save(using=using)
        else:
            return super().delete(using=using, keep_parents=keep_parents)
