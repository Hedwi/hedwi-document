.. _help-firewall:

.. _firewall:


Firewall configuration
----------------------------------------------------------------------------


Configure the firewall
=============================

..  csv-table:: 
    :header: "Host", "Port", "Purpose"
    :widths: 35, 35, 30

    "\*","TCP: 80","HTTP Service"
    "\*","TCP: 443","HTTPS Service"
    "\*","TCP: 41883","mqtt service"
    "\*","TCP: 25","smtp service"
    "\*","TCP: 465","smtp over ssl service"
    "\*","UDP: 50000-60000","meeting UDP connection for WebRTC"
    "\*","TCP: 7881","meeting TCP connection for WebRTC"
    

- 80/443 is the port of the HTTP and HTTPS service
- 80 port can be used to obtain the free ssl certificate
- 41883 is the port of the mqtt service
- 25 is the port of the smtp service
- 465 is the port of the smtp over ssl service
- 50000-60000 is the port range of the UDP connection for WebRTC
- 7881 is the port of the TCP connection for WebRTC
- the meet port range is in the meet configuration file


