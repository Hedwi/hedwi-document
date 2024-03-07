.. _help-manual-install:

.. _manual-install:


手动部署 
------------------------

0. 环境要求
=====================
最低1G内存 1核CPU  linux 64位系统


1. 安装依赖  
======================================

.. code-block:: bash

    centos / redhat

        yum install -y nginx supervisor redis

    debian / ubuntu 

        apt install -y nginx supervisor redis


2. 安装 meet server
===============================================


.. code-block:: bash

    export HEDWI_ROOT=/data/www   


    mkdir -p $HEDWI_ROOT/{hedwi-meetserver,redis}/logs;
    cd $HEDWI_ROOT ;
    wget 'https://download.hedwi.com/mail/linux/amd64/hedwi-meetserver.tar.gz'
    tar xf hedwi-meetserver.tar.gz;

    wget 'https://download.hedwi.com/mail/hedwi-meetserver-conf.tar.gz';
    tar xf hedwi-meetserver-conf.tar.gz;
    cp hedwi-meetserver-conf/hedwi-meetserver.conf .;
    cp hedwi-meetserver-conf/config.yaml.example hedwi-meetserver/config.yaml;
    按照示例修改配置文件中的域名

    supervisord -c /data/www/hedwi-meetserver.conf;


`配置文件示例 </meet/meet_conf.html>`_



3. nginx和ssl证书配置
===============================================


.. code-block:: json

    mkdir -p $HEDWI_ROOT/ssls/;
    cd $HEDWI_ROOT/ssls;
    mkdir example.com; //域名证书文件夹 此处为对应域名
    cd !$; 
    touch cert.crt; //补充证书
    touch cert.key; //补充证书私钥

    打开 /etc/nginx/conf.d/meet.conf 按照示例文件填充内容(需要替换域名和端口) 。

`Nginx配置文件示例 </meet/nginx_conf.html>`_


4. 重启nginx  
===============================================


.. code-block:: bash

    测试nginx配置是否正确
    nginx -t

    重启nginx 
    nginx -s reload

    如果nginx未启动 启动nginx
    nginx


5. 防火墙配置
===============================================

.. list-table:: 防火墙配置
   :widths: 25 25 25 25
   :header-rows: 1

   * - IP
     - 协议
     - 端口范围
     - 策略
   * - 0.0.0.0/0
     - UDP
     - 50000-60000
     - 允许
   * - 0.0.0.0/0
     - UDP
     - 3478-3500
     - 允许
   * - 0.0.0.0/0
     - UDP
     - 50000-60000
     - 允许
   * - 0.0.0.0/0
     - TCP
     - 7880-7890
     - 允许
   * - 0.0.0.0/0
     - TCP
     - 443
     - 允许
   * - 0.0.0.0/0
     - TCP
     - 80
     - 允许

6. 配置DNS
===============================================

.. code-block:: bash

    在域名服务商或者DNS解析服务商配置域名A记录指向服务器公网IP

