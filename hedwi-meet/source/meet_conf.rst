

.. _help-manual-meet_conf:

.. _meet_conf:


meetserver 配置内容
------------------------
web_port: 8888
port: 7882
log_level: info
rtc:
  tcp_port: 7883
  port_range_start: 50000
  port_range_end: 60000
  # use_external_ip should be set to true for most cloud environments where
  # the host has a public IP address, but is not exposed to the process.
  # LiveKit will attempt to use STUN to discover the true IP, and advertise
  # that IP with its clients
  use_external_ip: false
redis:
  # redis is recommended for production deploys
  address: 127.0.0.1:6379
keys:
  # your_api_key: <api_secret>
  key: secret   #需要修改
# when enabled, LiveKit will expose prometheus metrics on :6789/metrics
#prometheus_port: 6789
turn:
  enabled: true
  # domain must match tls certificate
  domain: example.com #需要修改域名
  # defaults to 3478. If not using a load balancer, must be set to 443.
  tls_port: 3479
  cert_file: /data/www/ssls/example.com/cert.crt  #需要修改
  key_file: /data/www/ssls/example.com/cert.key   #需要修改