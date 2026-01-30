Invoking FunctionGraph Event Function from Go
==============================================

Following pages demonstrate how to call a FunctionGraph implemented in
python using Go:

.. toctree::
    :maxdepth: 1

    Using API calls<apicalls/_index>
    Using gopertelekomcloud<gopertelekomcloud/_index>
    Using bash <bash/_index>


.. _ref_invoke-prerequisites:

Prerequisites
-------------

Environment Variables
^^^^^^^^^^^^^^^^^^^^^^^^^

Needed environment variables set

.. list-table:: Environment Variables
   :widths: 25 25
   :header-rows: 1

   * - Name
     - Description
   * - OTC_SDK_PROJECTID
     - Project ID
   * - OTC_SDK_REGION
     - Region, e.g. "eu-de"
   * - OTC_SDK_AK
     - Access Key
   * - OTC_SDK_SK
     - Secret Key
   * - OTC_USER_NAME
     - User name
   * - OTC_USER_PASSWORD
     - User password
   * - OTC_DOMAIN_NAME
     - Domain name
   * - OTC_IAM_ENDPOINT
     - IAM Endpoint, e.g. https://iam.eu-de.otc.t-systems.com/v3


Deployed FunctionGraph Event Function
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Example FunctionGraph written in python


.. literalinclude:: /../../samples-doc/invoke-fg/functionGraph.py
   :language: python
   :caption: Example FunctionGraph: functionGraph.py

Deploy this FunctionGraph as follows:

- Type: "Event Function":
- Name: "DefaultPython3_10_From_Go_SDK"
- Agency: "Use no agency"
- Runtime: "python 3.10"
- Application: "default"
- Handler: "index.handler"

