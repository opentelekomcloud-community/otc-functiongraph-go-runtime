DMS for Kafka (OPENSOURCE) Event Source
=======================================

DMS for Kafka is a message queuing service that provides Kafka premium
instances. If you create a Kafka trigger for a function, when a message is sent
to a Kafka instance topic, FunctionGraph will retrieve the message and trigger
the function to perform other operations. For details, see
`Using a Kafka Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_kafka_trigger.html>`__.

Kafka example event
-------------------

.. literalinclude:: /../../samples-doc/event-dms4kafka/resources/dms4kafka_event.json
    :language: json
    :caption: :github_repo_master:`dms4kafka_event.json <samples-doc/event-dms4kafka/resources/dms4kafka_event.json>`



Parameter description
---------------------


.. list-table::
   :header-rows: 1
   :widths: 20 15 40

   * - Parameter
     - Type
     - Description
   * - event_version
     - String
     - Event version
   * - event_time
     - String
     - Time when an event occurs
   * - trigger_type
     - String
     - Event type: **KAFKA**
   * - region
     - String
     - Region where a Kafka instance resides
   * - instance_id
     - String
     - Kafka instance ID
   * - messages
     - String[]
     - Message content
   * - topic_id
     - String
     - Message ID

Example
-------

.. literalinclude:: /../../samples-doc/event-dms4kafka/src/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-doc/event-dms4kafka/src/handler.go>`