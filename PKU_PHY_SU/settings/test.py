import configparser


# 读取机密信息
config = configparser.RawConfigParser()
config.read(filenames='PKU_PHY_SU/secret_config.ini', encoding='UTF-8')

APPID = config.get('IAAA', 'APPID')
APPKEY = config.get('IAAA', 'APPKEY')

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

ALLOWED_HOSTS = ['*']

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = '$l!yy=1_6r%fq8h+@n^p0gteowo&g=e*bct93&tz6a)giohgfp'


# Database
# https://docs.djangoproject.com/en/2.2/ref/settings/#databases
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'pku_phy_su_test',
        'USER': 'pku_phy_su_test',
        'PASSWORD': 'pku_phy_su_test',
        'HOST': 'localhost',
        'PORT': 3306,
    }
}


# 发送邮件
EMAIL_BACKEND = 'django.core.mail.backends.smtp.EmailBackend'
# 使用 SSL 连接
EMAIL_USE_SSL = True
# SMTP 服务地址和端口
EMAIL_HOST = config.get('Email', 'HOST')
EMAIL_PORT = config.getint('Email', 'PORT')
# 发送邮件的邮箱
EMAIL_HOST_USER = config.get('Email', 'USER')
EMAIL_HOST_PASSWORD = config.get('Email', 'PASSWORD')
EMAIL_FROM = config.get('Email', 'FROM')


# Redis 缓存配置
CACHES = {
    "default": {
        "BACKEND": "django_redis.cache.RedisCache",
        "LOCATION": "redis://127.0.0.1:6379/3",
        "OPTIONS": {
            "CLIENT_CLASS": "django_redis.client.DefaultClient",
        }
    }
}

# Broker配置，使用Redis作为消息中间件
CELERY_BROKER_URL = 'redis://127.0.0.1:6379/4'


print('TEST')
