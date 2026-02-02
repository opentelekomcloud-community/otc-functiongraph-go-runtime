Using the FunctionGraph context interface to retrieve Go function information
=============================================================================

When FunctionGraph runs your function, it passes a context object to the
handler.
This object provides methods and properties that provide information about
the invocation, function, and execution environment.

Context interface
-----------------

.. list-table:: **Table 1** Context interface methods
   :widths: 10 25
   :header-rows: 1

   * - Method Name
     - Description

   * - GetRequestID()
     - Get the request ID.

   * - GetRemainingTimeInMilliSeconds()
     - Get the remaining running time of a function.

   * - GetAccessKey()
     - Get the AccessKey (valid for 24 hours) with an agency.

       To use this method, you need to configure **agency** for the function.

       .. note::

         FunctionGraph has stopped maintaining the **GetAccessKey()** API in the Runtime
         SDK. You cannot use this API to obtain a temporary AK.

   * - GetSecretKey()
     - Get the SecretKey (valid for 24 hours) with an agency.

       To use this method, you need to configure the **agency** for the function.

       .. note::

         FunctionGraph has stopped maintaining the **GetSecretKey()** API in the Runtime
         SDK. You cannot use this API to obtain a temporary SK.

   * - GetSecurityAccessKey()
     - Get the SecurityAccessKey (valid for 24 hours) with an agency.

       To use this method, you need to configure a **agency** for the function.

   * - GetSecuritySecretKey()
     - Get the SecuritySecretKey (valid for 24 hours) with an agency.

       To use this method, you need to configure the **agency** for the function.

   * - GetSecurityToken()
     - Get the SecuritySecretToken (valid for 24 hours) with an agency.

       To use this method, you need to configure the **agency** for the function.

   * - GetUserData(String key)
     - Uses keys to obtain the values passed by environment variables.

   * - GetFunctionName()
     - Get the function name.

   * - GetRunningTimeInSeconds()
     - Get the function timeout.

   * - GetVersion()
     - Get the version of the function.

   * - GetMemorySize()
     - Allocated memory.

   * - GetCPUNumber()
     - Get the CPU resources used by the function.

   * - GetPackage()
     - Gets a function group.

   * - GetToken()
     - Get the token (valid for 24 hours) with an agency.

       To use this method, you need to configure the **agency** for the function.

   * - GetLogger()
     - Get the logger method provided by the context (by default, it will output information such as time and request ID).

   * - GetAlias()
     - Get the alias of a function

   * - NA: getInvokeProperty()
     - Get the invoke property

   * - NA: getTraceID()
     - Get the trace id (not yet available)

   * - NA: getInvokeID()
     - Get the invoke id

   * - NA: getState()
     - Get the state

   * - GetProjectID()
     - Get the project id

   * - GetHandler()
     - Get the function handler name

   * - GetInitializerHandler()
     - Get the function initializer handler name

   * - GetInitializerTimeout()
     - Get the function initializer timeout


.. warning::
  Results returned by using the ``GetToken()``, ``GetAccessKey()``, and
  ``GetSecretKey()`` or ``GetSecurityToken()``, ``GetSecurityAccessKey()``, and
  ``GetSecuritySecretKey()`` methods contain sensitive information.
  Exercise caution when using these methods.

Logging
-------------

FunctionGraph provides a logging interface through the context object.

It records user input logs by using the method
**Logf(format string, args ...interface{})**.

This method outputs logs in the format of *Time in UTC* *Request ID* *Output*
for example, **2017-10-25T09:10:03.328Z 473d369d-101a-445e-a7a8-315cca788f86 test log output.**

The logger can be accessed by using the **GetLogger()** method of the
context object.

.. code-block:: go
   :caption: Example

    ctx.GetLogger().Logf("Hello world")

The preceding code outputs the following log:

**2017-10-25T09:10:03.328Z 473d369d-101a-445e-a7a8-315cca788f86 Hello world**