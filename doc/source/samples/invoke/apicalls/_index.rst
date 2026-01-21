Calling FunctionGraph using API calls
==================================================

Prerequisites
-------------

See :ref:`ref_invoke-prerequisites`

Call Functiongraph using Username and Password synchronously
------------------------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg/src/invokeSync_UsernamePassword.go <samples-doc/invoke-fg/src/invokeSync_UsernamePassword.go>`.

As API calls require a token, this has to obtained by calling
the IAM token endpoint with username and password first.
(See `Obtaining a User Token <https://docs.otc.t-systems.com/identity-access-management/api-ref/apis/token_management/obtaining_a_user_token.html#en-us-topic-0057845583>`_)


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

For request signing the `otc-api-sign-sdk-go <https://github.com/opentelekomcloud-community/otc-api-sign-sdk-go>`_
can be used.


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke-fg/src
    go test -run TestInvokeSync_AKSK
