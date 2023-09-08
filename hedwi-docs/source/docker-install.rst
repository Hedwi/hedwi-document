.. _help-docker-install:

.. _docker-install:


使用 docker 部署（推荐）
------------------------

安装
=====================

.. code-block:: bash

    curl 'https://hedwi.com/install' |sh



配置
============

  如果使用自定义的域名 需要修改docs.dev中的 HEDWI_LICENSE 变量
  证书在 https://hedwi.com/license 申请

启动
============

.. code-block:: bash

    cd hedwi_pkgs/compose/docs && docker compose up -d

访问
============

    访问 http://127.0.0.1:40008 即可


注册 & 登录
============

    - 访问 http://127.0.0.1:40008/signup 注册管理员帐号
    - 访问 http://127.0.0.1:40008/domain 添加域名hedwi.com 或者申请的证书中的域名
    - 访问 http://localhost:40008/member 点击管理成员 添加成员 例如 support
    - 登录使用的账号格式是邮箱: 名称@域名, 例如support@hedwi.com 
    - 访问 http://localhost:40008/login  点击普通用户 使用刚才添加的普通用户账号登录