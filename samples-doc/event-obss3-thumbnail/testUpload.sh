#!/bin/bash

##############################################################################################
# This script uploads a test image to the OBS bucket to trigger the function
# Make sure you have
# - s3cmd installed and configured
#   Install s3cmd: https://s3tools.org/s3cmd
#   see also: # https://github.com/opentelekomcloud/obs-s3/blob/master/s3cmd/README.md
#
# - yq installed for parsing yaml
#   see: https://github.com/mikefarah/yq
#
##############################################################################################

TF_OUTPUT_FILE="tf_deployed_resources.yaml"

if [ ! -f ${TF_OUTPUT_FILE} ]; then
    echo "File ${TF_OUTPUT_FILE} not found! Deploy using terraform first."
else

  # For bucket name see output of terraform output
  # following command assumes that terraform output is stored in tf_deployed_resources.yaml
  OBS_INPUT_BUCKET=$(yq -r '.OBS_INPUT_BUCKET' ${TF_OUTPUT_FILE})

  echo "Uploading test image to bucket: $OBS_INPUT_BUCKET"

  s3cmd --access_key=${OTC_SDK_AK} --secret_key=${OTC_SDK_SK} --no-ssl \
    put ./resources/otc.jpg \
    s3://$OBS_INPUT_BUCKET/otc.jpg

fi