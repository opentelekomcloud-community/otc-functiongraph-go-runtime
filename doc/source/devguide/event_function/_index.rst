Event Functions
==========================

.. toctree::
   :maxdepth: 1
   :hidden:

   Create from scratch <scratch/_index>
   Container Image <container/_index>
   Trigger events <trigger_events/_index>

Event functions can be configured with event triggers and integrate
a variety of products (such as object storage service OBS,
distributed messaging service RabbitMQ version, cloud log service LTS, etc.).

Types of Event Functions in Go
------------------------------------

You can build FunctionGraph event functions in Go in the following ways:

* :ref:`From Scratch  <devguide_event_function_scratch_index>`
* :ref:`Using Container Image  <devguide_event_function_container_index>`

.. note::
  The option **Select template** is not supported for event functions written in **Go**.

FunctionGraph Go libraries
--------------------------------

The FunctionGraph Go runtime SDK provides the following libraries
to help you develop Go event functions:

- :github_repo_master:`github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context </go-runtime/go-api/context>`

  This library provides the Context structure, which contains
  the runtime information of the function and provides
  methods to obtain the runtime information.

- :github_repo_master:`github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime </go-runtime/pkg/runtime>`

  This library provides the Handler interface, which defines
  the method that must be implemented by a function handler.

- :github_repo_master:`github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-events </go-events>`

  This library provides type definitions for common event source integrations.


Open Telekom Cloud community provides following libraries for Go development:

* The community edition of :github_otc_community:`OTC SDK for API signing in Go <otc-api-sign-sdk-go>`
  provides utility methods to handle request signing.


Event Triggers
------------------------------------
Both methods support trigger events.

For available trigger events, see :ref:`Overview of Trigger Events  <devguide_event_function_trigger_events_index>`.
