.. _help-models:

.. _models:


database table structure update
----------------------------------------------------------------------------

2025-10-23
======================

message id change to bigint type

..  csv-table:: 
    :header: Table, Action, Field name, Field type
    :widths: 40, 20, 20, 20

	"chat_channel_message", "update", "id", "bigint" 
	"chat_user_message", "update", "id", "bigint"


2024-11-18
======================

..  csv-table:: 
    :header: Table, Action, Field name, Field type
    :widths: 40, 20, 20, 20

	"chat_channel_message", "add", "cmd_type", "smallint"
	"chat_user_message", "add", "cmd_type", "smallint"
	"chat_channel", "add", "kanban", "bigint"
	"chat_channel", "add", "drive", "bigint"
	"chat_channel_rel", "add", "type", "smallint"

