from django.db import models
from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin, BaseUserManager
from db.base_model import BaseModel
from phonenumber_field.modelfields import PhoneNumberField


# Create your models here.


class UserManager(BaseUserManager):
    def _create_user(self, identity_id, user_name, **kwargs):
        if 'email' not in kwargs:
            kwargs.update({'email': '%s@pku.edu.cn' % identity_id})
        else:
            kwargs['email'] = self.normalize_email(kwargs['email'])
        user = self.model(identity_id=identity_id, user_name=user_name, **kwargs)
        user.set_unusable_password()
        user.save()
        return user

    def create_stu(self, identity_id, user_name, **kwargs):
        kwargs['is_superuser'] = False
        kwargs['is_active'] = False
        return self._create_user(identity_id=identity_id, user_name=user_name, **kwargs)

    def create_tea(self, identity_id, user_name, **kwargs):
        kwargs['is_superuser'] = False
        kwargs['is_teacher'] = True
        kwargs['is_active'] = False
        return self._create_user(identity_id=identity_id, user_name=user_name, **kwargs)

    def create_user(self, identity_id, user_name, is_teacher, **kwargs):
        if is_teacher:
            return self.create_tea(identity_id, user_name, **kwargs)
        else:
            return self.create_stu(identity_id, user_name, **kwargs)

    def create_superuser(self, identity_id, user_name, **kwargs):
        kwargs['is_superuser'] = True
        kwargs['is_active'] = True
        kwargs['is_teacher'] = True
        kwargs['is_admin'] = True
        return self._create_user(identity_id=identity_id, user_name=user_name, **kwargs)


class User(AbstractBaseUser, BaseModel, PermissionsMixin):
    """
    User模型
    """
    identity_id = models.CharField(max_length=11, unique=True, verbose_name='职工号\\学号', db_index=True, primary_key=True)
    user_name = models.CharField(max_length=100, null=False, blank=False, verbose_name='姓名')
    email = models.EmailField(null=True, blank=True, verbose_name='电子邮件')
    is_active = models.BooleanField(default=False, verbose_name='是否激活')
    is_teacher = models.BooleanField(default=False, verbose_name='是否为教职工')
    is_admin = models.BooleanField(default=False, verbose_name='是否为管理员, 管理员可登陆cmsadmin管理页面')
    # is_staff = models.BooleanField(default=False)
    USERNAME_FIELD = 'identity_id'
    REQUIRED_FIELDS = ['user_name']
    EMAIL_FIELD = 'email'

    objects = UserManager()

    class Meta:
        verbose_name = "用户"
        verbose_name_plural = verbose_name

    def get_full_name(self):
        return self.user_name + self.identity_id

    def get_short_name(self):
        return self.user_name

    def save(self, *args, **kwargs):
        if self.email is None:
            self.email = '%s@pku.edu.cn' % self.identity_id
        super(User, self).save(*args, **kwargs)

    def __str__(self):
        return self.get_full_name()

    @property
    def is_staff(self):
        """
        判断是否可以访问django 默认 admin 管理站点, 这里和is_superuser相同
        :return:
        """
        return self.is_superuser


class UserProfile(BaseModel, models.Model):
    GenderChoices = (
        (0, '未知'),
        (1, '男'),
        (2, '女'),
    )
    user = models.OneToOneField(verbose_name='用户', to=User, on_delete=models.DO_NOTHING, primary_key=True)
    gender = models.SmallIntegerField(default=0, choices=GenderChoices, verbose_name='性别')
    telephone = PhoneNumberField(region='CN', null=False, blank=False, verbose_name='联系方式')

    birth = models.DateField(null=True, blank=True, verbose_name='生日')
    from apps.filemanager.models import Img
    head_picture = models.ForeignKey(to=Img, on_delete=models.SET_NULL, null=True, blank=True, default=None)

    class Meta:
        verbose_name = "用户详细信息"
        verbose_name_plural = verbose_name
