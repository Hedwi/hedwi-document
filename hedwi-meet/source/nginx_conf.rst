.. _help-manual-nginx:

.. _nginx_conf:


nginx configuration
----------------------------------------------------------------------------


.. code-block:: bash

    server {
        listen              443 ssl http2;
        listen              [::]:443 ssl http2;
        server_name         example.com;
        set                 $base /data/www/hedwi-meetserver;
        root                $base;

        # SSL
        ssl_certificate     /data/www/ssls/example.com/cert.crt;
        ssl_certificate_key /data/www/ssls/example.com/cert.key;

        # logging
        access_log          /var/log/nginx/meet.access.log combined buffer=512k flush=1m;
        error_log           /var/log/nginx/meet.error.log warn;

        # index.php
        index               index.html;

        location /api {

            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header Host $host;
                proxy_pass            http://127.0.0.1:8888;
        }


        location /meet {

            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header Host $host;
                proxy_pass            http://127.0.0.1:8888;
        }


        location / {

            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header Host $host;
                proxy_pass            http://127.0.0.1:7882;
        }


    }

    # HTTP redirect
    server {
        listen      80;
        listen      [::]:80;
        server_name example.com;
        return      301 https://example.com$request_uri;
    }