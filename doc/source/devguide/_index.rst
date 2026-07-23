.. _building_with_go:

Building with Go
========================
.. toctree::
   :hidden:


FunctionGraph Types
-------------------

FunctionGraph provides 2 types of functions:

* **Event Functions**

  Event functions can be configured with event triggers and integrate
  a variety of products
  (such as object storage service OBS, distributed messaging service
  RabbitMQ version, cloud log service LTS, etc.).

  See :doc:`Event Function <event_function/_index>`

* **HTTP Functions**

  HTTP functions support mainstream Web application frameworks and can
  be accessed through a browser or called directly by a URL.

  See :doc:`HTTP Functions <http_function/_index>`

Both types of functions can be built either from scratch or by using container images.

Supported Go Runtimes for building from scratch
-----------------------------------------------

FunctionGraph currently supports the following Go runtimes
for building functions from scratch:

.. list-table:: Supported Go runtimes
   :header-rows: 1

   * - Runtime
     - Description
     - Identifier
   * - Go 1.x
     - Supports Go 1.x
     - Go1.x

Supported Go Runtimes for building using container images
---------------------------------------------------------

For building functions using container images, you can use any
Go version that meets the requirements of your custom container image.


Set up development environment
---------------------------------
To build and run the Go runtime for FunctionGraph, you need to set up your development environment
by installing the Go programming language.


Operating system
^^^^^^^^^^^^^^^^^^^^

This guide assumes that you are using a Unix-like operating system such as

- Windows Subsystem for Linux (WSL)
  see `How to install Linux on Windows with WSL <https://learn.microsoft.com/en-us/windows/wsl/install>`_,
- Linux,
- macOS.

Install Go
^^^^^^^^^^^^^^^^^^^^
1. Download the Go installation package for your operating system from the official `Go website <https://golang.org/dl/>`_.
2. Follow the installation instructions provided on the website to install Go on your system.

Install an IDE
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
You can use any text editor or IDE to write Go code.

To maximize your Go development efficacy, follow the Go recommendations on
`Editor plugins and IDEs <https://golang.org/dl/>`_.

These provide the following features:

* Fully integrated debugging capabilities
* Syntax highlighting
* Code completion


.. note::
   Examples in this documentation were created using:

   - WSL and
   - Visual Studio Code.
