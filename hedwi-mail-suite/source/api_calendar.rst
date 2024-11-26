.. _api_calendar:

.. _api-calendar-example:


.. _calendar.list:

list calendar
-------------------------------------------------------------
interface address:
    * /api/calendar/collection/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/calendar/collection/list"
    
response example：

    * JSON::

        {
        }


create calendar
-------------------------------------------------------------
interface address:
    * /api/calendar/collection/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/calendar/collection/create"
    
response example：

    * JSON::

        {
        }


update calendar
-------------------------------------------------------------
interface address:
    * /api/calendar/collection/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/calendar/collection/update"
    
response example：

    * JSON::

        {
        }


list calendar event
-------------------------------------------------------------
interface address:
    * /api/calendar/event/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/calendar/event/list"
    
response example：

    * JSON::

        {
        }

get calendar event detail
----------------------------------------------------------------------------------------------------------------------------
interface address:
    * /api/calendar/event/detail/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/calendar/event/detail/:id"
    
response example：

    * JSON::

        {
        }

create calendar event
-------------------------------------------------------------
interface address:
    * /api/calendar/event/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/calendar/event/create"
    
response example：

    * JSON::

        {
        }

update calendar event
-------------------------------------------------------------
interface address:
    * /api/calendar/event/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/calendar/event/update"
    
response example：

    * JSON::

        {
        }

delete calendar event
-------------------------------------------------------------
interface address:
    * /api/calendar/event/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/calendar/event/delete"
    
response example：

    * JSON::

        {
        }

