DIS Event Source
================

Data Ingestion Service (DIS) can ingest large amounts of data in real time. You
can create a function to automatically poll a DIS stream and process all new
data records, such as website click streams, financial transactions, social
media streams, IT logs, and location-tracking events. FunctionGraph
periodically polls the stream for new data records. For details, see
`Using a DIS Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_dis_trigger.html>`__.

DIS example event
-----------------

.. code-block:: json

   {
     "ShardID": "shardId-0000000000",
     "Message": {
       "next_partition_cursor": "eyJnZXRJdGVyYXRvclBhcmFtIjp7InN0cmVhbS1uYW1lIjoiZGlzLXN3dGVzdCIsInBhcnRpdGlvbi1pZCI6InNoYXJkSWQtMDAwMDAwMDAwMCIsImN1cnNvci10eXBlIjoiVFJJTV9IT1JJWk9OIiwic3RhcnRpbmctc2VxdWVuY2UtbnVtYmVyIjoiNCJ9LCJnZW5lcmF0ZVRpbWVzdGFtcCI6MTUwOTYwNjM5MjE5MX0",
       "records": [
         {
           "partition_key": "shardId_0000000000",
           "data": "d2VsY29tZQ==",
           "sequence_number": "0"
         },
         {
           "partition_key": "shardId_0000000000",
           "data": "dXNpbmc=",
           "sequence_number": "1"
         },
         {
           "partition_key": "shardId_0000000000",
           "data": "RnVuY3Rpb25TdGFnZQ==",
           "sequence_number": "2"
         },
         {
           "partition_key": "shardId_0000000000",
           "data": "c2VydmljZQ==",
           "sequence_number": "3"
         }
       ],
       "millis_behind_latest": ""
     },
     "Tag": "latest",
     "StreamName": "dis-swtest"
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
   * - ShardID
     - String
     - shardId-0000000000
     - Partition ID
   * - next_partition_cursor
     - String
     - See the example.
     - Next partition cursor
   * - Records
     - Map
     - See the example.
     - Data records stored in a DIS stream
   * - partition_key
     - String
     - See the example.
     - Partition key
   * - data
     - String
     - See the example.
     - Data blocks, which are added by the data producer to the stream
   * - sequence_number
     - Int
     - See the example.
     - Record ID, which is automatically allocated by DIS
   * - Tag
     - String
     - latest
     - Stream tag
   * - StreamName
     - String
     - dis-swtest
     - Stream name

Example
-------

.. literalinclude:: /../../samples-events/dis/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/dis/handler.go>`
