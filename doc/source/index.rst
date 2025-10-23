OpenTelekomCloud FunctionGraph Go documentation
===============================================

This is the FunctionGraph Go Development Guide for running
FunctionGraph on OpenTelekomCloud.

.. toctree::
   :maxdepth: 10
   :hidden:

   Building with Go <devguide/_index>
   Event Functions <devguide/event_function/_index>
   HTTP Functions <devguide/http_function/_index>
   Samples <samples/_index>
   Best Practises <devguide/bestpractices/_index>


.. warning::

   Work in progress. Subject to change without notice.



Source Code
-----------

For source code, see :github_repo_master:`otc-functiongraph-go-runtime<>` on Github.

Documentation from source
^^^^^^^^^^^^^^^^^^^^^^^^^

This documentation is written using `Sphinx <https://www.sphinx-doc.org/en/master/index.html>`_
and `reStructuredText <https://www.sphinx-doc.org/en/master/usage/restructuredtext/index.html>`_.

.. note::

  To run documentation from source,
  install ``tox`` using `tox installation guide <https://tox.wiki/en/4.26.0/installation.html>`_
  and execute in root folder:

  .. code-block:: shell

     tox -e docs-auto

FunctionGraph User Guide
------------------------

For FunctionGraph usage, see :otc_fg_umn:`FunctionGraph User Guide <>`.


FunctionGraph Notes and Constraints
-----------------------------------

See :otc_fg_umn:`Notes and Constraints <service_overview/notes_and_constraints.html>`.


.. note:: **Warranty Disclaimer**
    
    THE OPEN SOURCE SOFTWARE IN THIS PRODUCT IS DISTRIBUTED IN THE HOPE THAT IT
    WILL BE USEFUL,BUT WITHOUT ANY WARRANTY; WITHOUT EVEN THE IMPLIED WARRANTY
    OF MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE.

    SEE THE APPLICABLE LICENSES FOR MORE DETAILS.
