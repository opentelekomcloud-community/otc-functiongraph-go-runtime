Calling FunctionGraph using bash scripts
==================================================

Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^
See environment variables section in: :ref:`ref_invoke-prerequisites`

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

.. code-block:: bash

    # execute the script to get the token without output to stdout
    MY_TOKEN=$(./tokenFromUsername.sh) > /dev/null

    # Output will be like:
    echo $MY_TOKEN

    # Output example:
    MIIGBQYJKoZIhvcNAQcCoIIF9jCCBfICAQExDTALBglghkgBZQMEAgEwggOKBgkqhkiG.....

Calling Functiongraph using Username and Password synchronously
----------------------------------------------------------------

See :docs_otc:`Executing a Function Synchronously <function-graph/api-ref/api/function_invocation/executing_a_function_synchronously.html#functiongraph-06-0125>`

.. code-block:: bash

   export MY_FUNCTION_URN="urn:fss:${OTC_SDK_REGION}:${OTC_SDK_PROJECTID}:function:[FUNCTION_NAME]:latest"

   export MY_FUNCTION_URN="urn:fss:eu-de:d52e41d2434941b194ce3f91b1b12f8a:function:default:go_go-doc-sample-container-event:latest"

   # execute curl
   curl -X POST \
    -H "Content-Type: application/json" \
    -H "x-auth-token: ${MY_TOKEN}" \
    -d '{"key":"Hello World of FunctionGraph"}' \
    https://functiongraph.${OTC_SDK_REGION}.otc.t-systems.com/v2/${OTC_SDK_PROJECTID}/fgs/functions/${MY_FUNCTION_URN}/invocations


Calling Functiongraph using Username and Password asynchronously
----------------------------------------------------------------

See :docs_otc:`Executing a Function Asynchronously <function-graph/api-ref/api/function_invocation/executing_a_function_asynchronously.html>`

.. code-block:: bash

   export MY_FUNCTION_URN="urn:fss:${OTC_SDK_REGION}:${OTC_SDK_PROJECTID}:function:[FUNCTION_NAME]:latest"

   export MY_FUNCTION_URN="urn:fss:eu-de:d52e41d2434941b194ce3f91b1b12f8a:function:default:go_go-doc-sample-container-event:latest"

   # execute curl
   curl -X POST \
    -H "Content-Type: application/json" \
    -H "x-auth-token: ${MY_TOKEN}" \
    -d '{"key":"Hello World of FunctionGraph"}' \
    https://functiongraph.${OTC_SDK_REGION}.otc.t-systems.com/v2/${OTC_SDK_PROJECTID}/fgs/functions/${MY_FUNCTION_URN}/invocations-async

.. code-block:: text

    # Output will be like:
    {"request_id": "c1db74a0-35a7-4717-b1db-1f66380d68d8"}
