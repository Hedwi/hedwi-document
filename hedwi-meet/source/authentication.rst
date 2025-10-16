.. _help-authentication:

.. _authentication:


Authentication
----------------------------------------------------------------------------

1. Authentication type
=============================

- none: No authentication
- ldap: LDAP authentication

2. How to configure authentication
============================================
    set MEET_AUTH_TYPE in .env file

3. How to configure LDAP authentication
============================================

    set the following variables in .env file

    MEET_LDAP_ENABLED="true"                  # whether to enable ldap authentication
    MEET_LDAP_ADDR="ldap://127.0.0.1:42389"   #ldap server address
    MEET_LDAP_BASE_DN="dc=hedwi,dc=com"       # ldap base dn

4. How to use authentication in mobile app
=============================================

    - ldap enabled, the mobile app will use ldap authentication

        password =  ldap password + " " + meeting password


    - ldap disabled, the mobile app will use meeting password and if meeting password is not set, leave it blank

        password =  meeting password or blank