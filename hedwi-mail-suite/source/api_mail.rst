.. _api_mail:

.. _api-mail-example:

.. _mail.list:

list mail
-------------------------------------------------------------
interface address:
    * /api/mail/list
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/mail/list"
    
response example：

    * JSON::

        {
        }


get mail detail
-------------------------------------------------------------
interface address:
    * /api/mail/detail/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/mail/detail"
    
response example：

    * JSON::

        {
        }


get mail original content
-------------------------------------------------------------
interface address:
    * /api/mail/original/:id
HTTP request method:
    * GET
request params:
    * -
response code:
    * common response

example::

    curl -X GET "$BASE_URL/api/mail/original/:id"
    
response example：

    * JSON::

        {
        }


search mail
-------------------------------------------------------------
interface address:
    * /api/mail/search
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/search"
    
response example：

    * JSON::

        {
        }

modify mail read state
-------------------------------------------------------------
interface address:
    * /api/mail/readstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/readstate"
    
response example：

    * JSON::

        {
        }

modify mail starred state
-------------------------------------------------------------
interface address:
    * /api/mail/starredstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/starredstate"
    
response example：

    * JSON::

        {
        }

modify mail important state
-------------------------------------------------------------
interface address:
    * /api/mail/importantstate
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/importantstate"
    
response example：

    * JSON::

        {
        }

delete mail
-------------------------------------------------------------
interface address:
    * /api/mail/delete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/delete"
    
response example：

    * JSON::

        {
        }


permanently delete mail
-------------------------------------------------------------
interface address:
    * /api/mail/foreverdelete
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/foreverdelete"
    
response example：

    * JSON::

        {
        }

archive mail
-------------------------------------------------------------
interface address:
    * /api/mail/archive
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/archive"
    
response example：

    * JSON::

        {
        }

unarchive mail
-------------------------------------------------------------
interface address:
    * /api/mail/unarchive
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/unarchive"
    
response example：

    * JSON::

        {
        }



move mail to label group
-------------------------------------------------------------
interface address:
    * /api/mail/movelabel
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/movelabel"
    
response example：

    * JSON::

        {
        }

delete mail label group
-------------------------------------------------------------
interface address:
    * /api/mail/deletelabel
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/deletelabel"
    
response example：

    * JSON::

        {
        }

apply mail label group
-------------------------------------------------------------
interface address:
    * /api/mail/applylabel
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/applylabel"
    
response example：

    * JSON::

        {
        }



send mail
-------------------------------------------------------------
interface address:
    * /api/mail/send/:domain
HTTP request method:
    * POST
request params:
    * -
response code:
    * common response

example::

    curl -X POST "$BASE_URL/api/mail/send/:domain"
    
response example：

    * JSON::

        {
        }
