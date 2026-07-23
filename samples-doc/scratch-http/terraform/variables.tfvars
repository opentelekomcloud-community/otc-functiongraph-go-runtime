# Terraform variables

# prefix of all resources
prefix        = "go"

# name of the function (will be prefixed)
function_name = "go-doc-scratch-http"

# resources will be tagged with this app_group tag
tag_app_group = "go-doc-scratch-http"

# name of zip file to deploy
zip_file_name = "deploy.zip"
    
# change to your API Gateway instance ID
# set as env var TF_VAR_API_GATEWAY_INSTANCE_ID or uncomment and set here
#API_GATEWAY_INSTANCE_ID="YOUR_API_GATEWAY_INSTANCE_ID"