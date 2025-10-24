APIG Event Source
=================

API Gateway (APIG) is an API hosting service that helps enterprises to build,
manage, and deploy APIs at any scale. With APIG, your function can be invoked
through HTTPS by using a custom REST API and a specified backend. You can map
each API operation (such as, GET and PUT) to a specific function. APIG invokes
the relevant function when an HTTPS request is sent to the API backend.

For more information about how to use HTTPS calls to trigger functions, see
`Using an APIG (Dedicated) trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_apig_dedicated_trigger.html#>`__.

Example APIG Event
------------------

.. code-block:: json

   {
     "body": "{\"test\":\"body\"}",
     "requestContext": {
       "apiId": "bc1dcffd-aa35-474d-897c-d53425a4c08e",
       "requestId": "11cdcdcf33949dc6d722640a13091c77",
       "stage": "RELEASE"
     },
     "queryStringParameters": {
       "responseType": "html"
     },
     "httpMethod": "GET",
     "pathParameters": {
       "path":"value"
     },
     "headers": {
       "accept-language": "en-US;q=0.3,en;q=0.2",
       "accept-encoding": "gzip, deflate, br",
       "x-forwarded-port": "443",
       "x-forwarded-for": "103.218.216.98",
       "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
       "upgrade-insecure-requests": "1",
       "host": "50eedf92-c9ad-4ac0-827e-d7c11415d4f1.apigw.region.cloud.com",
       "x-forwarded-proto": "https",
       "pragma": "no-cache",
       "cache-control": "no-cache",
       "x-real-ip": "103.218.216.98",
       "user-agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0"
     },
     "path": "/apig-event-template",
     "isBase64Encoded": true
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
   * - body
     - String
     - "{\"test\":\"body\"}"
     - Actual request in string format.
   * - requestContext
     - Map
     - See the example.
     - Request information, including the API gateway configuration, request ID, authentication information, and source.
   * - httpMethod
     - String
     - GET
     - HTTP method
   * - queryStringParameters
     - Map
     - See the example.
     - Query strings configured in APIG and their actual values
   * - pathParameters
     - Map
     - See the example.
     - Path parameters configured in APIG and their actual values
   * - headers
     - Map
     - See the example.
     - Complete headers
   * - path
     - String
     - /apig-event-template
     - Complete path
   * - isBase64Encoded
     - Boolean
     - True
     - Default value: true (see Notes below)

Notes
-----

- When calling a function using APIG, **isBase64Encoded** is valued true by
  default, indicating that the request body transferred to FunctionGraph is
  encoded using Base64 and must be decoded for processing.

- The function must return characters strings by using the following structure.

  .. code-block:: json

     {
       "isBase64Encoded": "true|false",
       "statusCode": "httpStatusCode",
       "headers": {"headerName":"headerValue"},
       "body": "..."
     }

Example
-------

.. literalinclude:: /../../samples-events/apig/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-events/apig/handler.go>`
