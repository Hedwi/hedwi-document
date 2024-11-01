.. _help-manual-meet_env:

.. _meet_env:


meetserver 配置
------------------------

只需要关注 『* 需要修改』的部分


.. code-block:: bash

    DOCKER_REGISTRY="registry.cn-shanghai.aliyuncs.com/"    #   如果访问docker hub比较慢，可以改成使用阿里云容器服务 "registry.cn-shanghai.aliyuncs.com/"
    DOMAIN="meet.example.com"                               # * 需要修改 域名 建议使用子域名
    BINDIP="127.0.0.1"                                      #   监听地址
    REDIS_ADDR="127.0.0.1:6379"                             #   redis 地址
    REDIS_USER=""                                           #   redis 用户名
    REDIS_PASSWORD=""                                       #   redis 密码
    MEET_SERVER="/"                                         #   meet server 地址
    MEET_KEY="changekeyandsecret"                           # * 需要修改 视频会议key  
    MEET_SECRET="c7056e2x9689f7f3ecx8c6e4055c4bc565dc9c5a"  # * 需要修改 视频会议secret 32位 
    MEET_ADDR="127.0.0.1:8888"                              #   视频会议web 监听地址
