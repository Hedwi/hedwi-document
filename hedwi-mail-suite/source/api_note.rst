.. _api_note:

.. _api-note-example:


.. _note.list:

list note
-------------------------------------------------------------
interface address:
    * /api/note/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/note/list"
    
response example：

    * JSON::

        {
        }




.. _note.create:

create note
-------------------------------------------------------------
interface address:
    * /api/note/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/create"
    
response example：

    * JSON::

        {
        }


.. _note.detail:

get note detail
-------------------------------------------------------------
interface address:
    * /api/note/detail/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/note/detail/:id"
    
response example：

    * JSON::

        {
        }


.. _note.delete:

delete note
-------------------------------------------------------------
interface address:
    	* /api/note/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/delete"
    
response example：

    * JSON::

        {
        }




.. _note.archive:

archive note
-------------------------------------------------------------
interface address:
    * /api/note/archive
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/archive"
    
response example：

    * JSON::

        {
        }


.. _note.unarchive:

unarchive note
-------------------------------------------------------------
interface address:
    * /api/note/unarchive
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/unarchive"
    
response example：

    * JSON::

        {
        }

.. _note.starredstate:

modify note starred state
-------------------------------------------------------------
interface address:
    * /api/note/starredstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/starredstate"
    
response example：

    * JSON::

        {
        }


.. _note.importantstate:

modify note important state
-------------------------------------------------------------
interface address:
    * /api/note/importantstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/note/importantstate"
    
response example：

    * JSON::

        {
        }
