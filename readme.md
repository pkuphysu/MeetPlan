## 部署网站
### 安装软件
```
sudo yum install supervisor
curl https://pyenv.run | bash
```

### 创建虚拟环境
```
pyenv install 3.8.2
pyenv global 3.8.2
pip install pipenv

cd /path/to/project
pipenv install
```
### 生成配置文件
```
cp webserver/uwsgi.sample.ini ./uwsgi.ini
```
更改`uwsgi.ini`里的`chdir`为项目的实际路径，合理修改其它选项

### 测试网站
启动
```
pipenv shell
uwsgi --ini uwsgi.ini
```
重启
```
uwsgi --reload uwsgi.pid
```
停止
```
uwsgi --stop uwsgi.pid
```

### 使用supervisor管理网站
