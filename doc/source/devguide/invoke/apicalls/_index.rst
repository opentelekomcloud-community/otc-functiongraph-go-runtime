Calling FunctionGraph using API calls
==================================================

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg <samples-doc/invoke-fg>`.

For details on API calls, see:

- :docs_otc:`OpenTelekomCloud API Documentation <function-graph/api-ref/api/function_invocation/index.html>`.

Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^
See environment variables section in: :ref:`ref_invoke-prerequisites`

Deployed FunctionGraph Event Function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Method A: Deploy using console
""""""""""""""""""""""""""""""""""""

Deploy the function manually as described in :ref:`ref_invoke-prerequisites`
using the FunctionGraph console.

Method B: Deploy using API calls
""""""""""""""""""""""""""""""""""""

Or deploy the function using the provided Go code using API calls.

.. code-block:: bash

    cd samples-doc/invoke-fg/src
    go test -run TestCreateFunction_UsernamePassword

this will create function with code and parameters from:

.. literalinclude:: /../../samples-doc/invoke-fg/src/functionConfig.go
   :language: go
   :caption: samples-doc/invoke-fg/src/functionConfig.go

After testing, you can delete the function using:

.. code-block:: bash

    cd samples-doc/invoke-fg/src
    go test -run TestDeleteFunction_UsernamePassword

Call Functiongraph using Username and Password synchronously
------------------------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg/src/invokeSync_UsernamePassword.go <samples-doc/invoke-fg/src/invokeSync_UsernamePassword.go>`.

As API calls require a token, this has to obtained by calling
the IAM token endpoint with username and password first.
(See :docs_otc:`Obtaining a User Token <identity-access-management/api-ref/apis/token_management/obtaining_a_user_token.html#en-us-topic-0057845583>`)


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke-fg/src
    go test -run TestInvokeSync_UsernamePassword

Call Functiongraph using AK/SK synchronously
-----------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg/src/invokeSync_AKSK.go <samples-doc/invoke-fg/src/invokeSync_AKSK.go>`.

Using AK/SK authentication, no token retrieval is necessary, but the
requests have to be signed with the AK/SK (or for temporal credentials with SecurityAccessKey/SecurityKey/SecurityToken).

For request signing the :github_otc_community:`otc-api-sign-sdk-go <otc-api-sign-sdk-go>`
can be used.


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke-fg/src
    go test -run TestInvokeSync_AKSK

.. note::

   Don't forget to clean up the deployed resources after testing.
