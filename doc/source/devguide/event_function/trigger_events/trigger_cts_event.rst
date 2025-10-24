Cloud Trace Service Trigger CTS
===============================

You can write a FunctionGraph function. According to the CTS cloud audit
service type and the event notification required for the operation subscription,
when the CTS cloud audit service obtains the subscribed operation record, the
collected operation record is passed as a parameter (CTS sample event) through
the CTS trigger to call the FunctionGraph function. Through the function, the
key information in the log is analyzed and processed, and the system, network
and other business modules are automatically repaired, or alarms are generated
through SMS, email, etc. to notify business personnel to handle. For the use of
CTS triggers, please refer to `Using a CTS Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_cts_trigger.html>`__.

CTS example event
-----------------

.. code-block:: json

   {
     "cts": {
       "time": 1529996047000,
       "user": {
         "name": "userName",
         "id": "5b726c4fbfd84821ba866bafaaf56aax",
         "domain": {
             "name": "domainName",
             "id": "b2b3853af40448fcb9e40dxj89505ba"
         }
       },
       "request": {},
       "response": {},
       "code": 204,
       "service_type": "vpc",
       "resource_type": "VPC",
       "resource_name": "workflow-2be1",
       "resource_id": "urn:fgs:{region}:2d1d891d93054bbaa69b9e866c0971ac:graph:workflow-2be1",
       "trace_name": "deleteGraph",
       "trace_type": "ConsoleAction",
       "record_time": 1529996047000,
       "trace_id": "69be64a7-0233-11e8-82e4-e5d37911193e",
       "trace_status": "normal"
     }
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
   * - time
     - Int
     - 1529996047000
     - (Epoch timestamp in milliseconds)
   * - user
     - Map
     - Reference examples
     - Information about the user who initiated this request
   * - request
     - Map
     - Reference examples
     - Event request content
   * - response
     - Map
     - Reference examples
     - Incident response content
   * - code
     - Int
     - 204
     - Event response code, such as 200, 400
   * - service_type
     - String
     - vpc
     - Abbreviation of the sender, such as vpc, ecs, etc.
   * - resource_type
     - String
     - VPC
     - The sender resource type, such as vm, vpn, etc.
   * - resource_name
     - String
     - workflow-2be1
     - Resource name, such as the name of a virtual machine in the ecs service
   * - trace_name
     - String
     - deleteGraph
     - Event name, such as: startServer, shutDown, etc.
   * - trace_type
     - String
     - ConsoleAction
     - The event source type, such as ApiCall
   * - record_time
     - String
     - 1529996047000
     - The time when the cts service receives this trace (Epoch timestamp in milliseconds)
   * - trace_id
     - String
     - 69be64a7-0233-11e8-82e4-e5d37911193e
     - Unique identifier for the event
   * - trace_status
     - String
     - normal
     - Status of the event

Example
-------

.. literalinclude:: /../../samples-events/cts/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/cts/handler.go>`
