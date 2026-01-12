

###########################################################
# Custom role to allow FunctionGraph to access LTS and OBS
###########################################################
resource "opentelekomcloud_identity_role_v3" "role" {
  display_name  = format("%s-%s-role", var.prefix, var.function_name)
  description   = "Role for FunctionGraph to access OBS"
  display_layer = "project"

  statement {
    effect = "Allow"
    action = [
      "functiongraph:*:*",
      "lts:*:*",
    ]
  }

  statement {
    effect = "Allow"
    action = [
      "obs:*:*",
    ]
    resource = [
      "OBS:*:*:object:*",
      format("OBS:*:*:bucket:%s", opentelekomcloud_s3_bucket.inbucket.bucket),
      format("OBS:*:*:bucket:%s", opentelekomcloud_s3_bucket.outbucket.bucket),
    ]
  }

}

##########################################################
# Agency for FunctionGraph
# Attention: Crating agency will take some time.
# Calls to function after creating agency will fail until
# agency is set up.
##########################################################
resource "opentelekomcloud_identity_agency_v3" "agency" {
  depends_on            = [opentelekomcloud_identity_role_v3.role]
  delegated_domain_name = "op_svc_cff"

  name        = format("%s-%s-agency", var.prefix, var.function_name)
  description = "Agency for FunctionGraph to access OBS"

  project_role {
    all_projects = true
    project      = "eu-de"
    roles = [
      opentelekomcloud_identity_role_v3.role.display_name
    ]
  }

}


##########################################################
# Create Function
##########################################################
resource "opentelekomcloud_fgs_function_v2" "MyFunction" {
  # depends_on       = [opentelekomcloud_obs_bucket_object.code_object]
  name             = format("%s_%s", var.prefix, var.function_name)
  app              = "default"
  agency           = opentelekomcloud_identity_agency_v3.agency.name
  handler          = var.function_handler_name

  description      = "Sample on how to create Thumbnails from images uploaded to OBS"
  memory_size      = 512
  timeout          = 30
  max_instance_num = 1
  runtime = var.function_runtime

  code_type     = "zip"
  func_code     = filebase64(var.zip_file_local)

  code_filename = basename(var.zip_file_local)

  log_group_id   = opentelekomcloud_lts_group_v2.MyLogGroup.id
  log_group_name = opentelekomcloud_lts_group_v2.MyLogGroup.group_name

  log_topic_id   = opentelekomcloud_lts_stream_v2.MyLogStream.id
  log_topic_name = opentelekomcloud_lts_stream_v2.MyLogStream.stream_name

  # set some environment variables
  user_data = jsonencode({
    "OUTPUT_BUCKET" : opentelekomcloud_s3_bucket.outbucket.bucket,
    "OBS_ENDPOINT" : "https://obs.otc.t-systems.com",
    # "RUNTIME_LOG_LEVEL" : "ERROR",
    # "RUNTIME_LOG_PATH" : "/tmp"
  })

  tags = {
    "app_group" = var.tag_app_group
  }

}

##########################################################
# Create Log Group
##########################################################
resource "opentelekomcloud_lts_group_v2" "MyLogGroup" {
  group_name  = format("%s_%s_%s", var.prefix, var.function_name, "log_group")
  ttl_in_days = 1

  tags = {
    "app_group" = var.tag_app_group
  }
}

##########################################################
# Create Log Stream
##########################################################
resource "opentelekomcloud_lts_stream_v2" "MyLogStream" {
  group_id    = opentelekomcloud_lts_group_v2.MyLogGroup.id
  stream_name = format("%s_%s_%s", var.prefix, var.function_name, "log_stream")

  tags = {
    "app_group" = var.tag_app_group
  }
}

##########################################################
# Input bucket for source images
##########################################################
resource "opentelekomcloud_s3_bucket" "inbucket" {
  bucket        = lower(format("%s-%s-%s", var.prefix, var.function_name, "images"))
  acl           = "private"

  # Warning: force_destroy will delete bucket on 
  # terraform destroy even if it contains objects
  force_destroy = true

  tags = {
    "app_group" = var.tag_app_group
  }

}

##########################################################
# Output bucket for thumbnail images
# For output a different bucket is used to avoid potential
# risk of recursive invocation of FunctionGraph
##########################################################
resource "opentelekomcloud_s3_bucket" "outbucket" {
  bucket        = lower(format("%s-%s-%s", var.prefix, var.function_name, "images-output"))
  acl           = "private"

  # Warning: force_destroy will delete bucket on 
  # terraform destroy even if it contains objects
  force_destroy = true

  tags = {
    "app_group" = var.tag_app_group
  }

}

##########################################################
# Create OBS Trigger listening for "ObjectCreate" in
# input bucket
##########################################################
resource "opentelekomcloud_fgs_trigger_v2" "obstrigger" {
  function_urn = opentelekomcloud_fgs_function_v2.MyFunction.urn
  type         = "OBS"
  event_data = jsonencode({
    "bucket" : opentelekomcloud_s3_bucket.inbucket.bucket
    "events" : [
      "s3:ObjectCreated:*"
    ]
    "name" : lower(format("%s-%s-%s", var.prefix, var.function_name, "event"))

  })
}


##########################################################
# Create Test Event
##########################################################
resource "opentelekomcloud_fgs_event_v2" "test_event" {
  function_urn = opentelekomcloud_fgs_function_v2.MyFunction.urn
  name         = "UploadTest"
  content = base64encode(jsonencode({
    "Records" = [{
      "eventVersion" = "2.0"
      "eventSource"  = "obs"
      "eventTime"    = "2025-10-24T08:30:00+08:00"
      "eventName"    = "ObjectCreated:PutObject"
      "awsRegion"    = "eu-de"
      "userIdentity" = {
        "principalId" = "EXAMPLE"
      }
      "requestParameters" = {
        "sourceIPAddress" = "EXAMPLE"
      }
      "s3" = {
        "configurationId" = "testConfigRule"
        "bucket" = {
          "name" = opentelekomcloud_s3_bucket.inbucket.id
          "ownerIdentity" = {
            "principalId" = "EXAMPLE"
          }
          "arn" = opentelekomcloud_s3_bucket.inbucket.arn
        }
        "object" = {
          "key" = "image.jpg"
          "size" = 1024
          "eTag" = "0123456789abcdef0123456789abcdef"
          "sequencer" = "0A1B2C3D4E5F678901"
        }
      }
    }]
  }))
}


output "INPUT_BUCKET" {
  value = opentelekomcloud_s3_bucket.inbucket.bucket
}
output "OUTPUT_BUCKET" {
  value = opentelekomcloud_s3_bucket.outbucket.bucket
}