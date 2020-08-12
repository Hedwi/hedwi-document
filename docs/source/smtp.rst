.. _help-smtp:

.. _smtp:

SMTP发信
------------

SMTP发信使用任意用户名和默认SMTP密码即可。

SMTP服务器：smtp.hedwi.net  端口: 25 / 25000/ 465

.. raw:: html

        推荐使用 <a target="_blank" href="https://www.mail-tester.com" ><img border="0" src="/_static/mail-tester_logo.svg" alt="mail-tester" style="width: 120px;" title="mailtester邮件测试"></a> 做测试

.. code-block:: shell

    ./swaks --auth \
	--server smtp.hedwi.net \
        --auth PLAIN  \
	--au username@example.com \
	--ap password \
	--to to@example.com \
        --tls \
	--h-Subject: "Hello" \
	--body 'Testing some Hedwi awesomness!'
