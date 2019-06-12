.. _api-auth:

.. _auth:

鉴权
------------

* 使用BasicAuth，需要在请求headers中增加Authorization
    
* Authorization生成算法：用冒号拼接字符串api和你的key，然后转成base64，再拼接"Basic "前缀即可。
* golang示例:  
  
.. code-block:: go

    key := "d5e76748-7b9a-11e9-bb91-f21898b25098"
    source := "api:" + key
    auth := "Basic " + b64.StdEncoding.EncodeToString([]byte(source))
    headers := map[string]string{"Authorization": auth}

