from django.db import models
from django.db.models.query import QuerySet


class SoftDeletableQuerySet(QuerySet):
    def delete(self):
        # 1.如果 obj = Student.objects.get(id=6).delete() 删除, 则是走不到该函数的
        # 2.如果 obj = Student.objects.filter(id=6).delete() 删除, 则才会走到该函数
        # print('In the delete of the MySoftDeletableQuerySet, self = {0}'.format(self))
        self.update(is_delete=True)


class BaseModelManager(models.Manager):
    _queryset_class = SoftDeletableQuerySet

    def get_queryset(self):
        return super(BaseModelManager, self).get_queryset().filter(is_delete=False)


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

