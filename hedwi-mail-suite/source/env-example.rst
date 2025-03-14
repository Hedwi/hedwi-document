
.. _help-env-example:

.. _env-example:

env 配置内容
----------------------------------------------------------------------------

只需要关注 『* 需要修改』的部分

.. code-block:: bash

  DOCKER_REGISTRY="registry.cn-shanghai.aliyuncs.com/"    #   如果访问docker hub比较慢，可以改成使用阿里云容器服务 "registry.cn-shanghai.aliyuncs.com/"
  DOMAIN="work.example.com"                               # * 需要修改 域名 建议使用子域名  
  ADMIN_EMAIL="xxx@example.com"                           # * 需要修改 管理员邮箱 
  ADMIN_PASSWORD="passwordneedchange"                     # * 需要修改 管理员密码
  BINDIP="127.0.0.1"                                      #   监听地址 不经过nginx代理的时候应该使用0.0.0.0或者指定地址
  MINIO_USER="changeusername"                             # * 需要修改 minio 用户名 
  MINIO_PASSWORD="changepassword"                         # * 需要修改 minio 密码
  MINIO_ADDR="127.0.0.1:9000"                             #   minio 监听地址
  MINIO_USESSL="false"                                    #   minio 是否使用https
  MINIO_WEB_ADDR="127.0.0.1:9001"                         #   minio web界面监听地址
  HEDWI_LICENSE=""                                        # * 需要修改  hedwi证书  在网站 https://hedwi.com/signup 注册管理员账号 然后在 https://hedwi.com/license 生成
  PUBLIC_DRIVE_PREFIX=""                                  #   minio访问前缀 与nginx配置对应 默认不需要填写
  CRDB_ADDR="127.0.0.1:26257"                             #   cockroachdb 监听地址
  CRDB_WEB_ADDR="127.0.0.1:42080"                         #   cockroachdb web 监听地址
  DB_NAME="hedwi_selfhost"                                #   数据库名称
  OFFICE_API_ADDR="127.0.0.1:42201"                       #   office api服务监听地址
  CONSUL_IP="127.0.0.1"                            #   consul 监听IP
  CONSUL_ADDR="127.0.0.1:8500"                            #   consul 监听地址
  REDIS_ADDR="127.0.0.1:6379"                             #   redis 地址
  REDIS_USER=""                                           #   redis默认监听127.0.0.1 没有密码
  REDIS_PASSWORD=""                                       #   redis默认监听127.0.0.1 没有密码
  IMAP_ADDR="127.0.0.1:42994"                             #   imap服务器监听地址
  IMAP_TLS_ADDR="127.0.0.1:42993"                         #   imap tls服务器监听地址
  HUB_ADDR="127.0.0.1:41002"                              #   hub rpc 服务监听地址
  SEARCH_ADDR="127.0.0.1:41701"                           #   search 搜索服务监听地址
  FILTER_ADDR="127.0.0.1:41003"                           #   邮件过滤服务监听地址
  NATS_URL="nats://127.0.0.1:42422"                       #   nats 消息队列监听地址
  ALLOW_ORIGIN="http://127.0.0.1:4200"                    #   跨域设置
  SMTP_ADDR="127.0.0.1:42025"                             #   smtp 服务器监听地址
  SMTP_TLS_ADDR="127.0.0.1:42465"                         #   smtp tls服务器监听地址
  MEET_SERVER="/"                                         #   会议服务器地址
  MEET_KEY="changekeyandsecret"                           # * 需要修改 视频会议创建使用的key 
  MEET_SECRET="c7056e2x9689f7f3ecx8c6e4055c4bc565dc9c5a"  # * 需要修改 视频会议创建使用的secret 
  MEET_ADDR="127.0.0.1:8888"                              #   视频会议web监听地址
  MQTT_ID="1"                                             #   mqtt client id  
  MQTT_IP="127.0.0.1"                                     #   mqtt 监听地址
  MQTT_PORT="1883"                                        #   mqtt 监听端口
  MQTT_PROTOCOL="wss"                                     #   mqtt 协议
  MQTT_PUBLIC_PORT="41883"                                #   mqtt 公网端口
  AI_TOKEN=""                                             #  需要修改 AI token
  AI_MODEL="deepseek-chat"                                #  需要修改 AI模型
  AI_TYPE="deepseek"                                      #  需要修改 AI类型
