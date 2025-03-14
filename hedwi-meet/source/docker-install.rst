.. _help-docker-install:

.. _docker-install:


docker install
----------------------------------------------------------------------------

0. Requirements 
=====================
Minimum 1G memory 1 core CPU linux 64-bit system

1. Configure DNS
=====================

.. code-block:: bash

Configure the A record of the domain name to the public IP address of the server in the domain name service provider or DNS resolution service provider


For example

..  csv-table:: 
    :header: "Host name", "Record type", "Record value"
    :widths: 35, 35, 30

    "meet","A","x.x.x.x"


- The host name can be any secondary domain name prefix, here we use meet
- example.com is your domain name
- x.x.x.x is the public IP address of the server
- The web access address is https://meet.example.com/meet  


2. Install docker, git 
======================================

.. code-block:: bash

  # If already installed, skip

  mkdir pkgs && cd pkgs
  wget 'https://download.docker.com/linux/static/stable/x86_64/docker-26.1.4.tgz'
  tar xf docker-26.1.4.tgz;
  chmod +x docker/*
  sudo cp docker/* /usr/local/bin/

  curl -SL https://github.com/docker/compose/releases/download/v2.27.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose

  yum install git # centos
  apt install git  # ubuntu/debian

  # start docker daemon
  sudo dockerd &


3. Install
=====================


.. code-block:: bash

  # github 国内使用github更快
  git clone https://github.com/Hedwi/hedwi-docker.git

  # gitee 国内使用gitee更快
  git clone https://gitee.com/hedwi/hedwi-docker.git

  cd hedwi-docker/meet
  cp env.example .env
  # Modify the .env configuration

  make

- `配置文件示例 </document/meet/zh-hans/meet_env.html>`_


-  `Configuration file example </document/meet/en-us/meet_env.html>`_


4. Automatically generate free SSL certificates
==========================================================

.. code-block:: bash

  make cert


5. Restart nginx  
===============================================

.. code-block:: bash

  make down && make

