.. _help-upgrade:

.. _upgrade:


Upgrade
----------------------------------------------------------------------------



1. Read update release notes
=================================


2. Backup
===============================
    Backup your data, such as database, files, etc.


3. Update database tables
===============================

    if need update database tables, update them before upgrade


4. Upgrade
===============================
    | # update docker image 
    | make update
    | # stop docker container 
    | make down   
    | # start docker container 
    | make      
