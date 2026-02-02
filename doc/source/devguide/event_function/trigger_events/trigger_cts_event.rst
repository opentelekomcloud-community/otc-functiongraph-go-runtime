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
CTS triggers, please refer to :docs_otc:`Using a CTS Trigger <function-graph/umn/creating_triggers/using_a_cts_trigger.html>`.

CTS example event
-----------------

.. literalinclude:: /../../samples-doc/event-cts/resources/cts_event.json
    :language: json
    :caption: :github_repo_master:`cts_event.json <samples-doc/event-cts/resources/cts_event.json>`


Parameter description
---------------------


.. list-table::
   :header-rows: 1
   :widths: 20 15 40

   * - Parameter
     - Type
     - Description
   * - time
     - Int
     - (Epoch timestamp in milliseconds)
   * - user
     - Map
     - Information about the user who initiated this request
   * - request
     - Map
     - Event request content
   * - response
     - Map
     - Incident response content
   * - code
     - Int
     - Event response code, such as 200, 400
   * - service_type
     - String
     - Abbreviation of the sender, such as vpc, ecs, etc.
   * - resource_type
     - String
     - The sender resource type, such as vm, vpn, etc.
   * - resource_name
     - String
     - Resource name, such as the name of a virtual machine in the ecs service
   * - trace_name
     - String
     - Event name, such as: startServer, shutDown, etc.
   * - trace_type
     - String
     - The event source type, such as ApiCall
   * - record_time
     - String
     - The time when the cts service receives this trace (Epoch timestamp in milliseconds)
   * - trace_id
     - String
     - Unique identifier for the event
   * - trace_status
     - String
     - Status of the event

For full description of all parameters see
:docs_otc:`CTS Event Reference <cloud-trace-service/umn/trace_references/trace_structure.html#id1>`.

Example
-------

.. literalinclude:: /../../samples-doc/event-cts/src/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-doc/event-cts/src/handler.go>`
