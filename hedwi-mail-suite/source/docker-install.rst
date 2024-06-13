.. _help-docker-install:

.. _docker-install:


docker部署 (推荐)
------------------------

0. 环境要求
=====================
最低2G内存 2核CPU  linux 64位系统

1. 配置DNS
===============================================

.. code-block:: bash

   在域名服务商或者DNS解析服务商配置域名A记录指向服务器公网IP

2. 安装docker, docker-compose, git
======================================


3. 安装
===============================================

.. code-block:: bash

  git clone https://github.com/Hedwi/hedwi-docker.git
  cd hedwi-docker/hedwi
  cp env.example .env
  #修改.env配置  修改所有备注『需要修改』的变量

  #修改nginx配置 nginx/conf.d/inbox.conf 中的域名  

  #https 443端口部分
  listen              443 ssl http2;
  listen              [::]:443 ssl http2;
  server_name         work.hedwi.com;  # 可以修改为任意子域名

  #http 80端口部分
  listen      80;
  listen      [::]:80;
  server_name work.hedwi.com;  #可以修改为任意子域名


  make  #运行起来之后 http 80端口可以访问 应该可以正常访问 http://配置的域名


`配置文件示例 </mail-suite/env-example.html>`_


4. 自动生成免费ssl证书
===============================================

.. code-block:: bash

  # 服务内置Let’s Encrypt HTTP-01 验证服务接口，运行此命令会请求/acme/create 接口创建证书
  # 获取证书成功之后会将证书写入 inbox/certs/cert.key  inbox/certs/cert.crt
  # 此命令会使用生成的免费证书替换掉 nginx/certs/cert.key nginx/certs/cert.crt
  make cert

5. 重启服务
===============================================

.. code-block:: bash

  make stop #停掉所有运行的docker容器
  make #启动docker容器

6. 使用https访问 登录
===============================================

访问 https://你的域名/hello 页面, 会自动根据配置文件中的邮箱和密码创建管理员账号（用于配置系统、添加用户、添加团队等）和一个普通用户账号admin@你的域名（用于使用系统服务），
