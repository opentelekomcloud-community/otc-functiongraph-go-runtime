package invoke_fg_openstack

const appCode = `
# -*- coding:utf-8 -*-
import json
def handler (event, context):
    return {
        "statusCode": 200,
        "isBase64Encoded": False,
        "body": json.dumps(event),
        "headers": {
            "Content-Type": "application/json"
        }
    }
`

const functionApp = "default"
const functionName = "DefaultPython3_10_From_Go_SDK"
const functionVersion = "latest"
