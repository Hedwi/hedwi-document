.. _help-manual-install:

.. _manual-install:


Manual deployment
----------------------------------------------------------------------------

0. Requirements
=====================
Minimum 1G memory 1 core CPU linux 64-bit system


1. Install dependencies
================================

.. code-block:: bash

    centos / redhat

        yum install -y nginx supervisor redis

    debian / ubuntu 

        apt install -y nginx supervisor redis


2. Install meet server
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
    # Modify the domain name in the configuration file.

    supervisord -c /data/www/hedwi-meetserver.conf;


`Configuration file example </meet/meet_conf.html>`_



3. nginx and ssl certificate configuration
===============================================


.. code-block:: bash

    mkdir -p $HEDWI_ROOT/ssls/;
    cd $HEDWI_ROOT/ssls;
    mkdir example.com; //Domain certificate folder, here is the corresponding domain name
    cd !$; 
    touch cert.crt; //Add certificate
    touch cert.key; //Add certificate private key

    Open /etc/nginx/conf.d/meet.conf and fill in the content according to the example file (replace the domain name and port).

`Nginx configuration file example </meet/nginx_conf.html>`_


4. Restart nginx  
===============================================


.. code-block:: bash

    # Test nginx configuration
    nginx -t

    # Restart nginx 
    nginx -s reload

    #If nginx is not started, start nginx
    nginx


5. Firewall configuration
===============================================

.. list-table:: Firewall configuration
   :widths: 25 25 25 25
   :header-rows: 1

   * - IP
     - Protocol
     - Port range
     - Policy
   * - 0.0.0.0/0
     - UDP
     - 50000-60000
     - Allow
   * - 0.0.0.0/0
     - UDP
     - 3478-3500
     - Allow
   * - 0.0.0.0/0
     - UDP
     - 50000-60000
     - Allow
   * - 0.0.0.0/0
     - TCP
     - 7880-7890
     - Allow
   * - 0.0.0.0/0
     - TCP
     - 443
     - Allow
   * - 0.0.0.0/0
     - TCP
     - 80
     - Allow

6. Configure DNS
===============================================

.. code-block:: bash

    # Configure the A record of the domain name in the domain name service or DNS resolution service to point to the public network IP of the server

