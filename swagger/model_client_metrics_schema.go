/*
 * Unleash API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.2.0-main
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Client usage metrics, accumulated in buckets of hour by hour by default
type ClientMetricsSchema struct {
	// The name of the application that is evaluating toggles
	AppName string `json:"appName"`
	// A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application
	InstanceId string `json:"instanceId,omitempty"`
	// Which environment the application is running in
	Environment string                     `json:"environment,omitempty"`
	Bucket      *ClientMetricsSchemaBucket `json:"bucket"`
}
