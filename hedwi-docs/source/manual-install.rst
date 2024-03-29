.. _help-manual-install:

.. _manual-install:


手动部署 
------------------------

环境要求
=====================
最低2G内存 2核CPU  linux 64位系统





minio / consul / nats / cockroach
======================================

- cockroach 数据库

    .. code-block:: bash





- minio

    .. code-block:: bash


        # 创建安装目录
        mkdir -p $HEDWI_ROOT/{consul,nats,hedis,hedwi-inbox,hedwi-task,hedwi-search}/logs ;

        cd $HEDWI_ROOT/minio ;
        wget https://dl.min.io/server/minio/release/linux-amd64/minio ;
        sudo chmod +x $HEDWI_ROOT/minio/minio ;

- consul

    .. code-block:: bash

        wget https://releases.hashicorp.com/consul/1.10.1/consul_1.10.1_linux_amd64.zip ;
        unzip consul_1.10.1_linux_amd64.zip ;
        mv consul $HEDWI_ROOT/consul/consul ;
        sudo chmod +x $HEDWI_ROOT/consul/consul ;
       

- nats

    .. code-block:: bash

        wget 'https://github.com/nats-io/nats-streaming-server/releases/download/v0.22.1/nats-streaming-server-v0.22.1-linux-amd64.tar.gz';
        tar xf nats-streaming-server-v0.22.1-linux-amd64.tar.gz ;
        mv nats-streaming-server-v0.22.1-linux-amd64/nats-streaming-server $HEDWI_ROOT/nats/nats;
        wget https://download.hedwi.com/mail/nats.conf
        mv nats.conf $HEDWI_ROOT/nats/


hedwi-mail
===============================================


- 安装 supervisor & unzip

    .. code-block:: bash

        ubuntu/debian:  
            sudo apt-get install python3-pip unzip ;        
        centos: 
            sudo yum install python34-pip.noarch unzip ;

        pip3 install supervisor ;


- hedwi-mail

    .. code-block:: bash

        cd $HEDWI_ROOT ;
        wget https://download.hedwi.com/mail/linux/amd64.zip ;
        unzip amd64.zip ;
        cd $HEDWI_ROOT/amd64 ;
        ls *|xargs -I {} cp -f {} $HEDWI_ROOT/{}/{} ;
        #增加可执行
        ls *|xargs -I {} chmod +x $HEDWI_ROOT/{}/{} ;

        cd $HEDWI_ROOT ; 
        wget https://download.hedwi.com/mail/hedwi-mail_supervisord.conf ;

        # todo 在supervisord配置文件中配置用户名 密码 
        # MINIO_ROOT_USER='user'        #kv server user  minio user
        # HEDWI_OSS_ACCESS_KEY='user'         #kv server user minio user
        # MINIO_ROOT_PASSWORD='password'    #kv server password minio password
        # HEDWI_OSS_ACCESS_SECRET='password'      #kv server password minio password
 
        
        访问https://hedwi.com/license 生成license （目前需要注册帐号生成license）
        10人免费 生成后配置到supervisord.conf文件中 HEDWI_LICENSE变量。

        supervisord -c hedwi-mail_supervisord.conf

- 防火墙配置
    对外开放25、465、993、443端口即可 本邮箱服务不需要nginx 如果使用nginx则需要自行修改配置文件中的端口
- 创建管理员帐号
    访问https://mail.example.com (此处可以是根域名或者任意子域名，只要配置好DNS记录就可以)。注册管理员帐号 第一个帐号就是管理员帐号。使用管理员帐号添加域名，配置DNS记录。