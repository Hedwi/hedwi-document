.. _api_drive:

.. _api-drive-example:


.. _drive-project.List:

获取文件夹列表
-----------------------
接口地址：
    * /api/drive/dirs
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/dirs"
    
返回参考：

    * JSON::

        {
        }



.. _drive.filelist:

获取文件列表
-----------------------
接口地址：
    * /api/drive/filelist
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/filelist"
    
返回参考：

    * JSON::

        {
        }

.. _drive.sharedlist:

获取共享文件列表
-----------------------
接口地址：
    * /api/drive/sharedlist
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/sharedlist"
    
返回参考：

    * JSON::

        {
        }

.. _drive.trashlist:

获取回收站文件列表
-----------------------
接口地址：
    * /api/drive/trash
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/trash"
    
返回参考：

    * JSON::

        {
        }

.. _drive.starredlist:

获取星标文件列表
-----------------------
接口地址：
    * /api/drive/starredlist
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/starredlist"
    
返回参考：

    * JSON::

        {
        }

.. _drive.recentlist:

获取最近访问文件列表
-----------------------
接口地址：
    * /api/drive/recentlist
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/recentlist"
    
返回参考：

    * JSON::

        {
        }

.. _drive.recentdocs:

获取最近Office文档列表
-----------------------
接口地址：
    * /api/drive/recentdocs
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/recentdocs"
    
返回参考：

    * JSON::

        {
        }

.. _drive.starreddocs:

获取星标Office文档列表
-----------------------
接口地址：
    * /api/drive/starreddocs
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/starreddocs"
    
返回参考：

    * JSON::

        {
        }

.. _drive.shareddocs:

获取共享Office文档列表
-----------------------
接口地址：
    * /api/drive/shareddocs
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/shareddocs"
    
返回参考：

    * JSON::

        {
        }

.. _drive.docslist:

获取Office文件列表
-----------------------
接口地址：
    * /api/drive/docslist
HTTP 请求方式：
    * GET
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X GET "$BASE_URL/api/drive/docslist"
    
返回参考：

    * JSON::

        {
        }

.. _drive.touchlist:

访问文件
-----------------------
接口地址：
    * /api/drive/touch
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/touch"
    
返回参考：

    * JSON::

        {
        }

.. _drive.upload:

上传文件
-----------------------
接口地址：
    * /api/drive/upload
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/upload"
    
返回参考：

    * JSON::

        {
        }

.. _drive.create:

创建文件/文件夹
-----------------------
接口地址：
    * /api/drive/create
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/create"
    
返回参考：

    * JSON::

        {
        }


.. _drive.move:

移动文件/文件夹
-----------------------
接口地址：
    * /api/drive/move
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/move"
    
返回参考：

    * JSON::

        {
        }



.. _drive.delete:

删除文件/文件夹
-----------------------
接口地址：
    * /api/drive/delete
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/delete"
    
返回参考：

    * JSON::

        {
        }



.. _drive.rename:

重命名文件/文件夹
-----------------------
接口地址：
    * /api/drive/rename
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/rename"
    
返回参考：

    * JSON::

        {
        }



.. _drive.downloadlink:

获取下载链接
-----------------------
接口地址：
    * /api/drive/downloadlink
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/downloadlink"
    
返回参考：

    * JSON::

        {
        }



.. _drive.clone:

克隆文件/文件夹
-----------------------
接口地址：
    * /api/drive/clone
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/clone"
    
返回参考：

    * JSON::

        {
        }



.. _drive.starred:

星标文件/文件夹
-----------------------
接口地址：
    * /api/drive/starred
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/starred"
    
返回参考：

    * JSON::

        {
        }



.. _drive.restore:

恢复文件/文件夹
-----------------------
接口地址：
    * /api/drive/restore
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/restore"
    
返回参考：

    * JSON::

        {
        }



.. _drive.search:

搜索文件/文件夹
-----------------------
接口地址：
    * /api/drive/search
HTTP 请求方式：
    * POST
请求参数：
    * -
响应代码：
    * 共通返回

示例::

    curl -X POST "$BASE_URL/api/drive/search"
    
返回参考：

    * JSON::

        {
        }


