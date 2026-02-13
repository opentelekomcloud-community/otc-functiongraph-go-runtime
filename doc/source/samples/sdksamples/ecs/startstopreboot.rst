ECS Sample to start/stop/reboot an instance
============================================

.. toctree::
   :hidden:

This sample demonstrates how to start/stop/reboot an ECS instance
using FunctionGraph and:

* :otc_docs:`OpenTelekomCloud Rest API for ECS batch operations <elastic-cloud-server/api-ref/apis_recommended/batch_operations/starting_ecss_in_a_batch.html>` 
* :github_go_sign_sdk:`otc-api-sign-sdk-go <>`

Prerequisites
^^^^^^^^^^^^^^^^^^^^^
* For this example an ECS instance must exist.
* The function must have permissions to start/stop/reboot the ECS instance.

  This can be achieved by creating an agency with a policy
  granting the permission `ecs:*:start`, `ecs:*:stop`, and `ecs:*:reboot` and
  specifying this agency when creating the function.
  (E.g. create an agency with `ECS User` System-defined policy).


Source
-------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/sdk-ecs</samples-doc/sdk-ecs>`.

.. tabs::

  .. tab:: main.go

     .. literalinclude:: /../../samples-doc/sdk-ecs/src/main.go
        :language: go
        :caption: /main.go

  .. tab:: go.mod

    This files contains the main program.

     .. literalinclude:: /../../samples-doc/sdk-ecs/go.mod
        :language: go
        :caption: /go.mod

    


Build the project
-----------------

To build the project, navigate to the project directory and run the following
command:


.. tabs::
  .. tab:: Linux

    For Windows Subsystem for Linux (WSL) or Linux/MacOS systems use following makefile:

     .. literalinclude:: /../../samples-doc/sdk-ecs/Makefile
        :language: makefile
        :caption: /Makefile

     Then run the command:

     .. code-block:: bash

        make all

     This command builds the project in **target** folder and creates a zip file
     in the project folder.

     The generated zip file is: **deploy.zip**


  .. tab:: Windows

     For Windows use following script:

     .. literalinclude:: /../../samples-doc/sdk-ecs/build.cmd
        :language: bat
        :caption: /build.cmd

     Then run the command:

     .. code-block:: bash

        build.cmd

     This command builds the project in **target_win** folder and creates a zip file
     in the project folder.

     The generated zip file is: **deploy.zip**



Deploy the function
-------------------

Use :fg_console:`FunctionGraph console <>` to create a function with following settings:

Create function
*******************

**Create With**:  Create from scratch

**Basic Information**

* **Function Type**  Event Function
* **Region**  <YOUR REGION>
* **Function Name** <YOUR FUNCTION NAME>
* **Agency**  Specify an agency with policy to start ECS instance
* **Runtime**  Go 1.x

Upload code
*******************

Use **Upload** > **Local ZIP** and upload *deploy.zip*
from previous step.

Configure function
*******************

* In **Configuration** > **Basic Settings** > **Handler**:
  set value to name as defined in **handler.txt**

* In **Configuration** > **Environment Variables** add following variables:

    .. list-table:: Environment variables
      :widths: 20 20 25
      :header-rows: 1

      * - Environment variable name
        - Value
        - Remarks

      * - ECS_INSTANCE_ID
        - <ID of ecs instance>
        - ID of ECS instance to start

      * - ECS_ENDPOINT_URL
        - <ecs endpoint>
        - Default: https://ecs.eu-de.otc.t-systems.com
          see :otc_docs:`Regions and Endpoints<regions-and-endpoints/index.html>`

      * - ECS_ACTION
        - <action>
        - Action to perform on the ECS instance ("start", "stop", "reboot"),
          default: "start"

      * - ECS_ACTION_TYPE
        - <action type>
        - Action type to perform on the ECS instance for reboot/stop ("SOFT", "HARD"),
          default: "SOFT"

Test the function
-------------------

Create Test Event
*******************

In **Code** create a Test Event using "Blank Template"
(Event is not used in function).

Test function
*******************

Click **Test** to test function.

The function execution result is displayed in the
**Execution Result** section.
