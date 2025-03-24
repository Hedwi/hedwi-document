.. _help-docker-install:

.. _docker-install:


docker deploy
----------------------------------------------------------------------------


About
======================

To be a nice google workspace alternative


0. Requirements
=====================
Minimum 1G memory 1 CPU 64-bit Linux system

1. configure DNS
===============================================

.. code-block:: bash

   Configure A record in the domain name service provider or DNS resolution service provider

For example:

..  csv-table:: 
    :header: "hostname", "record type", "record value"
    :widths: 35, 35, 30

    "work","A","x.x.x.x"


- The hostname can be any secondary domain name prefix, here use work
- example.com is your domain name
- x.x.x.x is the public IP address of the server
- The web access address is https://work.example.com

2. Install docker, docker-compose, git
==========================================

.. code-block:: bash

  # If already installed, skip

  mkdir pkgs && cd pkgs
  wget 'https://download.docker.com/linux/static/stable/x86_64/docker-26.1.4.tgz'
  tar xf docker-26.1.4.tgz;
  chmod +x docker/*
  sudo cp docker/* /usr/local/bin/

  curl -SL https://github.com/docker/compose/releases/download/v2.27.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose

  yum install git # centos
  apt install git  # ubuntu/debian


3. Install
===============================================

.. code-block:: bash


  # github 国外用户使用github更快
  git clone https://github.com/Hedwi/hedwi-docker.git

  # gitee 国内用户使用gitee更快
  git clone https://gitee.com/hedwi/hedwi-docker.git

  cd hedwi-docker/hedwi
  cp env.example .env

| Modify the .env configuration, modify all variables marked with 'Need modify'

  make  

| After running, http 80 port can be accessed, should be able to access http://configured domain name
| The service will automatically create an administrator account (for configuring the system, adding users, adding teams, etc.) and a normal user account admin@your domain name (for using the system service)

`Configuration file example <./env-example.html>`_


4. Automatic generation of free SSL certificates
==============================================================

| The service has built-in Let’s Encrypt HTTP-01 verification service interface, running this command will request the /acme/create interface to create a certificate
| After obtaining the certificate successfully, the certificate will be written to inbox/certs/cert.key inbox/certs/cert.crt
| This command will use the generated free certificate to replace nginx/certs/cert.key nginx/certs/cert.crt

.. code-block:: bash

  make cert

5. Restart the service
===============================================

| Stop all running docker containers

.. code-block:: bash
  make down 

| Start docker containers

.. code-block:: bash
  make 

6. Use https to access and login
===============================================

Access https://work.example.com/login?admin=true (work.example.com is the domain name used), use the administrator account (the email and password configured) to log in
Access https://work.example.com/hello page, continue to add users, teams, etc.
