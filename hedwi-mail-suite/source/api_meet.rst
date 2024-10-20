.. _api_meet:

.. _api-meet-example:


.. _meet.gettoken:
获取加入会议的授权
-----------------------
接口地址：
    * /api/meet/gettoken
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/meet/gettoken"
    
返回参考：

    * JSON::

        {
        }


.. _meet.create:
创建会议
-----------------------
接口地址：
    * /api/meet/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/meet/create"
    
返回参考：

    * JSON::

        {
        }



.. _meet.enter:
进入会议
-----------------------
接口地址：
    * /api/meet/enter
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/meet/enter"
    
返回参考：

    * JSON::

        {
        }



.. _meet.info:
获取会议信息
-----------------------
接口地址：
    * /api/meet/info/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/meet/info/:id"
    
返回参考：

    * JSON::

        {
        }