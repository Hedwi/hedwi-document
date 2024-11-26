.. _api_kanban:

.. _api-kanban-example:


.. _kanban-project.List:

list kanban project
-------------------------------------------------------------
interface address:
    * /api/kanban/project/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/kanban/project/list"
    
response example：

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

create kanban project
-------------------------------------------------------------
interface address:
    * /api/kanban/project/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/project/create"
    
response example：

    * JSON::

        {
        }


.. _kanban-project.move:

move kanban project
-------------------------------------------------------------
interface address:
    * /api/kanban/project/move
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/project/move"
    
response example：

    * JSON::

        {
        }



.. _kanban-project.update:

update kanban project
-------------------------------------------------------------
interface address:
    * /api/kanban/project/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/project/update"
    
response example：

    * JSON::

        {
        }



delete kanban project
-------------------------------------------------------------
interface address:
    * /api/kanban/project/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/project/delete"
    
response example：

    * JSON::

        {
        }



.. _kanban-list.create:

create kanban list
-------------------------------------------------------------
interface address:
    * /api/kanban/list/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/list/create"
    
response example：

    * JSON::

        {
        }



move kanban list
-------------------------------------------------------------
interface address:
    * /api/kanban/list/move
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/list/move"
    
response example：

    * JSON::

        {
        }



update kanban list
-------------------------------------------------------------
interface address:
    * /api/kanban/list/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/project/update"
    
response example：

    * JSON::

        {
        }



delete kanban list
-------------------------------------------------------------
interface address:
    * /api/kanban/list/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/list/delete"
    
response example：

    * JSON::

        {
        }



get kanban card list
-------------------------------------------------------------
interface address:
    * /api/kanban/card/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/kanban/card/list"
    
response example：

    * JSON::

        {
        }



create kanban card
-------------------------------------------------------------
interface address:
    * /api/kanban/card/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/card/create"
    
response example：

    * JSON::

        {
        }


get kanban card detail
-------------------------------------------------------------
interface address:
    * /api/kanban/card/detail/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/kanban/card/detail/:id"
    
response example：

    * JSON::

        {
        }


move kanban card
-------------------------------------------------------------
interface address:
    * /api/kanban/card/move
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/card/move"
    
response example：

    * JSON::

        {
        }



update kanban card
-------------------------------------------------------------
interface address:
    * /api/kanban/card/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/card/update"
    
response example：

    * JSON::

        {
        }



delete kanban card
-------------------------------------------------------------
interface address:
    * /api/kanban/card/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/kanban/card/delete"
    
response example：

    * JSON::

        {
        }
