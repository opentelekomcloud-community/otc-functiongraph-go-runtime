##########################################################
# Create Function
##########################################################
resource "opentelekomcloud_fgs_function_v2" "MyFunction" {
  name = format("%s_%s", var.prefix, var.function_name)
  app  = "default"
  #  agency      = var.agency_name
  handler          = var.function_handler_name
  description      = "Sample for timer triggered FunctionGraph using terraform."
  memory_size      = 128
  timeout          = 30
  max_instance_num = 400

  runtime       = "Go1.x"
  code_type     = "zip"
  func_code     = filebase64(format("${path.module}/../%s", var.zip_file_name))
  code_filename = var.zip_file_name

  log_group_id   = opentelekomcloud_lts_group_v2.MyLogGroup.id
  log_group_name = opentelekomcloud_lts_group_v2.MyLogGroup.group_name

  log_topic_id   = opentelekomcloud_lts_stream_v2.MyLogStream.id
  log_topic_name = opentelekomcloud_lts_stream_v2.MyLogStream.stream_name

}
