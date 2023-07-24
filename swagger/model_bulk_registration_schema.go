/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An application registration. Defines the format POSTed by our server-side SDKs when they're starting up
type BulkRegistrationSchema struct {
	// A list of applications this app registration has been registered through. If connected directly to Unleash, this is an empty list.   This can be used in later visualizations to tell how many levels of proxy or Edge instances our SDKs have connected through
	ConnectVia []BulkRegistrationSchemaConnectVia `json:"connectVia,omitempty"`
	// The name of the application that is evaluating toggles
	AppName string `json:"appName"`
	// Which environment the application is running in
	Environment string `json:"environment"`
	// A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application
	InstanceId string `json:"instanceId"`
	// How often (in seconds) the application refreshes its features
	Interval float64     `json:"interval,omitempty"`
	Started  *DateSchema `json:"started,omitempty"`
	// Enabled [strategies](https://docs.getunleash.io/reference/activation-strategies) in the application
	Strategies []string `json:"strategies,omitempty"`
	// The version the sdk is running. Typically <client>:<version>
	SdkVersion string `json:"sdkVersion,omitempty"`
}