.. _api-auth:

.. _auth:

鉴权
------------

使用BasicAuth，需要在请求headers中增加Authorization
    
Authorization生成算法：用冒号拼接字符串api和你的key，然后转成base64，再拼接"Basic "前缀即可。

- golang示例:  
  
.. code-block:: go

    key := "d5e76748-7b9a-11e9-bb91-f21898b25098"
    source := "api:" + key
    auth := "Basic " + b64.StdEncoding.EncodeToString([]byte(source))
    headers := map[string]string{"Authorization": auth}

- python示例:  

  完整示例 :download:`auth.py <auth.py>`

.. code-block:: python

    from base64 import b64encode
    key = "d5e76748-7b9a-11e9-bb91-f21898b25098"
    source = "api:" + key
    auth = "Basic " + b64encode(source)
    headers = {"Authorization": auth}
 
- php示例：

.. code-block:: php

    function sendEmail() {
        $data = array("from" => "user@example.hedwi.com>", "to" => "support@hedwi.com",
                      "subject" => "Hello", "text" => "Testing email from Hedwi!");
        $ch = curl_init("https://api.hedwi.com/mail/example.hedwi.com");
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLINFO_HEADER_OUT, true);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
        curl_setopt($ch, CURLOPT_USERPWD, "api:d5e76748-7b9a-11e9-bb91-f21898b25098");
        $result = curl_exec($ch);
        curl_close($ch);
        return $result
    }
