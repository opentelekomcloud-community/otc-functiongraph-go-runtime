.. _deploying_an_event_function_using_a_container_image_built_with_go:

Deploying an Event Function container Image using Terraform
===============================================================

This section describes on how to deploy an Event Function packed in a container image using Terraform.


Prerequisite
------------

* Terraform configured according to :ref:`ref_terraform_setup`
* Image created as described in :ref:`creating_an_event_function_using_a_container_image_built_with_go`

Full sample can be found in :github_repo_master:`samples-doc/container-event <samples-doc/container-event>`.


Terraform Scripts
---------------------
Terraform deployment scripts can be found in:
:github_repo_master:`samples-doc/container-event/terraform <samples-doc/container-event/terraform>`

provider.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script configures the OpenTelekomCloud provider for Terraform.

.. literalinclude:: /../../samples-doc/container-event/terraform/provider.tf
   :language: terraform
   :caption: provider.tf

variables.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script defines the variables used in the Terraform scripts.

.. literalinclude:: /../../samples-doc/container-event/terraform/variables.tf
   :language: terraform
   :caption: variables.tf

variables.tfvars
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This file provides values for the variables defined in **variables.tf**.

.. literalinclude:: /../../samples-doc/container-event/terraform/variables.tfvars
   :language: terraform
   :caption: variables.tfvars

agency.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates an agency with required permissions
to allow FunctionGraph to pull images from SWR.

.. literalinclude:: /../../samples-doc/container-event/terraform/agency.tf
   :language: terraform
   :caption: agency.tf

loggroup.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a log group and log stream in LTS.

.. literalinclude:: /../../samples-doc/container-event/terraform/loggroup.tf
   :language: terraform
   :caption: loggroup.tf

function.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates the Event Function using the container image.

.. literalinclude:: /../../samples-doc/container-event/terraform/function.tf
   :language: terraform
   :caption: function.tf

testevent.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a test event to test the Event Function.

.. literalinclude:: /../../samples-doc/container-event/terraform/testevent.tf
   :language: terraform
   :caption: testevent.tf


Deployment
-------------------------------------

MakefileTF
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

To simplify the development and testing process,
see **MakefileTF** in the **container-event** folder:

.. literalinclude:: /../../samples-doc/container-event/MakefileTF
   :language: make
   :caption: MakefileTF
   :tab-width: 2

This **MakefileTF** imports all targets from the build **Makefile** and
provides additional targets for Terraform deployment.

Initialize Terraform
"""""""""""""""""""""""""""""""""""""

To initialize Terraform, run following command in
the **container-event** folder:

.. code-block:: bash

    make -f MakefileTF tf_init

Plan Terraform deployment
"""""""""""""""""""""""""""""""""""""

To plan the Terraform deployment, run following command in
the **container-event** folder:

.. code-block:: bash

    make -f MakefileTF tf_plan

Deploy using Terraform
"""""""""""""""""""""""""""""""""""""

To deploy the Event Function using Terraform, run following command in
the **container-event** folder:

.. code-block:: bash

    make -f MakefileTF tf_apply

