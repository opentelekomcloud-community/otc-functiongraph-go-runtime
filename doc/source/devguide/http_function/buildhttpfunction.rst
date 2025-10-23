Building an HTTP Function Using Go
===================================

.. toctree::
   :hidden:


Introduction
------------

This chapter describes how to deploy services on FunctionGraph using Go.

HTTP functions do not support direct code deployment using Go.
This section uses binary conversion as an example to describe how to
deploy Go programs on FunctionGraph.

This sample uses the `go-restful <https://github.com/emicklei/go-restful>`_
framework to implement the HTTP function.

Procedure
------------

Building a code package
^^^^^^^^^^^^^^^^^^^^^^^^^

Create the source file main.go. The code is as follows:

.. literalinclude:: /../../samples-doc/buildinghttpfunction/main.go
    :language: go
    :caption: :github_repo_master:`main.go <samples-doc/buildinghttpfunction/main.go>`


In **main.go**, an HTTP server is started using port **8000**,
and an API whose path is **/hello** is registered.

When the API is invoked, **"nice to meet you"** is returned.

Create the **go.mod** file to manage the dependencies of the project.

The content is as follows:

.. literalinclude:: /../../samples-doc/buildinghttpfunction/go.mod
    :language: go
    :caption: :github_repo_master:`go.mod <samples-doc/buildinghttpfunction/go.mod>`

Compiling and packaging
^^^^^^^^^^^^^^^^^^^^^^^^^

To ease building and packaging, create a Makefile in the project root directory
like the following:

.. literalinclude:: /../../samples-doc/buildinghttpfunction/Makefile
    :language: Makefile
    :caption: :github_repo_master:`Makefile <samples-doc/buildinghttpfunction/Makefile>`

The Makefile automates the build process for your Go FunctionGraph HTTP
function.

You can run the **make build** command in the project root directory
to compile your function code and generate the executable file named
**go-http-demo** in the **target** directory.

The **make zip** command creates a deployment package named **deploy.zip**
that contains the executable file and a **bootstrap** file.

The **bootstrap** file is required by FunctionGraph to identify
the entry file of the function.

.. code-block::

    /opt/function/code/go-http-demo

You can then upload this deployment package to FunctionGraph.


Testing using console
^^^^^^^^^^^^^^^^^^^^^

**Example Test Event**

In FunctionGraph console, you can create a test event to test the
HTTP function.

The following is an example test event:

.. literalinclude:: /../../samples-doc/buildinghttpfunction/test_event.json
   :language: json
   :caption: :github_repo_master:`test_event.json <samples-doc/buildinghttpfunction/test_event.json>`

The preceding example test event simulates an HTTP GET request
to the **/hello** API of the HTTP function.

When you use this test event to test the HTTP function,
the function returns the following response:

.. code-block:: json

    {
      "body": "bmljZSB0byBtZWV0IHlvdQ==",
      "headers": {
          "Content-Length": [
              "16"
          ],
          "Content-Type": [
              "text/plain; charset=utf-8"
          ],
          "Date": [
              "Thu, 23 Oct 2025 14:08:57 GMT"
          ]
      },
      "statusCode": 200,
      "isBase64Encoded": true
    }

The **body** field contains the Base64-encoded string of
**"nice to meet you"**.

The **statusCode** field indicates that the request is processed successfully.
