module invoke-fg-openstack

go 1.17

// use dedicated devel branch until PR#907 is released
require github.com/opentelekomcloud/gophertelekomcloud v0.9.6-0.20260122085539-793ddb638df5

require gopkg.in/yaml.v2 v2.3.0 // indirect

// replace github.com/opentelekomcloud/gophertelekomcloud => ../../../../opentelekomcloud/gophertelekomcloud
