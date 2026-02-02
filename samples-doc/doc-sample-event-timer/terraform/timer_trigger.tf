##########################################################
# Create Trigger
##########################################################
resource "opentelekomcloud_fgs_trigger_v2" "timer_cron" {
  function_urn = opentelekomcloud_fgs_function_v2.MyFunction.urn
  type         = "TIMER"
  event_data = jsonencode({
    "name" : "CronTrigger_3m",
    "schedule_type" : "Cron",
    "user_event" : "Created by terraform script",
    "schedule" : "@every 3m"
  })
}
