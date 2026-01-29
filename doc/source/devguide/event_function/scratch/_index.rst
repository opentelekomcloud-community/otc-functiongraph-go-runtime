.. _devguide_event_function_scratch_index:

Building FunctionGraph Event Functions with Go from scratch
==========================================================================

.. toctree::
   :hidden:

   Handler <handler>
   Context <context>
   Initializer <initializer>

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

- :github_repo_master:`github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events </go-runtime/events>`

  This library provides type definitions for common event source integrations.


OpenTelekomCloud community provides following libraries for Go development:

* The community edition of `OTC SDK for API signing in Go <https://github.com/opentelekomcloud-community/otc-api-sign-sdk-go>`_
  provides utility methods to handle request signing.
