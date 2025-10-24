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

##########################################################
# Create Log Group
##########################################################
resource "opentelekomcloud_lts_group_v2" "MyLogGroup" {
  group_name  = format("%s_%s_%s", var.prefix, var.function_name, "log_group")
  ttl_in_days = 1
}

##########################################################
# Create Log Stream
##########################################################
resource "opentelekomcloud_lts_stream_v2" "MyLogStream" {
  group_id    = opentelekomcloud_lts_group_v2.MyLogGroup.id
  stream_name = format("%s_%s_%s", var.prefix, var.function_name, "log_stream")
}

##########################################################
# Create Test Event
##########################################################
resource "opentelekomcloud_fgs_event_v2" "test_event" {
  function_urn = opentelekomcloud_fgs_function_v2.MyFunction.urn
  name         = "timer-test-event"
  content = base64encode(jsonencode({
    "version" = "v1.0",
    "time" = "2025-10-24T08:30:00+08:00",
    "trigger_type" = "TIMER",
    "trigger_name" = "timer-test-event",
    "user_event" = "Test event created by terraform script"
  }))
}


##########################################################
# Create Trigger
##########################################################
resource "opentelekomcloud_fgs_trigger_v2" "timer_cron" {
  function_urn = opentelekomcloud_fgs_function_v2.MyFunction.urn
  type         = "TIMER"
  event_data = jsonencode({
    "name" : "CronTrigger_3m"
    "schedule_type" : "Cron",
    "user_event" : "Created by terraform script",
    "schedule" : "@every 3m"
  })
}