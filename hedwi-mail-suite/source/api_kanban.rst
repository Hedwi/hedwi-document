.. _api_kanban:

.. _api-kanban-example:

Kanban API
===================== 

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
