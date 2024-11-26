.. _api_contacts:

.. _api-contacts-example:

.. _contact.list:

list contact
-------------------------------------------------------------
interface address:
    * /api/contact/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/contact/list"
    
response example：

    * JSON::

        {
        }


create contact
-------------------------------------------------------------
interface address:
    * /api/contact/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/create"
    
response example：

    * JSON::

        {
        }


search contact
-------------------------------------------------------------
interface address:
    * /api/contact/search
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/search"
    
response example：

    * JSON::

        {
        }


update contact
-------------------------------------------------------------
interface address:
    * /api/contact/update
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/update"
    
response example：

    * JSON::

        {
        }




get contact detail
-------------------------------------------------------------
interface address:
    * /api/contact/detail/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/contact/detail/:id"
    
response example：

    * JSON::

        {
        }



delete contact
-------------------------------------------------------------
interface address:
    * /api/contact/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/delete"
    
response example：

    * JSON::

        {
        }


list contact label
-------------------------------------------------------------
interface address:
    * /api/contact/label/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/contact/label/list"
    
response example：

    * JSON::

        {
        }

create contact label
-------------------------------------------------------------
interface address:
    * /api/contact/label/create
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/label/create"
    
response example：

    * JSON::

        {
        }

delete contact label
-------------------------------------------------------------
interface address:
    * /api/contact/label/deletelabel
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/label/deletelabel"
    
response example：

    * JSON::

        {
        }


apply contact label
-------------------------------------------------------------
interface address:
    * /api/contact/label/applylabel
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/label/applylabel"
    
response example：

    * JSON::

        {
        }


modify contact starred state
-------------------------------------------------------------
interface address:
    * /api/contact/label/starredstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/contact/label/starredstate"
    
response example：

    * JSON::

        {
        }

