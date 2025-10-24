Document Database Service DDS
=============================

Using DDS triggers, each time a table in the database is updated, a
Functiongraph function can be triggered to perform additional work. For more
information about how to use DDS triggers, see `Using a DDS Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_dds_trigger.html>`__.

DDS example event
-----------------

.. code-block:: json

   {
     "records": [
       {
         "event_source": "dds",
         "event_name": "insert",
         "region": "{region}",
         "event_version": "1.0",
         "dds": {
           "size_bytes": "100",
           "token": "{\"_data\": \"825D8C2F4D0000001529295A100474039A3412A64BA89041DC952357FB4446645F696400645D8C2F8E5BECCB6CF5370D6A0004\"}",
           "full_document": "{\"_id\": {\"$oid\": \"5d8c2f8e5beccb6cf5370d6a\"},\"name\": \"dds\",\"age\": {\"$numberDouble \": \"52.0\"}}",
           "ns": "{\"db\": \"functiongraph\",\"coll\": \"person\"}"
         },
         "event_source_id": "e6065860-f7b8-4cca-80bd-24ef2a3bb748"
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
   * - region
     - String
     - {region}
     - The region where the DDS instance is located
   * - event_version
     - String
     - 1.0
     - Event protocol version
   * - event_source
     - String
     - dds
     - Source of the event
   * - event_name
     - String
     - insert
     - Event name
   * - size_bytes
     - Int
     - 100
     - The number of bytes in the message
   * - token
     - String
     - Reference examples
     - Base64 encoded data
   * - full_document
     - String
     - Reference examples
     - Complete file information
   * - ns
     - String
     - Reference examples
     - Column Name
   * - event_source_id
     - String
     - e6065860-f7b8-4cca-80bd-24ef2a3bb748
     - Event source unique identifier
     - 


Example
-------

.. literalinclude:: /../../samples-events/dds/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/dds/handler.go>`
