Define FunctionGraph function handler in Go
==========================================================

The Go function handler is the method in your function code that
processes events.
When your function is invoked, FunctionGraph runs the handler method.
Your function runs until the handler returns a response, exits, or times out.


Setting up the Go handler project
---------------------------------

A typical Go FunctionGraph project is typically structured as follows:

.. code-block:: console
  :caption: Project structure

  /project-root
   ├─ src
   |  └─ main.go
   ├─ go.mod
   ├─ go.sum
   └─ Makefile

The main logic for the function resides in Go file **main.go**.
When deploying to FunctionGraph make sure to specify the correct handler:

The name of the handler is the **name of the Go executable** (here: main) file.


Example code for Go FunctionGraph function
------------------------------------------

A FunctionGraph function written in `Go <https://golang.org/>`_ is authored as
a Go executable.
You can initialize a Go FunctionGraph function project the same way you would
initialize any other Go project using the following *go mod init* command:

.. code-block:: bash

   go mod init example


Here, **example** is the module name. You can replace this with anything.

This command initializes your project in the current folder and generates
the **go.mod** file that lists your project's dependencies.

Use the **go get** command to add any external dependencies to your project.

For example, for all FunctionGraph functions in Go, you must include the
**github.com/opentelekomcloud-community/otc-functiongraph-go-runtime** package,
which implements the FunctionGraph programming model for Go.

Include this package with the following go get command:

.. code-block:: bash

    # enable Go modules to get specific version of the runtime package
    go env -w GO111MODULE=on

    # install FunctionGraph Go runtime package for latest version
    go get github.com/opentelekomcloud-community/otc-functiongraph-go-runtime

    # or for a specified version
    go get github.com/opentelekomcloud-community/otc-functiongraph-go-runtime@v0.0.1

Your function code should live in a Go file. In the following example,
we name this file **main.go**. In this file, you implement your core function
logic in a **handleRequest** method, as well as a **main()** function that
starts the runtime with this handler.

Example Go FunctionGraph function code
--------------------------------------

The following example shows a simple FunctionGraph function written in Go.

.. literalinclude:: /../../samples-doc/example/src/main.go
    :language: go
    :caption: :github_repo_master:`main.go <samples-doc/example/src/main.go>`

Handler name is: **main**


Syntax for creating a handler function in Go:
---------------------------------------------

``func`` **Handler** ( **payload** \[\] ``byte`` , **ctx**  ``context.RuntimeContext)``

* Entry function name (**Handler**):

  Entry function name.

* Execution event body (**payload**)

  The execution event parameters entered by the user in the function
  execution interface, in the format of a JSON object.

* Context (**ctx**):

  The function execution context provided by Runtime.


Return values of the handler function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

``return`` **object**, **error**

* If the **error** parameter returned by a function is **not nil**,
  the function execution fails.
* If the **error** parameter returned by a function is **nil**,
  FunctionGraph supports only the following types of values:

  .. list-table:: Supported return value types
     :header-rows: 1
     :widths: 25 75

     * - **object** Type
       - Description
     * - nil
       - The HTTP response body is empty.
     * - \[\]byte
       - The content in this byte array is the body of an HTTP response.
     * - string
       - The content in this string is the body of an HTTP response.
     * - Other
       - FunctionGraph returns a value for JSON encoding, and uses
         the encoded object as the body of an HTTP response.

         The **Content-Type** header of the HTTP response is set to **application/json**.

.. tabs::

   .. tab:: Linux

      .. literalinclude:: /../../samples-doc/example/Makefile
          :language: Makefile
          :caption: :github_repo_master:`Makefile <samples-doc/example/Makefile>`

      The Makefile automates the build process for your Go FunctionGraph function.

      You can run the **make build** command in the project root directory
      to compile your function code and generate the executable file named
      **example**.

      The **make zip** command creates a deployment package named **deploy.zip**
      that contains the executable file.

   .. tab:: Windows

      .. literalinclude:: /../../samples-doc/example/build.cmd
         :language: bat
         :caption: :github_repo_master:`build.cmd <samples-doc/example/build.cmd>`

      Run the **build.cmd** script in the project root directory to compile your function code
      and generate the executable file named **example** in the **target_win** directory.

      The script also creates a deployment package named **deploy.zip** in the project
      root directory using the Windows tar.exe command.

You can then upload this deployment package to FunctionGraph.

For general deployment instructions see
:otc_fg_umn:`Building Functions<building_functions/index.html>` in User Guide.


Accessing and using the FunctionGraph context object
----------------------------------------------------

The :doc:`Context<./context>` interface allows functions to obtain the
function execution context, such as information about the invocation,
function, execution environment, and so on.

The context is of type ``ctx context.RuntimeContext``
and is the second argument of the handler function.

* :github_repo_master:`ctx context.RuntimeContext <go-runtime/go-api/context/client-context.go>`

To produce logs in OpenTelekomCloud Log Tank Servics (LTS) you can use
``ctx.GetLogger()`` to get a RuntimeLogger object for logging.

.. code-block:: go

  ctx.GetLogger().Logf("Received payload: %s", string(payload))

Besides of logging, you can also use the context object for
function monitoring.
For more information about the context object,
see :doc:`Using the FunctionGraph context object to retrieve function information.<./context>`

Accessing environment variables
-------------------------------

Environment variables defined in ``OpenTelekomCloud`` ->
``Configuration`` -> ``Environment Variables`` can be accessed using:

.. code-block:: java

  // accessing an environment variable named "ENV_VAR1"
  ctx.getUserData("ENV_VAR1");
