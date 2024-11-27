.. _help-manual-cluster-install:

.. _manual-cluster-install:


Manual cluster deployment
----------------------------------------------------------------------------

Docker installation
=====================

.. code-block:: bash

    curl -O 'https://hedwi.com/cluster' | sh

Start
============

.. code-block:: bash

    docker compose up -d

- Firewall configuration
     Open ports 25, 465, 993, 443 for external access. The email service does not require nginx. If using nginx, you need to modify the port in the configuration file.
- Create an administrator account
     Access https://127.0.0.1:40008 (This can be the root domain name or any subdomain, as long as the DNS record is configured). Register an administrator account. The first account is the administrator account. Use the administrator account to add a domain name and configure the DNS record.