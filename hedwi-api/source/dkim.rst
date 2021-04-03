.. _help-dkim:

.. _dkim:


DKIM - 域名密钥识别邮件
------------------------

引用自维基百科 `域名密钥识别邮件 <https://zh.wikipedia.org/wiki/%E5%9F%9F%E5%90%8D%E5%AF%86%E9%92%A5%E8%AF%86%E5%88%AB%E9%82%AE%E4%BB%B6>`_ 


域名密钥识别邮件（DomainKeys Identified Mail，DKIM）是一套电子邮件认证机制，使用公开密钥加密的基础提供了数字签名与身份验证的功能，以检测寄件者、主旨、内文、附件等部分有否被伪冒或窜改。

一般来说，发送方会在电子邮件的标头插入DKIM-Signature及电子签名信息。而接收方则透过DNS查询得到公开密钥后进行验证。


历史
===========

DKIM是2004年在雅虎的DomainKeys和思科系统的Identified Internet Mail两套相近的基础上合并并改进而来的协议，大多数的运作方式与DomainKeys相同。在2007年2月时，DKIM被列入互联网工程工作小组（IETF）的标准提案（Proposed Standard），并于同年5月成为正式标准（Standards Track）并得到主要的电子邮件服务供应商如雅虎、Gmail、美国在线、FastMail率先支持。

优点
============


DKIM的主要优点是可以让寄件者有效地表明身份，让收件者可凭借公钥确认寄件者并非伪冒、内文未经窜改，提高电子邮件的可信度。邮件过滤器可使用白名单及黑名单机制更可靠地检测网络钓鱼或垃圾电邮。


hedwi dkim配置
===================

请参考域名详情页DNS配置部分


外部链接
===================

+ （英文） `DomainKeys Identified Mail (DKIM) <http://www.dkim.org/>`_  DKIM的信息
+ （英文） `RFC 4408 <https://tools.ietf.org/html/rfc4408>`_
