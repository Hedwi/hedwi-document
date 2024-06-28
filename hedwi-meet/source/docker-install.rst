.. _help-docker-install:

.. _docker-install:


docker部署 (推荐)
------------------------

0. 环境要求
=====================
最低1G内存 1核CPU  linux 64位系统

1. 配置DNS
===============================================

.. code-block:: bash

   在域名服务商或者DNS解析服务商配置域名A记录指向服务器公网IP


2. 安装docker, git 
======================================


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

