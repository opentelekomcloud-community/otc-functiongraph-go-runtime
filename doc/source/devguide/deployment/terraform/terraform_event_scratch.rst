.. _deploying_an_event_function_using_terraform:

Deploying an Event Function using Terraform
===========================================

This section describes on how to deploy an Event Function using Terraform.

Prerequisite
------------

* Terraform configured according to :ref:`ref_terraform_setup`

Full sample can be found in :github_repo_master:`doc-sample-event-timer <samples-doc/doc-sample-event-timer>`.

Terraform Scripts
---------------------

Terraform deployment scripts can be found in:
:github_repo_master:`samples-doc/doc-sample-event-timer/terraform <samples-doc/doc-sample-event-timer/terraform>`

provider.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script configures the OpenTelekomCloud provider for Terraform.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/provider.tf
   :language: terraform
   :caption: provider.tf


variables.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script defines the variables used in the Terraform scripts.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/variables.tf
   :language: terraform
   :caption: variables.tf


loggroup.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a log group and log stream in LTS.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/loggroup.tf
   :language: terraform
   :caption: loggroup.tf

function.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates the Event Function using the container image.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/function.tf
   :language: terraform
   :caption: function.tf

timer_trigger.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates the timer trigger.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/timer_trigger.tf
   :language: terraform
   :caption: timer_trigger.tf

testevent.tf
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This script creates a test event to test the Event Function.

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/terraform/testevent.tf
   :language: terraform
   :caption: testevent.tf

Deployment
-------------------------------------

MakefileTF
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

To simplify the development and testing process,
see **Makefile** in the project root folder:

.. literalinclude:: /../../samples-doc/doc-sample-event-timer/Makefile
   :language: make
   :caption: Makefile
   :tab-width: 2

Initialize Terraform
"""""""""""""""""""""""""""""""""""""

To initialize Terraform, run following command in
the project root folder:

.. code-block:: bash

    make tf_init

Plan Terraform deployment
"""""""""""""""""""""""""""""""""""""

To plan the Terraform deployment, run following command in
the project root folder:

.. code-block:: bash

    make tf_plan

Deploy using Terraform
"""""""""""""""""""""""""""""""""""""

To deploy the Event Function using Terraform, run following command in
the project root folder:

.. code-block:: bash

    make tf_apply


Cleanup deployed resources
"""""""""""""""""""""""""""""""""""""

.. note::

   To destroy the deployed resources, run following command in
   the project root folder:

   .. code-block:: bash

      make tf_destroy
