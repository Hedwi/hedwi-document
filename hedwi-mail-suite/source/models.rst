.. _help-models:

.. _models:


database table structure update
----------------------------------------------------------------------------


2024-10-30
======================


2024-11-18
======================

    chat_channel_message  add `cmd_type` field type int16 smallint
    chat_user_message     add `cmd_type` field type int16 smallint

    chat_channel   add `kanban` field type int64 bigint
    chat_channel   add `drive` field type int64 bigint

    chat_channel_rel   add `type` field type int16 smallint

