Calling FunctionGraph using gopertelekomcloud
==============================================

For details on gopertelekomcloud,
see `gopertelekomcloud <https://github.com/opentelekomcloud/gophertelekomcloud>`_  on Github.

Prerequisites
-------------

See :ref:`ref_invoke-prerequisites`

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg-openstack <samples-doc/invoke-fg-openstack>`.


Call Functiongraph using Username and Password synchronously
-------------------------------------------------------------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/invoke-fg-openstack/src/invokeSync_UsernamePassword.go <samples-doc/invoke-fg-openstack/src/invokeSync_UsernamePassword.go>`.


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
