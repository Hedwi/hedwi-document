.. _help-docker-install:

.. _docker-install:


docker部署
------------------------

0. 环境要求
=====================
最低1G内存 1核CPU  linux 64位系统

1. 配置DNS
===============================================

.. code-block:: bash

   在域名服务商或者DNS解析服务商配置域名A记录指向服务器公网IP


例如

..  csv-table:: 
    :header: "主机名", "记录类型", "记录值"
    :widths: 35, 35, 30

    "meet","A","x.x.x.x"


- 主机名可以是任意二级域名前缀，这里使用meet
- example.com为你的域名
- x.x.x.x 为服务器的IP地址
- 网页访问地址为 https://meet.example.com/meet  


2. 安装docker, git 
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
  cd hedwi-docker/meet
  cp env.example .env
  #修改.env配置

  make

`配置文件示例 </meet/meet_env.html>`_


4. 自动生成免费ssl证书
===============================================

.. code-block:: bash

  make cert


5. 重启nginx  
===============================================

.. code-block:: bash

  make down && make

