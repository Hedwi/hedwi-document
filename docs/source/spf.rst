.. _help-spf:

.. _spf:


SPF - 发件人策略框架
------------------------

引用自维基百科 `发件人策略框架 <https://zh.wikipedia.org/wiki/%E5%8F%91%E4%BB%B6%E4%BA%BA%E7%AD%96%E7%95%A5%E6%A1%86%E6%9E%B6>`_ 


发件人策略框架（英语：Sender Policy Framework；简称SPF； RFC 4408）是一套电子邮件认证机制，可以确认电子邮件确实是由网域授权的邮件服务器寄出，防止有人伪冒身份网络钓鱼或寄出垃圾电邮。SPF允许管理员设定一个DNS TXT记录或SPF记录设定发送邮件服务器的IP范围，如有任何邮件并非从上述指明授权的IP地址寄出，则很可能该邮件并非确实由真正的寄件者寄出（邮件上声称的“寄件者”为假冒）。


用途
===============


SPF 记录的用途是阻止垃圾邮件发件人发送假冒您的域中的“发件人”地址的电子邮件。收件人可以参考 SPF 记录来确定号称来自您的域的邮件是否来自授权邮件服务器。对于大多主流的邮件服务商，鉴别发送者的SPF记录有助于抵御垃圾邮件给接收者带来的骚扰。

原理
===============

SMTP协议本身没有机制鉴别寄件人的真正身份，电子邮件的“寄件者”一栏可以填上任何名字，于是伪冒他人身份来网络钓鱼或寄出垃圾邮件便相当容易，而真正来源却不易追查。

SPF 记录允许域名管理员公布所授权的邮件服务器IP地址，接收服务器会在收到邮件时验证发件人在 SMTP 会话中执行 MAIL FROM 命令时的邮件地址是否与域名 SPF 记录中所指定的源 IP 匹配，以判断是否为发件人域名伪造。

邮件接收方的收件服务器在接受到邮件后，首先检查域名的SPF记录，来确定发件人的IP地址是否被包含在SPF记录里面，如果在，就认为是一封来自于被授权服务器的邮件，否则会认为是一封伪造的邮件并根据相关政策退回或放进收件人的杂件箱。


简单来说，SPF 记录其实就是一条有特殊语法的 TXT 记录，它由“匹配机制”和“修饰符”2 部分组成。“修饰符”通常只作为可选项，一般情况下 Exchange 管理员只会用到并处理包含“匹配机制”的 SPF 记录。
SPF 记录的匹配机制主要用于定义和指定可由该域名发送邮件的主机，其定义方式包括：


- all 匹配任何主机，它写在 SPF 记录最后以匹配在其前面所列出的主机。
- ip4 匹配 IPv4 地址或网络范围。
- ip6 匹配 IPv6 地址或网络范围。
- a 匹配主机名或域名。
- mx 匹配域名的 MX 记录，当出站与入站邮件为同一服务器时通常采用此种机制。
- ptr 通过 DNS 反向记录来匹配发件人 IP 和域名，由于会增加 DNS 负载，一般不采用此种机制。
- exists 只检查域是否在 DNS 中存在。
- include 将发件人 IP 和 SPF 记录指向另一个域，这种匹配机制通常用于云服务，如 Exchange Online Protection。

SPF 记录的匹配机制会结合一些限定词来使用，以告诉服务器找到一条匹配记录时该怎么办。常见的限定词有：

- \+ 放行，如果没有明确指定限定词，则为默认值。
- – 硬拒绝，直接拒绝来自未经授权主机的邮件。
- ~ 软拒绝，邮件可被接受，也可被标记为垃圾邮件。
- ? 中性，不考虑邮件是否被接受。

SPF 记录会结合“匹配机制”和“限定词”使用，例如：在 SPF 记录的末尾处写有 -all 则表明在不匹配前面所列主机时，接收服务器需要将邮件全部拒绝。

hedwi spf配置
================

TXT记录值为 v=spf1 include:spf.hedwi.net ~all

外部链接
================

- `Sender Policy Framework项目网站 <https://web.archive.org/web/20080513111421/http://www.openspf.org/>`_
- `RFC 4408 <https://tools.ietf.org/html/rfc4408>`_


来源
==========

- `http://www.ietf.org/rfc/rfc4408.txt <http://www.ietf.org/rfc/rfc4408.txt>`_
- `https://support.google.com/a/answer/33786?hl=zh-Hans <https://support.google.com/a/answer/33786?hl=zh-Hans>`_
        - `http://service.exmail.qq.com/cgi-bin/help?subtype=1&&id=20012&&no=1000580t <http://service.exmail.qq.com/cgi-bin/help?subtype=1&&id=20012&&no=1000580>`_
        - `http://www.sysgeek.cn/sender-policy-framework-spf-exchange/ <http://www.sysgeek.cn/sender-policy-framework-spf-exchange/>`_

