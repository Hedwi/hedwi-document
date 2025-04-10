.. _help-firewall:

.. _firewall:


Firewall configuration
----------------------------------------------------------------------------


Configure the firewall
=============================

..  csv-table:: 
    :header: "Host", "Port", "Purpose"
    :widths: 35, 35, 30

    "\*","TCP: 443","HTTP Service"
    "\*","UDP: 50000-60000","UDP connection for WebRTC"
    "\*","TCP: 7881","TCP connection for WebRTC"


- 443 is the port of the HTTP service
- 50000-60000 is the port range of the UDP connection for WebRTC
- 7881 is the port of the TCP connection for WebRTC
- the port range is in the configuration file


