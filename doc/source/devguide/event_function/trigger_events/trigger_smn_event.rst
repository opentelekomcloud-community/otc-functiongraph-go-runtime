SMN Event Source
================

Simple Message Notification (SMN) sends messages to email addresses, mobile
phones, or HTTP/HTTPS URLs. If you create a function with an SMN trigger,
messages published to a specified topic will be passed as a parameter to invoke
the function. Then, the function processes the event, for example, publishing
messages to other SMN topics or sending them to other cloud services. For
details, see :docs_otc:`Using an SMN Trigger <function-graph/umn/creating_triggers/using_an_smn_trigger.html>`.

SMN example event
-----------------

.. literalinclude:: /../../samples-doc/event-smn/resources/smn_event.json
    :language: json
    :caption: :github_repo_master:`smn_event.json <samples-doc/event-smn/resources/smn_event.json>`


Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 25 15 35

   * - Parameter
     - Type
     - Description
   * - event_version
     - String
     - Event version
   * - topic_urn
     - String
     - ID of an SMN event
   * - timestamp
     - String
     - Time when an event occurs
   * - message_attributes
     - Map
     - Message attributes
   * - message
     - String
     - Message content
   * - type
     - String
     - Event type
   * - message_id
     - String
     - Message ID. The ID of each message is unique.
   * - subject
     - String
     - Subject of message
   * - event_subscription_urn
     - String
     - Subscription ID
   * - event_source
     - String
     - Event source: **smn**

Example
-------

.. literalinclude:: /../../samples-doc/event-smn/src/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-doc/event-smn/src/handler.go>`
