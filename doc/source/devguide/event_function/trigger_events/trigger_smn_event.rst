SMN Event Source
================

Simple Message Notification (SMN) sends messages to email addresses, mobile
phones, or HTTP/HTTPS URLs. If you create a function with an SMN trigger,
messages published to a specified topic will be passed as a parameter to invoke
the function. Then, the function processes the event, for example, publishing
messages to other SMN topics or sending them to other cloud services. For
details, see `Using an SMN Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_smn_trigger.html>`__.

SMN example event
-----------------

.. code-block:: json

   {
     "record": [
       {
         "event_version": "1.0",
         "smn": {
           "topic_urn": "urn:smn:{region}:0162c0f220284698b77a3d264376343a:{function_name}",
           "timestamp": "2018-01-09T07:11:40Z",
           "message_attributes": "None",
           "message": "this is smn message content",
           "type": "notification",
           "message_id": "a51671f77d4a479cacb09e2cd591a983",
           "subject": "this is smn message subject"
         },
         "event_subscription_urn": "urn:fss:{region}:0162c0f220284698b77a3d264376343a:function:default:read-smn-message:latest",
         "event_source": "smn"
       }
     ]
   }

Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 25 15 25 35

   * - Parameter
     - Type
     - Example Value
     - Description
   * - event_version
     - String
     - 1.0
     - Event version
   * - topic_urn
     - String
     - See the example.
     - ID of an SMN event
   * - timestamp
     - String
     - 2018-01-09T07:11:40Z
     - Time when an event occurs
   * - message_attributes
     - Map
     - None
     - Message attributes
   * - message
     - String
     - this is smn message content
     - Message content
   * - type
     - String
     - notification
     - Event type
   * - message_id
     - String
     - a51671f77d4a479cacb09e2cd591a983
     - Message ID. The ID of each message is unique.
   * - subject
     - String
     - this is smn message subject
     - Subject of message
   * - event_subscription_urn
     - String
     - See the example.
     - Subscription ID
   * - event_source
     - String
     - smn
     - Event source

Example
-------

.. literalinclude:: /../../samples-events/smn/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/smn/handler.go>`
