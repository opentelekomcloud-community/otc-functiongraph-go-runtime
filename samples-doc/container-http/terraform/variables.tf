# prefix will be prepended to all resource names
variable "prefix" {
  type    = string
  default = "sample"
}

# FunctionGraph: Function name
variable "function_name" {
  type    = string
  default = "function_name"
}

variable "image_url" {
  type    = string
  default = "your_image_url_here"
}

# Resource tag:
variable "tag_app_group" {
  type    = string
  default = "your_tag_app_group_here"
}

variable "API_GATEWAY_INSTANCE_ID" {
  type = string
  default = "your_api_gateway_instance_id"
}