# write yaml file with deployed resource info
resource "local_file" "output_yaml" {
  filename = "${path.module}/../tf_deployed_resources.yaml"
  content  = <<-EOT
  FUNCTION_ARN: ${opentelekomcloud_fgs_function_v2.MyFunction.urn}
  OBS_INPUT_BUCKET: ${opentelekomcloud_s3_bucket.inbucket.bucket}
  OBS_OUTPUT_BUCKET: ${opentelekomcloud_s3_bucket.outbucket.bucket}  
  EOT
  
}