.. _transferringKeys-ref:

Transferring Secret Keys through the request header
----------------------------------------------------

The key information of HTTP functions can be transferred only
through request headers.

To obtain the AK, SK, and token of an HTTP function as shown in Table 1,
perform the following steps:

1. Log in to the :fg_console:`FunctionGraph console <>` console and go
   to the details page of the HTTP function to be configured
2. Choose **Configuration** > **Advanced Settings** and enable
   **Include Keys**.

    .. figure:: console_advanced_settings.png
      :scale: 50 %
      :alt: Enable transferring secret keys

      Enable transferring secret keys through the request header.


.. list-table:: Table 1: Key information transferred by HTTP functions
   :header-rows: 1

   * - Header Name
     - Description
   * - X-CFF-Auth-Token
     - A token is an access credential issued to an IAM user
       to bear its identity and permissions.
   * - X-CFF-Security-Access-Key
     - A temporary access key is issued by the system to IAM users.
   * - X-CFF-Security-Secret-Key
     - A temporary SecurityToken is issued by the system to IAM users.
   * - X-CFF-Security-Token
     - A temporary SecurityToken is issued by the system to IAM users.

.. note::

   - The temporary Security-Access-Key, Security-Secret-Key and SecurityToken
     must be used together.


On details how to use the Security-Access-Key, Security-Secret-Key,
and Security-Token to call other OTC services using API request,
see :github_go_sign_sdk:`Developer guide for request signing for Go<>`


.. hint::

  In Terraform set **enable_auth_in_header**  to **true** in the resource
  **opentelekomcloud_functiongraph_function_v2** to enable transferring
  keys through the request header.
