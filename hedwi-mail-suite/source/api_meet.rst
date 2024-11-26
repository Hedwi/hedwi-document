.. _api_meet:

.. _api-meet-example:


.. _meet.gettoken:
get join meet token
-------------------------------------------------------------
interface address:
    * /api/meet/gettoken
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/meet/gettoken"
    
response example：

    * JSON::

        {
        }


.. _meet.create:
create meet
-------------------------------------------------------------
interface address:
    * /api/meet/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/meet/create"
    
response example：

    * JSON::

        {
        }



.. _meet.enter:
enter meet
-------------------------------------------------------------
interface address:
    * /api/meet/enter
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/meet/enter"
    
response example：

    * JSON::

        {
        }



.. _meet.info:
get meet info
-------------------------------------------------------------
interface address:
    * /api/meet/info/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/meet/info/:id"
    
response example：

    * JSON::

        {
        }