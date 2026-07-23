.. _deploying_an_http_function_using_terraform:

Deploying a HTTP Function using Terraform
=========================================

This section describes on how to deploy an HTTP Function using Terraform.

Prerequisite
------------

* Terraform configured according to :ref:`ref_terraform_setup`
* Exiting API Gateway instance to create the trigger for the HTTP Function.


Full sample can be found on Github: :github_repo_master:`Building HTTP Function <samples-doc/scratch-http>`.

Terraform Scripts
---------------------
Terraform deployment scripts can be found in:
:github_repo_master:`samples-doc/scratch-http/terraform <samples-doc/scratch-http/terraform>`

provider.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script configures the provider for Terraform.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/provider.tf
   :language: terraform
   :caption: provider.tf

variables.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script defines the variables used in the Terraform scripts.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/variables.tf
   :language: terraform
   :caption: variables.tf

variables.tfvars
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

This file provides values for the variables defined in **variables.tf**.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/variables.tfvars
   :language: terraform
   :caption: variables.tfvars

loggroup.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a log group and log stream in LTS.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/loggroup.tf
   :language: terraform
   :caption: loggroup.tf

function.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates the HTTP Function using the container image.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/function.tf
   :language: terraform
   :caption: function.tf

In the **function.tf** script, the relevant part to create
the HTTP Function using code from zip file is as follows:

.. code-block:: terraform

    handler   = "bootstrap"
    code_type = "zip"
    runtime   = "http"
    func_code     = filebase64(format("${path.module}/../%s", var.zip_file_name))
    code_filename = var.zip_file_name

api_trigger.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates an API Gateway trigger and all relevant resources to
invoke the HTTP Function:

- API gateway group (opentelekomcloud_apigw_group_v2)
- API resource (opentelekomcloud_apigw_api_v2) that connects
  this API to the FunctionGraph function
- publish API using (opentelekomcloud_apigw_api_publishment_v2)

.. literalinclude:: /../../samples-doc/scratch-http/terraform/api_trigger.tf
   :language: terraform
   :caption: api_trigger.tf


.. note::

    To have full control over resources being created this sample does not use
    **opentelekomcloud_fgs_trigger_v2** for APIG Trigger creation,

testevent.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a test event to test the HTTP Function.

.. literalinclude:: /../../samples-doc/scratch-http/terraform/testevent.tf
   :language: terraform
   :caption: testevent.tf

Deployment
-------------------------------------

MakefileTF
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

To simplify the development and testing process,
see **MakefileTF** in the project root folder:

.. literalinclude:: /../../samples-doc/scratch-http/MakefileTF
   :language: make
   :caption: MakefileTF
   :tab-width: 2

This **MakefileTF** imports all targets from the build **Makefile** and
provides additional targets for Terraform deployment.

Initialize Terraform
"""""""""""""""""""""""""""""""""""""

To initialize Terraform, run following command in
the project root folder:

.. code-block:: bash

    make -f MakefileTF tf_init

Plan Terraform deployment
"""""""""""""""""""""""""""""""""""""

To plan the Terraform deployment, run following command in
the project root folder:

.. code-block:: bash

    make -f MakefileTF tf_plan

Deploy using Terraform
"""""""""""""""""""""""""""""""""""""

To deploy the HTTP Function using Terraform, run following command in
the project root folder:

.. code-block:: bash

    make -f MakefileTF tf_apply

Test deployed Function
"""""""""""""""""""""""""""""""""""""
To test the deployed HTTP Function using the API Gateway trigger, run following command in
the project root folder:

.. code-block:: bash

    make -f MakefileTF test_deployed


Cleanup deployed resources
"""""""""""""""""""""""""""""""""""""

.. note::

   To destroy the deployed resources, run following command in
   the project root folder:

   .. code-block:: bash

        make -f MakefileTF tf_destroy
