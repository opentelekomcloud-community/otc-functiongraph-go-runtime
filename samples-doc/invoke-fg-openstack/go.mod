module invoke-fg-openstack

go 1.17

// use dedicated branch until gophertelekomcloud fixes for invoke are merged
require github.com/opentelekomcloud/gophertelekomcloud v0.9.6-0.20260121093111-7e86a575a4b3

require gopkg.in/yaml.v2 v2.3.0 // indirect

// replace github.com/opentelekomcloud/gophertelekomcloud => ../../../../opentelekomcloud/gophertelekomcloud
