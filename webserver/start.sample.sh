#!/bin/bash

NAME="PKU_PHY_SU"    ##定义一个名字，用来作为进程名
DIR=/path/to/PKU_PHY_SU    ##你的项目路径
USER=root   ##启动该项目使用的用户
GROUP=root ##启动该项目使用的用户的组
WORKERS=3
BIND=unix:${DIR}/webserver/gunicorn.sock ##不需要你手工创建该文件！！！这是自动创建的用来通信的套接字
DJANGO_SETTINGS_MODULE=PKU_PHY_SU.settings.base ##改成你的项目
DJANGO_WSGI_MODULE=PKU_PHY_SU.wsgi ##改成你的项目
LOG_LEVEL=info

cd $DIR

export DJANGO_SETTINGS_MODULE=$DJANGO_SETTINGS_MODULE

##下面的/usr/local/bin/gunicorn  改成你自己安装后的gunicorn的路径，可以使用whereis gunicorn 来找到
exec /usr/local/bin/gunicorn ${DJANGO_WSGI_MODULE}:application \
  --name $NAME \
  --workers $WORKERS \
  --user=$USER \
  --group=$GROUP \
  --bind=$BIND \
  --log-level=$LOG_LEVEL \
  --log-file=${DIR}/webserver/gunicorn.log # 需要事先创建log文件