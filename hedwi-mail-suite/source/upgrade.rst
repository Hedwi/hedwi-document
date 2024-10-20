.. _help-upgrade:

.. _upgrade:


Upgrade
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
