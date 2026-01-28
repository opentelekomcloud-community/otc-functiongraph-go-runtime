# write yaml file with deployed resource info
resource "local_file" "output_yaml" {
  filename = "${path.module}/../tf_deployed_resources.yaml"
  content  = <<-EOT
  FUNCTION_ARN: ${opentelekomcloud_fgs_function_v2.MyFunction.urn}

  EOT

}
