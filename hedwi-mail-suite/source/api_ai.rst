.. _api_ai:

.. _api-ai-example:


.. _ai.List:

列出对话列表
-----------------------
接口地址：
    * /api/aichat/thread/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/aichat/thread/list"
    
返回参考：

    * JSON::

        {
        }



创建对话
-----------------------
接口地址：
    * /api/aichat/thread/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/thread/create"
    
返回参考：

    * JSON::

        {
        }

列出对话列表
-----------------------
接口地址：
    * /api/aichat/thread/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/aichat/thread/list"
    
返回参考：

    * JSON::

        {
        }

移动对话
-----------------------
接口地址：
    * /api/aichat/thread/move
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/thread/move"
    
返回参考：

    * JSON::

        {
        }

更新对话
-----------------------
接口地址：
    * /api/aichat/thread/update
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/thread/update"
    
返回参考：

    * JSON::

        {
        }

删除对话
-----------------------
接口地址：
    * /api/aichat/thread/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/thread/delete"
    
返回参考：

    * JSON::

        {
        }

列出对话消息
-----------------------
接口地址：
    * /api/aichat/content/list/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/aichat/content/list/:id"
    
返回参考：

    * JSON::

        {
        }


创建对话消息
-----------------------
接口地址：
    * /api/aichat/content/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/content/create"
    
返回参考：

    * JSON::

        {
        }


删除对话消息
-----------------------
接口地址：
    * /api/aichat/content/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/aichat/content/delete"
    
返回参考：

    * JSON::

        {
        }


对话消息流
-----------------------
接口地址：
    * /api/aichat/stream/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/aichat/stream/:id"
    
返回参考：

    * JSON::

        {
        }
