.. _python-send-email-code-example:

.. _python-code-example:

 使用python发送邮件
-------------------------------------

使用BasicAuth，需要在HTTP请求headers中增加Authorization
    
Authorization生成算法：用冒号拼接字符串api和你的key，然后转成base64，再拼接"Basic "前缀即可。

- golang示例:  
  
.. code-block:: go

    key := "d5e76748-7b9a-11e9-bb91-f21898b25098"
    source := "api:" + key
    auth := "Basic " + b64.StdEncoding.EncodeToString([]byte(source))
    headers := map[string]string{"Authorization": auth}


.. code-block:: go

    import "fmt"
    import b64 "encoding/base64"
    import "github.com/levigross/grequests"

    func main() {
        auth := b64.StdEncoding.EncodeToString([]byte(
            "api:d5e76748-7b9a-11e9-bb91-f21898b25098"))
        ro := &grequests.RequestOptions{
            Headers: map[string]string{"Authorization": "Basic " + auth},
            Data: map[string]string{
                "from": "user@example.hedwi.com",
                "to": "support@hedwi.com",
                 "subject": "Hello",
                 "text": "Testing email from Hedwi!",
            },
        }
        r, _ := grequests.Post("https://api.hedwi.com/mail/example.hedwi.com", ro)
        fmt.Println(r.String())
    }
              

- python示例:  

.. code-block:: python

    from base64 import b64encode
    key = "d5e76748-7b9a-11e9-bb91-f21898b25098"
    source = "api:" + key
    auth = "Basic " + b64encode(source)
    headers = {"Authorization": auth}
 

.. code-block:: python

    #!/usr/bin/env python
    # -*- coding: utf-8 -*-

    import string
    import requests
    from base64 import b64encode

    def auth(key):
        source = "api:" + key
        auth = "Basic " + b64encode(source)
        headers = {"Authorization": auth}
        return headers


    def main():

        key = "d5e76748-7b9a-11e9-bb91-f21898b25098"

        data = {
            "from": "user@example.hedwi.com>",
            "to": "support@hedwi.com",
            "subject": "Hello",
            "text": "Testing email from Hedwi!"
        }
        url = "https://api.hedwi.com/mail/example.hedwi.com"

        use_headers = True
        if use_headers:
            headers = auth(key)
            #使用headers
            r = requests.post(url, headers = headers, data = data)
            print r.content
        else:
            #使用requests basic auth
            r = requests.post(url, auth = ("api", key), data = data)
            print r.content

    if __name__ == "__main__":
        main()


- PHP示例：

.. code-block:: php

    <?php
        function sendEmail() {
            $data = array(
                "from" => "user@example.hedwi.com>",
                "to" => "support@hedwi.com",
                "subject" => "Hello",
                "text" => "Testing email from Hedwi!"
            );
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
    ?>



- C#示例：

.. code-block:: C#

    public static RestResponse SendSimpleMessage() {
        RestClient client = new RestClient("https://api.hedwi.com");
        client.Authenticator = new HttpBasicAuthenticator(
            "api","d5e76748-7b9a-11e9-bb91-f21898b25098");
        RestRequest request = new RestRequest();
        request.AddParameter("domain", "example.hedwi.com", ParameterType.UrlSegment);
        request.Resource = "mail/{domain}";
        request.AddParameter("from", "user@example.hedwi.com");
        request.AddParameter("to", "support@hedwi.com");
        request.AddParameter("subject", "Hello");
        request.AddParameter("text", "Testing email from Hedwi!");
        request.Method = Method.POST;
        return client.Execute(request);
    }

- Ruby示例

.. code-block:: Ruby

    def send_simple_message
      RestClient.post "https://api:d5e76748-7b9a-11e9-bb91-f21898b25098"
      "@api.hedwi.com/mail/example.hedwi.com",
      :from => "user@example.hedwi.com",
      :to => "support@hedwi.com",
      :subject => "Hello",
      :text => "Testing email from Hedwi!"
    end
