Sample for Timer Triggered FunctionGraph
================================================

This sample demonstrates how to create a timer triggered FunctionGraph
using Go runtime and deploy it using Terraform.

Source for this sample can be found in:
:github_repo_master:`/samples-doc/doc-sample-event-timer <samples-doc/doc-sample-event-timer>`.

Overview
--------
Following diagram shows components used in this example:
  



The timer event triggers the FunctionGraph function at specified intervals.

Deployment
----------
This sample can be deployed using Terraform,
see :ref:`ref_terraform_setup` for setup details.

Terraform deployment scripts can be found in:
:github_repo_master:`/samples-doc/doc-sample-event-timer/terraform <samples-doc/doc-sample-event-timer/terraform>`