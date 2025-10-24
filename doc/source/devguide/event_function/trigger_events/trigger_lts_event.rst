Log Trigger LTS
===============

You can write FunctionGraph functions to process logs subscribed to Cloud Log
Service. When Cloud Log Service collects subscribed logs, you can call
FunctionGraph functions by passing the collected logs as parameters (LTS sample
events). FunctionGraph function code can be customized, analyzed, or loaded
into other systems. For the use of LTS log triggers, please refer to
`Using a LTS Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_lts_trigger.html>`__.

Example LTS Event
-----------------

.. code-block:: json

   {
     "lts": {
       "data": "=="
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
   * - data
     - String
     - Reference examples
     - Base64 encoded data

Example
-------

.. literalinclude:: /../../samples-events/lts/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/lts/handler.go>`
