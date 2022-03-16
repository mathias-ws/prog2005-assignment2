module status_endpoint/v1

go 1.17

require web_client v1.0.0

require constants v1.0.0

require custom_errors v1.0.0 // indirect

replace constants => ../constants

replace custom_errors => ../custom_errors

replace logic => ../logic

replace utilities => ../utilities

replace web_client => ../web_client

replace web_server => ../web_server

replace assignment-2 => ../assignment-2

replace status_endpoint => ../status_endpoint

replace policy_endpoint => ../polic_endpoint
