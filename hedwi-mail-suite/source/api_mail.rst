.. _api_mail:

.. _api-mail-example:

.. _mail.list:

获取邮件列表
-----------------------
接口地址：
    * /api/mail/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/mail/list"
    
返回参考：

    * JSON::

        {
        }


获取邮件详情
-----------------------
接口地址：
    * /api/mail/detail/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/mail/detail"
    
返回参考：

    * JSON::

        {
        }


获取邮件原始内容
-----------------------
接口地址：
    * /api/mail/original/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/mail/original/:id"
    
返回参考：

    * JSON::

        {
        }


搜索邮件
-----------------------
接口地址：
    * /api/mail/search
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/search"
    
返回参考：

    * JSON::

        {
        }

修改邮件阅读状态
-----------------------
接口地址：
    * /api/mail/readstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/readstate"
    
返回参考：

    * JSON::

        {
        }

修改邮件星标状态
-----------------------
接口地址：
    * /api/mail/starredstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/starredstate"
    
返回参考：

    * JSON::

        {
        }

修改邮件是否重要状态
-----------------------
接口地址：
    * /api/mail/importantstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/importantstate"
    
返回参考：

    * JSON::

        {
        }

删除邮件
-----------------------
接口地址：
    * /api/mail/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/delete"
    
返回参考：

    * JSON::

        {
        }


永久删除邮件
-----------------------
接口地址：
    * /api/mail/foreverdelete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/foreverdelete"
    
返回参考：

    * JSON::

        {
        }

归档邮件
-----------------------
接口地址：
    * /api/mail/archive
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/archive"
    
返回参考：

    * JSON::

        {
        }

取消归档邮件
-----------------------
接口地址：
    * /api/mail/unarchive
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/unarchive"
    
返回参考：

    * JSON::

        {
        }



移动邮件到标签分组
-----------------------
接口地址：
    * /api/mail/movelabel
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/movelabel"
    
返回参考：

    * JSON::

        {
        }

删除邮件标签分组
-----------------------
接口地址：
    * /api/mail/deletelabel
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/deletelabel"
    
返回参考：

    * JSON::

        {
        }

应用邮件标签分组
-----------------------
接口地址：
    * /api/mail/applylabel
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/applylabel"
    
返回参考：

    * JSON::

        {
        }



发送邮件
-----------------------
接口地址：
    * /api/mail/send/:domain
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/mail/send/:domain"
    
返回参考：

    * JSON::

        {
        }
