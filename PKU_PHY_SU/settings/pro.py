import configparser


# 读取机密信息
config = configparser.RawConfigParser()
config.read(filenames='PKU_PHY_SU/secret_config.ini', encoding='UTF-8')

APPID = config.get('IAAA', 'APPID')
APPKEY = config.get('IAAA', 'APPKEY')

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = False

ALLOWED_HOSTS = ['*']

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = config.get('Django', 'SECRET_KEY')


# Database
# https://docs.djangoproject.com/en/2.2/ref/settings/#databases
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': config.get('DataBase', 'NAME'),
        'USER': config.get('DataBase', 'USER'),
        'PASSWORD': config.get('DataBase', 'PASSWORD'),
        'HOST': config.get('DataBase', 'HOST'),
        'PORT': config.getint('DataBase', 'PORT'),
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
        "LOCATION": "redis://%s:%s/pku_phy_1" % (config.get('Redis', 'HOST'), config.get('Redis', 'PORT')),
        "OPTIONS": {
            "CLIENT_CLASS": "django_redis.client.DefaultClient",
        }
    }
}