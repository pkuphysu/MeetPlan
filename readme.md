
### 安装软件
```
sudo yum install supervisor
curl https://pyenv.run | bash
```

### 进入虚拟环境
```
pyenv install 3.8.2
pyenv virtualenv 3.8.2 PKU_PHY_SU
pyenv activate PKU_PHY_SU
pip install -r requirements.txt
```
### 生成运行脚本
```
cd webserver
cp start.sample.sh start.sh
vim start.sh
```
更改`start.sh`里的`DIR`为项目的实际路径后运行`chmod u+x start.sh`
```
cp supervisor.sample.conf supervisor.conf
vim supervisor.conf
```
更改`supervisor.conf`里的`command`，然后保存退出做软连接
```
sudo ln -s /your/path/supervisor.conf   /etc/supervisor/conf.d/
```

### 使用supervisor运行django

```
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl status djangopro
```

