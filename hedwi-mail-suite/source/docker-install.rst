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

例如

..  csv-table:: 
    :header: "主机名", "记录类型", "记录值"
    :widths: 35, 35, 30

    "work","A","x.x.x.x"


- 主机名可以是任意二级域名前缀，这里使用work
- example.com为你的域名
- x.x.x.x 为服务器的IP地址
- 网页访问地址为 https://work.example.com

2. 安装docker, docker-compose, git
======================================

.. code-block:: bash

  # 已经安装的话可以跳过

  mkdir pkgs && cd pkgs
  wget 'https://download.docker.com/linux/static/stable/x86_64/docker-26.1.4.tgz'
  tar xf docker-26.1.4.tgz;
  chmod +x docker/*
  sudo cp docker/* /usr/local/bin/

  curl -SL https://github.com/docker/compose/releases/download/v2.27.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

  yum install git # centos
  apt install git  # ubuntu/debian


3. 安装
===============================================

.. code-block:: bash

  git clone https://github.com/Hedwi/hedwi-docker.git
  cd hedwi-docker/hedwi
  cp env.example .env
  #修改.env配置  修改所有备注『需要修改』的变量

  make  #运行起来之后 http 80端口可以访问 应该可以正常访问 http://配置的域名

  # 服务会自动根据配置文件中的邮箱和密码创建管理员账号（用于配置系统、添加用户、添加团队等）和一个普通用户账号admin@你的域名（用于使用系统服务）


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

  make down #停掉所有运行的docker容器
  make #启动docker容器

6. 使用https访问 登录
===============================================

访问 https://work.example.com/login?admin=true (work.example.com为使用的域名), 使用管理员账号登录(配置的邮箱和密码)
访问 https://work.example.com/hello 页面,  继续添加用户、团队等等。
