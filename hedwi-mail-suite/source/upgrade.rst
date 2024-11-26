.. _help-upgrade:

.. _upgrade:


Upgrade
----------------------------------------------------------------------------

0. requirements
=====================
minimum 2G memory 2 CPU 64-bit linux system

1. configure DNS
=====================

.. code-block:: bash

   configure domain A record to point to the server's public IP

For example:

..  csv-table:: 
    :header: "hostname", "record type", "record value"
    :widths: 35, 35, 30

    "work","A","x.x.x.x"


- hostname can be any subdomain prefix, here use work
- example.com is your domain name
- x.x.x.x is the server's IP address
- web access address is https://work.example.com
