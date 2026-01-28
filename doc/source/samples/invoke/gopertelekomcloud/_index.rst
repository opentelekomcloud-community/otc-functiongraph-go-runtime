Calling FunctionGraph using gopertelekomcloud
==============================================

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg-openstack/src/invokeSync_UsernamePassword.go <samples-doc/invoke-fg-openstack/src/invokeSync_UsernamePassword.go>`.


For details on gopertelekomcloud,
see `gopertelekomcloud <https://github.com/opentelekomcloud/gophertelekomcloud>`_  on Github.

Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^
See environment variables section in: :ref:`ref_invoke-prerequisites`

Deployed FunctionGraph Event Function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Method A: Deploy using console
""""""""""""""""""""""""""""""""""""

Deploy the function manually as described in :ref:`ref_invoke-prerequisites`
using the OpenTelekomCloud Console.

Method B: Deploy using gopertelekomcloud
""""""""""""""""""""""""""""""""""""""""

Or deploy the function using the provided Go code using gopertelekomcloud.

.. code-block:: bash

    cd samples-doc/invoke-fg-openstack/src
    go test -run TestCreateFunction_UsernamePassword


this will create function with code and parameters from:

.. literalinclude:: /../../samples-doc/invoke-fg-openstack/src/functionConfig.go
   :language: go
   :caption: samples-doc/invoke-fg-openstack/src/functionConfig.go


After testing, you can delete the function using:

.. code-block:: bash

    cd samples-doc/invoke-fg-openstack/src
    go test -run TestDeleteFunction_UsernamePassword



Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg-openstack <samples-doc/invoke-fg-openstack>`.


Call Functiongraph using Username and Password synchronously
-------------------------------------------------------------




To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke-fg-openstack/src
    go test -run TestInvokeSync_UsernamePassword

Call Functiongraph using AK/SK synchronously
-----------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg-openstack/src/invokeSync_AKSK.go <samples-doc/invoke-fg-openstack/src/invokeSync_AKSK.go>`.


To run the sample, execute:

.. code-block:: bash

    cd samples-doc/invoke-fg-openstack/src
    go test -run TestInvokeSync_AKSK


.. note::

   Don't forget to clean up the deployed resources after testing.
