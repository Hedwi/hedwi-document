.. _api_kanban:

.. _api-kanban-example:


.. _kanban-project.List:

列出看板项目
-----------------------
接口地址：
    * /api/kanban/project/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/kanban/project/list"
    
返回参考：

    * JSON::

        {
            "code": 0,
            "message": "",
            "data": {
                "list": [
                    {
                       "name":       "",
                        "id":         "",
                        "creator_id": "",
                        "team_id":    "",
                        "color":      "",
                        "icon":       "",
                        "icon_color": "",
                        "name_color": "",
                    },
                ], 
                "uid": "1234567890", 
                "email": "hello@example.com", 
            }, 
        }


.. _kanban-project.create:

创建看板项目
-----------------------
接口地址：
    * /api/kanban/project/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/project/create"
    
返回参考：

    * JSON::

        {
        }


.. _kanban-project.move:

移动看板项目
-----------------------
接口地址：
    * /api/kanban/project/move
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/project/move"
    
返回参考：

    * JSON::

        {
        }



.. _kanban-project.update:

更新看板项目
-----------------------
接口地址：
    * /api/kanban/project/update
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/project/update"
    
返回参考：

    * JSON::

        {
        }



删除看板项目
-----------------------
接口地址：
    * /api/kanban/project/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/project/delete"
    
返回参考：

    * JSON::

        {
        }



.. _kanban-list.create:

创建看板列表
-----------------------
接口地址：
    * /api/kanban/list/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/list/create"
    
返回参考：

    * JSON::

        {
        }



移动看板列表
-----------------------
接口地址：
    * /api/kanban/list/move
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/list/move"
    
返回参考：

    * JSON::

        {
        }



更新看板列表
-----------------------
接口地址：
    * /api/kanban/list/update
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/project/update"
    
返回参考：

    * JSON::

        {
        }



删除看板列表
-----------------------
接口地址：
    * /api/kanban/list/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/list/delete"
    
返回参考：

    * JSON::

        {
        }



获取任务列表
-----------------------
接口地址：
    * /api/kanban/card/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/kanban/card/list"
    
返回参考：

    * JSON::

        {
        }



创建任务
-----------------------
接口地址：
    * /api/kanban/card/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/card/create"
    
返回参考：

    * JSON::

        {
        }


获取任务详情
-----------------------
接口地址：
    * /api/kanban/card/detail/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/kanban/card/detail/:id"
    
返回参考：

    * JSON::

        {
        }


移动任务
-----------------------
接口地址：
    * /api/kanban/card/move
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/card/move"
    
返回参考：

    * JSON::

        {
        }



更新任务
-----------------------
接口地址：
    * /api/kanban/card/update
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/card/update"
    
返回参考：

    * JSON::

        {
        }



删除任务
-----------------------
接口地址：
    * /api/kanban/card/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/kanban/card/delete"
    
返回参考：

    * JSON::

        {
        }
