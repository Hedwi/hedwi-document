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
  #修改.env配置
  make

`配置文件示例 </mai-suite/hedwi_env.html>`_


4. 自动生成免费ssl证书
===============================================

.. code-block:: bash

  make cert


服务内置了证书申请服务，获取成功之后会将证书写入 inbox/certs/cert.key  inbox/certs/cert.crt 
使用生成的免费证书替换掉 nginx/certs/cert.key nginx/certs/cert.crt
后面会增加make步骤 自动执行此操作


5. 重启nginx  
===============================================

.. code-block:: bash

  make