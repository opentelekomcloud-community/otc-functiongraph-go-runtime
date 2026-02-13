HTTP Functions
==========================

.. toctree::
   :hidden:

   Create from scratch <scratch/_index>
   Container Image <container/_index>

HTTP functions support mainstream Web application frameworks
and can be accessed through a browser or called directly by a URL.

Types of Building Go HTTP Functions
------------------------------------

You can build FunctionGraph HTTP functions in Go in the following ways:

* :ref:`From Scratch  <devguide_http_function_scratch_index>`
* :ref:`Using Container Image  <devguide_http_function_container_index>`

.. note::
  The option "Select template" is not supported for Go HTTP functions.

.. _general_constraints_go_http:

General Constraints for HTTP Functions
-----------------------------------------

Following are the general constraints of HTTP functions:

- HTTP functions can be used with APIG triggers.
  According to the forwarding protocol between FunctionGraph and APIG,
  a valid HTTP function response must contain:

  - **body(String)**,
  - **statusCode(int)**,
  - **headers(Map)**, and
  - **isBase64Encoded(boolean)**.

  By default, the response is encoded using Base64.
  The default value of **isBase64Encoded** is **true**.

  For details about the constraints, see :ref:`Base64 Decoding and
  Response Structure<ref_apig_event_base64>`.

- The bound IP address is **127.0.0.1.**

- By default, port **8000** is enabled for HTTP functions, no other port can be used.

- By default, an account can create a maximum of 400 functions.
  (This quota can be increased upon request.)

- HTTP functions cannot be executed for a long time,
  invoked asynchronously, or retried.

- When a function initiates an HTTP request, the request IP address
  is dynamic for private network access and fixed for public network access.

Common Request Headers of HTTP Functions
-----------------------------------------

HTTP request headers are an important part of the HTTP protocol for
passing metadata.
When a function is invoked, specific metadata or configuration information
can be passed. Following Table describes the common request headers carried
by functions by default.

.. list-table:: Common Request Headers of HTTP Functions
   :header-rows: 1

   * - Header Name
     - Description
   * - X-CFF-Request-Id
     - ID of the current request.
   * - X-CFF-Memory
     - Memory allocated to the function.
   * - X-CFF-Timeout
     - Function timeout.
   * - X-CFF-Func-Version
     - Function version.
   * - X-CFF-Func-Name
     - Function name.
   * - X-CFF-Project-Id
     - Project ID of the function.
   * - X-CFF-Package
     - App to which the function belongs.
   * - X-CFF-Region
     - Region where the function is located.


.. note::

  The key information of HTTP functions:

     - X-CFF-Auth-Token
     - X-CFF-Security-Access-Key
     - X-CFF-Security-Secret-Key
     - X-CFF-Security-Token

   can be transferred only through request headers.
   For details about how to obtain the AK, SK, and token of HTTP functions,
   see :ref:`Transferring Secret Keys Through the Request Header <transferringKeys-ref>`

