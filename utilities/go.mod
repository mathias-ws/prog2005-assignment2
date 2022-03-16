module utilities/v1

go 1.17

require github.com/stretchr/testify v1.7.0

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace web_client => ../web_client
replace constants => ../constants
replace custom_errors => ../custom_errors
replace logic => ../logic
replace utilities => ../utilities
replace web_client => ../web_client
replace web_server => ../web_server
replace assignment-2 => ../assignment-2
replace status_endpoint => ../status_endpoint
replace policy_endpoint => ../polic_endpoint