.. _help-manual-meet_env:

.. _meet_env:


meetserver configuration
----------------------------------------------------------------------------

Only pay attention to the part marked with 『* Need modify』


.. code-block:: bash

    DOCKER_REGISTRY="registry.cn-shanghai.aliyuncs.com/"    #   If accessing docker hub is slow, you can change to use Aliyun Container Service "registry.cn-shanghai.aliyuncs.com/"
    DOMAIN="meet.example.com"                               # * Need modify domain name Suggest using subdomain
    BINDIP="127.0.0.1"                                      #   Listening address
    REDIS_ADDR="127.0.0.1:6379"                             #   redis address
    REDIS_USER=""                                           #   redis user name
    REDIS_PASSWORD=""                                       #   redis password
    MEET_SERVER="/"                                         #   meet server address
    MEET_KEY="changekeyandsecret"                           # * Need modify video conference key  
    MEET_SECRET="c7056e2x9689f7f3ecx8c6e4055c4bc565dc9c5a"  # * Need modify video conference secret 32-bit 
    MEET_ADDR="127.0.0.1:8888"                              #   meeting web listening address
    MEET_AUTH_TYPE="ldap"                                   #   Authentication type  ldap or none  none means no authentication
    MEET_LDAP_ENABLED="true"                                #   Whether to enable ldap authentication
    MEET_LDAP_ADDR="ldap://127.0.0.1:42389"                 #   Need modify
    MEET_LDAP_BASE_DN="dc=hedwi,dc=com"                     #   Need modify