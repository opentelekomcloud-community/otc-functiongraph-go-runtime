S3 Event Thumbnail Sample
=========================

.. toctree::
   :hidden:


This is a sample that processes an image uploaded to OBS,
resizes it to fit within a maximum dimension, and uploads the resized
image back to another OBS using FunctionGraph with OBS trigger event.



Overview
--------

Following diagram shows components used in this example:

.. image:: ./thumbnail.drawio.svg
  :width: 800
  :alt: Components

Source
------

Source for this sample can be found in:
:github_repo_master:`/samples-doc/event-obss3-thumbnail</samples-doc/event-obss3-thumbnail>`.

Deployment
----------
This sample can be deployed using Terraform,
see :ref:`ref_terraform_setup` for setup details.

Terraform deployment scripts can be found in:
:github_repo_master:`/samples-doc/event-obss3-thumbnail/terraform <samples-doc/event-obss3-thumbnail/terraform>`

../Makefile
~~~~~~~~~~~

Adapt file *Makefile* variables according to your needs.

.. literalinclude:: /../../samples-doc/event-obss3-thumbnail/Makefile
   :language: make
   :caption: /Makefile

main.tf
~~~~~~~

This files contains all resources to be created for this sample.

.. literalinclude:: /../../samples-doc/event-obss3-thumbnail/terraform/main.tf
   :language: terraform
   :caption: /main.tf

provider.tf
~~~~~~~~~~~

This file contains the terraform provider configuration.

.. note::

   Check especially the **backend "s3"** configuration for **bucket** and **key**.

.. literalinclude:: /../../samples-doc/event-obss3-thumbnail/terraform/provider.tf
   :language: terraform
   :caption: /provider.tf

To deploy use following command in directory where *Makefile* is located:

.. code-block:: bash

  make deploy


.. note::

  This terraform deployment creates an agency with
  permissions for FunctionGraph to access OBS.

  The creation of an agency with permissions is a time consuming task
  and may take up to several minutes to complete.

  Until the agency is fully, testing the function may fail with
  permission errors like:

  **`Error fetching object otc.jpg from bucket
  csharp-doc-sample-event-obss3-thumbnail-csharp-images:
  One or more errors occurred. (Access Denied)`**

References
----------

In this sample the `AWS SDK for .Net <https://github.com/aws/aws-sdk-net>`_
is used to access OBS service:

.. code-block:: xml
  :caption: :github_repo_master:`event_obss3_thumbnail.csproj <samples-doc/event-obss3-thumbnail/src/event_obss3_thumbnail.csproj>`

    <ItemGroup>
      <PackageReference Include="AWSSDK.S3" Version="3.7.4.3" />
      <PackageReference Include="AWSSDK.Core" Version="3.7.4.3" />
    </ItemGroup>

Documentation on the `AWS SDK for .Net` can be found here:
`Amazon S3 examples using SDK for .NET <https://docs.aws.amazon.com/sdk-for-net/v3/developer-guide/csharp_s3_code_examples.html>`_

Upload image file to source bucket
----------------------------------

Linux/Ubuntu
~~~~~~~~~~~~~~~~~~~~~~

.. note::

   For *s3cmd* tool installation see: `S3cmd <https://github.com/opentelekomcloud/obs-s3/tree/master/s3cmd>`_
   on github.


To upload image to source bucket, following script an be used for Ubuntu users:

.. literalinclude:: /../../samples-doc/event-obss3-thumbnail/testUpload.sh
   :language: bash
   :caption: /samples-doc/event-obss3-thumbnail/testUpload.sh


Microsoft Windows
~~~~~~~~~~~~~~~~~~~~~~

.. note::
   For *Microsoft Windows*, see `OBS Browser+ <https://docs.otc.t-systems.com/object-storage-service/tool-guide/index.html#>`_



Alternatives (unsupported)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Huawei obsutil
""""""""""""""

.. note::

  Huawei *obsutil* is available for Windows, Linux and Mac from Huawei,
  see: `obsutil Introduction <https://support.huaweicloud.com/intl/en-us/utiltg-obs/obs_11_0001.html>`_

.. literalinclude:: /../../samples-doc/event-obss3-thumbnail/testUpload.cmd
   :language: bash
   :caption: /samples-doc/event-obss3-thumbnail/testUpload.cmd

