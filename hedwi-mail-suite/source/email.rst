.. _api-email:

.. _mail:


发送邮件
================

interface address:

* https://api.hedwi.com/mail/:domain

HTTP请求方式：

*  POST

鉴权：

* 参考 :ref:`api-auth` 章节

request params:

* from

  - 发件人

* to

  - 收件人

* subject

  - 邮件主题

* html

  - 邮件html内容

* text

  - 邮件纯文本内容


response code:

* common response

example::
    
    curl -X POST -s --user 'api:d5e76748-7b9a-11e9-bb91-f21898b25098' \
    'https://api.hedwi.com/mail/example.hedwi.com' \
    -F from='user@example.hedwi.com' \
    -F to='support@hedwi.com' \
    -F subject='Hello' \
    -F text='Testing email from Hedwi!'

response example：

.. code-block:: json

    {
        "code":0,
        "message":"干得漂亮！要开始使用请从这里注册帐号: http://hedwi.com/signup",
        "ts":1560340668
    }


模板发送邮件
======================

interface address:

* https://api.hedwi.com/mail/:domain

HTTP请求方式：

*  POST

鉴权：

* 参考 :ref:`api-auth` 章节

request params:

* from

  - 发件人

* to

  - 收件人

* subject

  - 邮件主题

* template

  - 模板名称

* hedwi-variables

  - 模板变量  

response code:

* common response

example::
    
    curl -X POST -s --user 'api:d5e76748-7b9a-11e9-bb91-f21898b25098' \
    'https://api.hedwi.com/mail/example.hedwi.com' \
    -F from='user@example.hedwi.com' \
    -F to='support@hedwi.com' \
    -F subject='Hello' \
    -F template='order' \
    -F hedwi-variables='{"name": "hedwi"}'

response example：

.. code-block:: json

    {
        "code":0,
        "message":"干得漂亮！要开始使用请从这里注册帐号: http://hedwi.com/signup",
        "ts":1560340668
    }

模板变量调用示例：

.. code-block:: html

    <html>
        <body>
            <h1>你好 {{name}}</h1>
        </body>
    </html>


