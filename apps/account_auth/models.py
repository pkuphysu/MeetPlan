from django.db import models
from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin, BaseUserManager
from db.base_model import BaseModel, SoftDeletableQuerySet
from django.core.validators import RegexValidator


# Create your models here.


class UserManager(BaseUserManager):
    _queryset_class = SoftDeletableQuerySet

    def get_queryset(self, is_delete=False):
        return super(UserManager, self).get_queryset().filter(is_delete=is_delete)

    def _create_user(self, identity_id, user_name, **kwargs):
        if 'email' not in kwargs:
            kwargs.update({'email': '%s@pku.edu.cn' % identity_id})
        else:
            kwargs['email'] = self.normalize_email(kwargs['email'])
        user = self.model(identity_id=identity_id, user_name=user_name, **kwargs)
        user.set_unusable_password()
        user.save()
        return user

    def _create_stu(self, identity_id, user_name, **kwargs):
        kwargs['is_superuser'] = False
        kwargs['is_active'] = False
        return self._create_user(identity_id=identity_id, user_name=user_name, **kwargs)

    def _create_tea(self, identity_id, user_name, **kwargs):
        kwargs['is_superuser'] = False
        kwargs['is_teacher'] = True
        kwargs['is_active'] = False
        return self._create_user(identity_id=identity_id, user_name=user_name, **kwargs)

    def create_user(self, identity_id, user_name, is_teacher, **kwargs):
        if is_teacher:
            return self._create_tea(identity_id, user_name, **kwargs)
        else:
            return self._create_stu(identity_id, user_name, **kwargs)

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
    identity_id = models.CharField(max_length=16, unique=True, verbose_name='职工号\\学号', db_index=True)
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


class BaseProfile(BaseModel):
    GenderChoices = (
        (False, '男'),
        (True, '女'),
    )
    user = models.OneToOneField(verbose_name='用户', to=User, on_delete=models.DO_NOTHING, db_index=True)
    gender = models.BooleanField(default=False, choices=GenderChoices, verbose_name='性别')

    birth = models.DateField(null=True, blank=True, verbose_name='生日')
    from apps.filemanager.models import MyImg
    head_picture = models.ForeignKey(to=MyImg, on_delete=models.SET_NULL, null=True, blank=True, default=None)

    class Meta:
        verbose_name = "基本信息"
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.user.user_name


class Department(BaseModel, models.Model):
    department = models.CharField(verbose_name='系所', unique=True, max_length=100)

    class Meta:
        verbose_name = '系所\\办公室'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.department


class Major(BaseModel):
    department = models.ForeignKey(to=Department, on_delete=models.DO_NOTHING, verbose_name='系所')
    major = models.CharField(verbose_name='专业', max_length=100)

    class Meta:
        verbose_name = '专业'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.major


class Grade(BaseModel):
    grade = models.CharField(verbose_name='年级', max_length=4)

    class Meta:
        verbose_name = '年级'
        verbose_name_plural = verbose_name

    def __str__(self):
        return '{}级'.format(self.grade)


class StudentProfile(BaseModel):
    GRADUATE_CHOICES = (
        (True, '研究生'),
        (False, '本科生')
    )
    user = models.OneToOneField(verbose_name='学生', to=User, on_delete=models.DO_NOTHING, db_index=True)

    is_graduate = models.BooleanField(verbose_name='研究生\\本科生', default=False, choices=GRADUATE_CHOICES)

    phone_regex = RegexValidator(regex=r'^\+?1?\d{9,15}$',message="号码格式错误。")
    phone_number = models.CharField(validators=[phone_regex], max_length=17, blank=True)

    department = models.ForeignKey(to=Department, on_delete=models.DO_NOTHING, verbose_name='系所')
    major = models.ForeignKey(to=Major, on_delete=models.DO_NOTHING, verbose_name='专业')
    grade = models.ForeignKey(to=Grade, on_delete=models.DO_NOTHING, verbose_name='年级')

    dorm = models.CharField(verbose_name='宿舍', max_length=32)

    class Meta:
        verbose_name = '学生详细信息'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.user.user_name


class TeacherProfile(BaseModel):
    user = models.OneToOneField(verbose_name='教师', to=User, on_delete=models.DO_NOTHING, db_index=True)

    phone_regex = RegexValidator(regex=r'^\+?1?\d{9,15}$', message="号码格式错误。座机请加区号")
    phone_number = models.CharField(validators=[phone_regex], max_length=17, blank=True)

    department = models.ForeignKey(to=Department, on_delete=models.DO_NOTHING, verbose_name='系所')
    office = models.CharField(verbose_name='办公室', max_length=100)

    introduce = models.TextField(verbose_name='个人简介', blank=True, null=True)

    class Meta:
        verbose_name = '教职工详细信息'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.user.user_name
