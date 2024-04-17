.. _help-docker-cluster-install:

.. _docker-cluster-install:


使用 docker 进行集群部署
------------------------

docker 安装
=====================

.. code-block:: bash

	todo

	curl 'https://hedwi.com/cluster' |sh

启动
============

.. code-block:: bash

    docker compose up -d

- 防火墙配置
    对外开放25、465、993、443端口即可 本邮箱服务不需要nginx 如果使用nginx则需要自行修改配置文件中的端口
- 创建管理员帐号
    访问https://127.0.0.1:40008 (此处可以是根域名或者任意子域名，只要配置好DNS记录就可以)。注册管理员帐号 第一个帐号就是管理员帐号。使用管理员帐号添加域名，配置DNS记录。