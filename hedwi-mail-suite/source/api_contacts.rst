.. _api_contacts:

.. _api-contacts-example:

.. _contact.list:

获取联系人列表
-----------------------
接口地址：
    * /api/contact/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/contact/list"
    
返回参考：

    * JSON::

        {
        }


创建联系人
-----------------------
接口地址：
    * /api/contact/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/create"
    
返回参考：

    * JSON::

        {
        }


搜索联系人
-----------------------
接口地址：
    * /api/contact/search
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/search"
    
返回参考：

    * JSON::

        {
        }


 更新联系人信息
-----------------------
接口地址：
    * /api/contact/update
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/update"
    
返回参考：

    * JSON::

        {
        }




获取联系人详情
-----------------------
接口地址：
    * /api/contact/detail/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/contact/detail/:id"
    
返回参考：

    * JSON::

        {
        }



删除联系人
-----------------------
接口地址：
    * /api/contact/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/delete"
    
返回参考：

    * JSON::

        {
        }


联系人标签列表
-----------------------
接口地址：
    * /api/contact/label/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/contact/label/list"
    
返回参考：

    * JSON::

        {
        }

创建联系人标签
-----------------------
接口地址：
    * /api/contact/label/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/label/create"
    
返回参考：

    * JSON::

        {
        }

删除联系人标签
-----------------------
接口地址：
    * /api/contact/label/deletelabel
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/label/deletelabel"
    
返回参考：

    * JSON::

        {
        }


应用联系人标签
-----------------------
接口地址：
    * /api/contact/label/applylabel
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/label/applylabel"
    
返回参考：

    * JSON::

        {
        }


修改联系人星标状态
-----------------------
接口地址：
    * /api/contact/label/starredstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/contact/label/starredstate"
    
返回参考：

    * JSON::

        {
        }

