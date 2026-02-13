Calling FunctionGraph using curl and bash
==================================================

This section demonstrates how to call a FunctionGraph implemented in
python using curl.

For details on API calls, see:

- :docs_otc:`OpenTelekomCloud API Documentation <function-graph/api-ref/api/function_invocation/index.html>`.
- :docs_otc:`Authentication <identity-access-management/api-ref/calling_apis/authentication.html#iam-02-0510>`.
- :docs_otc:`Obtaining a User Token Through Password Authentication <identity-access-management/api-ref/apis/token_management/obtaining_a_user_token_through_password_authentication.html>`.


Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^
See environment variables section in: :ref:`ref_invoke-prerequisites`

Deployed FunctionGraph Event Function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
Deploy the function manually as described in :ref:`ref_invoke-prerequisites`
using the FunctionGraph console.

Getting Token from Username and Password
------------------------------------------------------------

To get a token for authentication from Username and Password,
you can use the provided bash script:

.. literalinclude:: /../../utils/tokenFromUsername.sh
   :language: bash
   :caption: utils/tokenFromUsername.sh

This file needs execution permissions, e.g. set using:

.. code-block:: bash

    chmod +x tokenFromUsername.sh

Use the script as follows to get the token and store it in a variable:

Option 1: Using subshell to get token into variable
------------------------------------------------------------

.. code-block:: bash

    # execute the script to get the token without output to stdout
    OTC_X_AUTH_TOKEN=$(./tokenFromUsername.sh) > /dev/null

    # Output will be like:
    echo $OTC_X_AUTH_TOKEN

    # Output example:
    MIIGBQYJKoZIhvcNAQcCoIIF9jCCBfICAQExDTALBglghkgBZQMEAgEwggOKBgkqhkiG.....

Option 2: Sourcing the script to get token into environment variable
---------------------------------------------------------------------

.. code-block:: bash

   # source the script to get the token into environment variable OTC_X_AUTH_TOKEN
   source ./tokenFromUsername.sh > /dev/null

   # Output will be like:
   echo $OTC_X_AUTH_TOKEN

   # Output example:
   MIIGBQYJKoZIhvcNAQcCoIIF9jCCBfICAQExDTALBglghkgBZQMEAgEwggOKBgkqhkiG.....

Calling Functiongraph using Username and Password synchronously
----------------------------------------------------------------

See :docs_otc:`Executing a Function Synchronously <function-graph/api-ref/api/function_invocation/executing_a_function_synchronously.html#functiongraph-06-0125>`

.. code-block:: bash

   export MY_FUNCTION_NAME="DefaultPython3_10_From_Go_SDK"
   export MY_FUNCTION_URN="urn:fss:${OTC_SDK_REGION}:${OTC_SDK_PROJECTID}:function:default:${MY_FUNCTION_NAME}:latest"

   # execute curl
   curl -X POST \
    -H "Content-Type: application/json" \
    -H "x-auth-token: ${OTC_X_AUTH_TOKEN}" \
    -d '{"key":"Hello World of FunctionGraph"}' \
    https://functiongraph.${OTC_SDK_REGION}.otc.t-systems.com/v2/${OTC_SDK_PROJECTID}/fgs/functions/${MY_FUNCTION_URN}/invocations


.. code-block:: text

    # Output will be like:
    {"statusCode": 200, "isBase64Encoded": false, "body": "{\"key\": \"Hello World of FunctionGraph\"}", "headers": {"Content-Type": "application/json"}}


Calling Functiongraph using Username and Password asynchronously
----------------------------------------------------------------

See :docs_otc:`Executing a Function Asynchronously <function-graph/api-ref/api/function_invocation/executing_a_function_asynchronously.html>`

.. code-block:: bash

   export MY_FUNCTION_NAME="DefaultPython3_10_From_Go_SDK"
   export MY_FUNCTION_URN="urn:fss:${OTC_SDK_REGION}:${OTC_SDK_PROJECTID}:function:default:${MY_FUNCTION_NAME}:latest"

   # execute curl
   curl -X POST \
    -H "Content-Type: application/json" \
    -H "x-auth-token: ${OTC_X_AUTH_TOKEN}" \
    -d '{"key":"Hello World of FunctionGraph"}' \
    https://functiongraph.${OTC_SDK_REGION}.otc.t-systems.com/v2/${OTC_SDK_PROJECTID}/fgs/functions/${MY_FUNCTION_URN}/invocations-async

.. code-block:: text

    # Output will be like:
    {"request_id": "c1db74a0-35a7-4717-b1db-1f66380d68d8"}
