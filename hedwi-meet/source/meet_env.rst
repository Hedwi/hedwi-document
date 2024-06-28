.. _help-manual-meet_env:

.. _meet_env:


meetserver 配置
------------------------


.. code-block:: bash

    DOCKER_REGISTRY="registry.cn-shanghai.aliyuncs.com/" # 如果访问docker hub比较慢，可以改成使用阿里云容器服务 "registry.cn-shanghai.aliyuncs.com/"
    DOMAIN="meet.example.com" # 域名 建议使用子域名  需要修改
    BINDIP="127.0.0.1" # 监听地址
    REDIS_ADDR="127.0.0.1:6379"  # redis 地址
    REDIS_USER="" # redis 用户名
    REDIS_PASSWORD="" # redis 密码
    HEDWI_MEET_SERVER="/"  # meet server 地址
    HEDWI_MEET_KEY="changekeyandsecret"  # 视频会议key  需要修改
    HEDWI_MEET_SECRET="c7056e2x9689f7f3ecx8c6e4055c4bc565dc9c5a" # 视频会议secret 32位 需要修改
    HEDWI_MEET_ADDR="127.0.0.1:8888" # 视频会议web 监听地址