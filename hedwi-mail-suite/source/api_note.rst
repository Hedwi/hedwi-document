.. _api_note:

.. _api-note-example:


.. _note.list:

列出笔记
-----------------------
接口地址：
    * /api/note/list
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/note/list"
    
返回参考：

    * JSON::

        {
        }




.. _note.create:

创建笔记
-----------------------
接口地址：
    * /api/note/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/create"
    
返回参考：

    * JSON::

        {
        }


.. _note.detail:

获取笔记详情
-----------------------
接口地址：
    * /api/note/detail/:id
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/note/detail/:id"
    
返回参考：

    * JSON::

        {
        }


.. _note.delete:

删除笔记
-----------------------
接口地址：
    	* /api/note/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/delete"
    
返回参考：

    * JSON::

        {
        }




.. _note.archive:

归档笔记
-----------------------
接口地址：
    * /api/note/archive
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/archive"
    
返回参考：

    * JSON::

        {
        }


.. _note.unarchive:

取消归档笔记
-----------------------
接口地址：
    * /api/note/unarchive
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/unarchive"
    
返回参考：

    * JSON::

        {
        }

.. _note.starredstate:

设置笔记星标状态
-----------------------
接口地址：
    * /api/note/starredstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/starredstate"
    
返回参考：

    * JSON::

        {
        }


.. _note.importantstate:

设置笔记重要状态
-----------------------
接口地址：
    * /api/note/importantstate
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/note/importantstate"
    
返回参考：

    * JSON::

        {
        }
