Invoking FunctionGraph Event Function from Go
==============================================

This example demonstrates how to call a FunctionGraph implemented in
python using Go.

Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^

Needed environment variables set

.. list-table:: Environment Variables
   :widths: 25 25
   :header-rows: 1

   * - Name
     - Description
   * - OTC_SDK_PROJECTID
     - Project ID
   * - OTC_TENANT_NAME
     - Region, e.g. "eu-de"
   * - OTC_SDK_AK
     - Access Key
   * - OTC_SDK_SK
     - Secret Key
   * - OTC_USER_NAME
     - User name
   * - OTC_USER_PASSWORD
     - User password
   * - OTC_DOMAIN_NAME
     - Domain name
   * - OTC_DOMAIN_NAME
     - IAM Endpoint, e.g. https://iam.eu-de.otc.t-systems.com/v3


Deployed FunctionGraph Event Function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Example FunctionGraph written in python


.. literalinclude:: /../../samples-doc/invoke-fg/functionGraph.py
   :language: python
   :caption: Example FunctionGraph: functionGraph.py

Deploy this FunctionGraph as follows:

- Type: "Event Function":
- Name: "DefaultPython3_10"
- Agency: "Use no agency"
- Runtime: "python 3.10"
- Application: "default"
- Handler: "index.handler"


Call Functiongraph using Username and Password
-----------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke_fg/src/invokeUsernamePassword.go <samples-doc/invoke_fg/src/invokeUsernamePassword.go>`.

As API calls require a token, this has to obtained by calling
the IAM token endpoint with username and password first.
(See `Obtaining a User Token <https://docs.otc.t-systems.com/identity-access-management/api-ref/apis/token_management/obtaining_a_user_token.html#en-us-topic-0057845583>`_)


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke_fg
    go run src/invokeUsernamePassword.go


Call Functiongraph using AK/SK
-----------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke_fg/src/invokeAKSK.go <samples-doc/invoke_fg/src/invokeAKSK.go>`.

Using AK/SK authentication, no token retrieval is necessary, but the
requests have to be signed with the AK/SK (or for temporal credentials with SecurityAccessKey/SecurityKey/SecurityToken).

For request signing the `otc-api-sign-sdk-go <https://github.com/opentelekomcloud-community/otc-api-sign-sdk-go>`_
can be used.


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke_fg
    go run src_aksk/invokeAKSK.go

