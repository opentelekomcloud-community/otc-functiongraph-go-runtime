.. _deploying_an_http_function_using_a_container_image_built_with_go:

Deploying an HTTP Function container Image using Terraform
===============================================================

This section describes on how to deploy an HTTP Function packed in a container image using Terraform.

Prerequisite
------------

* Terraform configured according to :ref:`ref_terraform_setup`
* Image created as described in :ref:`creating_an_http_function_using_a_container_image_built_with_go`
* Exiting API Gateway instance to create the trigger for the HTTP Function.
  see section :ref:`variables.tfvars <deploying_an_http_function_using_a_container_image_built_with_go_variables_tfvars>` below.

Full sample can be found in :github_repo_master:`samples-doc/container-http <samples-doc/container-http>`.

Terraform Scripts
---------------------
Terraform deployment scripts can be found in:
:github_repo_master:`samples-doc/container-http/terraform <samples-doc/container-http/terraform>`

provider.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script configures the OpenTelekomCloud provider for Terraform.

.. literalinclude:: /../../samples-doc/container-http/terraform/provider.tf
   :language: terraform
   :caption: provider.tf

variables.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script defines the variables used in the Terraform scripts.

.. literalinclude:: /../../samples-doc/container-http/terraform/variables.tf
   :language: terraform
   :caption: variables.tf

.. _deploying_an_http_function_using_a_container_image_built_with_go_variables_tfvars:

variables.tfvars
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

This file provides values for the variables defined in **variables.tf**.

.. literalinclude:: /../../samples-doc/container-http/terraform/variables.tfvars
   :language: terraform
   :caption: variables.tfvars

agency.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates an agency with required permissions
to allow FunctionGraph to pull images from SWR.

.. literalinclude:: /../../samples-doc/container-http/terraform/agency.tf
   :language: terraform
   :caption: agency.tf

loggroup.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a log group and log stream in LTS.

.. literalinclude:: /../../samples-doc/container-http/terraform/loggroup.tf
   :language: terraform
   :caption: loggroup.tf

function.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates the HTTP Function using the container image.

.. literalinclude:: /../../samples-doc/container-http/terraform/function.tf
   :language: terraform
   :caption: function.tf

In the **function.tf** script, the relevant part to create
the HTTP Function using a custom container image is as follows:

.. code-block:: terraform

    handler   = "-"
    code_type = "Custom-Image-Swr"
    runtime   = "http"

    custom_image {
      url = var.image_url
    }

were `var.image_url` is the address of the container image in SWR.
in the format:

.. code-block:: text

    swr.<REGION>.otc.t-systems.com/<SWR-ORGANIZATION>/<IMAGE_NAME>:<TAG>


api_trigger.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates an API Gateway trigger to invoke the HTTP Function.

.. literalinclude:: /../../samples-doc/container-http/terraform/api_trigger.tf
   :language: terraform
   :caption: api_trigger.tf

testevent.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a test event to test the HTTP Function.

.. literalinclude:: /../../samples-doc/container-http/terraform/testevent.tf
   :language: terraform
   :caption: testevent.tf


Deployment
-------------------------------------

MakefileTF
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

To simplify the development and testing process,
see **MakefileTF** in the **container-http** folder:

.. literalinclude:: /../../samples-doc/container-http/MakefileTF
   :language: make
   :caption: MakefileTF
   :tab-width: 2

This **MakefileTF** imports all targets from the build **Makefile** and
provides additional targets for Terraform deployment.

Initialize Terraform
"""""""""""""""""""""""""""""""""""""

To initialize Terraform, run following command in
the **container-http** folder:

.. code-block:: bash

    make -f MakefileTF initTerraform

Plan Terraform deployment
"""""""""""""""""""""""""""""""""""""

To plan the Terraform deployment, run following command in
the **container-http** folder:

.. code-block:: bash

    make -f MakefileTF plan

Deploy using Terraform
"""""""""""""""""""""""""""""""""""""

To deploy the HTTP Function using Terraform, run following command in
the **container-http** folder:

.. code-block:: bash

    make -f MakefileTF deploy

Test deployed Function
"""""""""""""""""""""""""""""""""""""
To test the deployed HTTP Function using the API Gateway trigger, run following command in
the **container-http** folder:

.. code-block:: bash

    make -f MakefileTF test_deployed


Cleanup deployed resources
"""""""""""""""""""""""""""""""""""""

.. note::

   To destroy the deployed resources, run following command in
   the **container-http** folder:

   .. code-block:: bash

        make -f MakefileTF tf_destroy
