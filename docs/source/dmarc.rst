.. _help-dmarc:

.. _dmarc:

DMARC-基于域的消息认证，报告和一致性
-------------------------------------

引用自维基百科 `DMARC-基于域的消息认证，报告和一致性 <https://zh.wikipedia.org/wiki/%E5%9F%BA%E4%BA%8E%E5%9F%9F%E7%9A%84%E6%B6%88%E6%81%AF%E8%AE%A4%E8%AF%81%EF%BC%8C%E6%8A%A5%E5%91%8A%E5%92%8C%E4%B8%80%E8%87%B4%E6%80%A7>`_ 


基于域的消息认证，报告和一致性（DMARC，Domain-based Message Authentication, Reporting and Conformance）是一套以SPF及DKIM为基础的电子邮件认证机制，可以检测及防止伪冒身份、对付网络钓鱼或垃圾电邮。

网域管理员可以在域名系统公布相关政策，让外界得知旗下域名的电子邮件提供何种方式（SPF及/或DKIM）认证身份，以及如果寄件者身份未能百分之百确认时，收件者可以如何处理邮件（放进杂件箱或直接回绝）及回报。回报机制可以让网域管理员了解是否有第三者正在伪冒其网域身份寄出电邮。

外部链接
~~~~~~~~~~

- `DMARC项目官方网站 <https://dmarc.org/>`_
