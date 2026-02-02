Log Trigger LTS
===============

You can write FunctionGraph functions to process logs subscribed to Cloud Log
Service. When Cloud Log Service collects subscribed logs, you can call
FunctionGraph functions by passing the collected logs as parameters (LTS sample
events). FunctionGraph function code can be customized, analyzed, or loaded
into other systems. For the use of LTS log triggers, please refer to
:docs_otc:`Using a LTS Trigger <function-graph/umn/creating_triggers/using_an_lts_trigger.html>`.

Example LTS Event
-----------------

.. literalinclude:: /../../samples-doc/event-lts/resources/lts_event.json
    :language: json
    :caption: :github_repo_master:`lts_event.json <samples-doc/event-lts/resources/lts_event.json>`


Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 20 15 40

   * - Parameter
     - Type
     - Description
   * - data
     - String
     - Base64 encoded data

Example
-------

.. literalinclude:: /../../samples-doc/event-lts/src/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-doc/event-lts/src/handler.go>`