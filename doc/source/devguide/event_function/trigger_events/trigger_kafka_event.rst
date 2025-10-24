DMS for Kafka (OPENSOURCE) Event Source
=======================================

DMS for Kafka is a message queuing service that provides Kafka premium
instances. If you create a Kafka trigger for a function, when a message is sent
to a Kafka instance topic, FunctionGraph will retrieve the message and trigger
the function to perform other operations. For details, see
`Using a Kafka Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_kafka_trigger.html>`__.

Kafka example event
-------------------

.. code-block:: json

   {
     "event_version": "v1.0",
     "event_time": 1576737962,
     "trigger_type": "KAFKA",
     "region": "{region}",
     "instance_id": "81335d56-b9fe-4679-ba95-7030949cc76b",
     "records": [
       {
         "messages": [
           "kafka message1",
           "kafka message2",
           "kafka message3",
           "kafka message4",
           "kafka message5"
         ],
         "topic_id": "topic-test"
       }
     ]
   }

Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 20 15 25 40

   * - Parameter
     - Type
     - Example Value
     - Description
   * - event_version
     - String
     - v1.0
     - Event version
   * - event_time
     - String
     - 2018-01-09T07:50:50.028Z
     - Time when an event occurs
   * - trigger_type
     - String
     - KAFKA
     - Event type
   * - region
     - String
     - {region}
     - Region where a Kafka instance resides
   * - instance_id
     - String
     - 81335d56-b9fe-4679-ba95-7030949cc76b
     - Kafka instance ID
   * - messages
     - String[]
     - See the example.
     - Message content
   * - topic_id
     - String
     - topic-test
     - Message ID

Example
-------

.. literalinclude:: /../../samples-events/kafka/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/kafka/handler.go>`
