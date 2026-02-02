
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