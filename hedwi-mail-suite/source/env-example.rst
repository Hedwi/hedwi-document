
.. _help-env-example:

.. _env-example:

env 配置内容
------------------------

只需要关注 『需要修改』的部分

.. code-block:: bash

  DOCKER_REGISTRY="registry.cn-shanghai.aliyuncs.com/" #  如果不使用docker镜像，改成空
  DOMAIN="work.example.com"  # 域名 建议使用子域名  需要修改
  ADMIN_EMAIL="xxx@example.com" # 管理员邮箱 需要修改
  ADMIN_PASSWORD="passwordneedchange" # 管理员密码 需要修改
  BINDIP="127.0.0.1"  # 监听地址
  MINIO_USER="changeusername"  # minio 用户名 需要修改
  MINIO_PASSWORD="changepassword" # minio 密码 需要修改
  MINIO_ADDR="127.0.0.1:9000" #  minio 监听地址
  MINIO_USESSL="false" # minio 是否使用https
  MINIO_WEB_ADDR="127.0.0.1:9001" # minio web界面监听地址
  HEDWI_LICENSE="" #  hedwi证书  在网站 https://hedwi.com/signup 注册管理员账号 然后在 https://hedwi.com/license 生成  需要修改
  CRDB_ADDR="127.0.0.1:26257" # cockroachdb 监听地址
  CRDB_WEB_ADDR="127.0.0.1:42080" # cockroachdb web 监听地址
  DB_NAME="hedwi_selfhost" #  数据库名称
  OFFICE_API_ADDR="127.0.0.1:42201" #office api服务监听地址
  CONSUL_ADDR="127.0.0.1:8500" # consul 监听地址
  REDIS_ADDR="127.0.0.1:6379" # redis 地址
  REDIS_USER="" #  redis默认监听127.0.0.1 没有密码
  REDIS_PASSWORD="" #  redis默认监听127.0.0.1 没有密码
  IMAP_ADDR="127.0.0.1:42994" # imap服务器监听地址
  IMAP_TLS_ADDR="127.0.0.1:42993" # imap tls服务器监听地址
  HUB_ADDR="127.0.0.1:41002" # hub rpc 服务监听地址
  SEARCH_ADDR="127.0.0.1:41701" # search 搜索服务监听地址
  FILTER_ADDR="127.0.0.1:41003" # 邮件过滤服务监听地址
  NATS_URL="nats://127.0.0.1:42422" # nats 消息队列监听地址
  ALLOW_ORIGIN="http://127.0.0.1:4200" # 跨域设置
  SMTP_ADDR="127.0.0.1:42025" # smtp 服务器监听地址
  SMTP_TLS_ADDR="127.0.0.1:42465" # smtp tls服务器监听地址
  HEDWI_MEET_SERVER="/" # 会议服务器地址
  HEDWI_MEET_KEY="changekeyandsecret" # 视频会议创建使用的key 需要修改
  HEDWI_MEET_SECRET="c7056e2x9689f7f3ecx8c6e4055c4bc565dc9c5a" #  视频会议创建使用的secret 需要修改
  HEDWI_MEET_ADDR="127.0.0.1:8888" # 视频会议web监听地址
