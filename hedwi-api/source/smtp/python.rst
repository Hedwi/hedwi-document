.. _help-smtp-python:

.. _smtp-python:


Python smtp 发邮件
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


Python可以使用yagmail发信，比较方便。

.. code-block:: python

        #!/usr/bin/env python
        #coding=utf-8

        import yagmail

        #建议: 同类型邮件的发信用户名最好固定，可以减少邮件被当做垃圾邮件的几率

        user = "support@notice.domain.com"  # 任意用户名@域名 推荐在hedwi配置子域名
        password = "*******************" #  SMTP 默认密码  参见域名详情页SMTP部分
        smtp_server = "smtp.hedwi.net" 
        port = "465"
        yag = yagmail.SMTP(user=user, password=password, host=smtp_server, port=port, smtp_ssl=True)

        to = "user@example.com" # 收件人
        subject = 'smtp 测试'  # 邮件主题
        contents = ['This is the body, and here is just text '] # 邮件内容
        yag.send(to, subject, contents, group_messages=False)

        #


当然也可以使用标准库发信

.. code-block:: python

        #!/usr/bin/env python
        #coding=utf-8

        import smtplib, ssl

        user = "support@notice.domain.com" #  发信人 任意用户名@域名  
        password = "*******************" #  SMTP 默认密码  参见域名详情页SMTP部分
        to = "user@example.com"  #  收信人
        smtp_server = "smtp.hedwi.net"  # 邮件服务器
        port = 465  # 端口


        message = """\
        Subject: Hedwi Mail test

        This message is sent from Hedwi."""

        context = ssl.create_default_context()
        server = smtplib.SMTP_SSL(smtp_server, port, context)
        server.login(user, password)
        server.sendmail(user, to, message)
        server.quit()

        # 
