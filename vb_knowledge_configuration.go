/*
 * Voicebase V3 API
 *
 * APIs for speech recognition and speech analytics, powering insights every business needs.
 *
 * API version: 3.0.1
 * Contact: support@voicebase.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package voicebase

type VbKnowledgeConfiguration struct {

	// Whether knowledge discovery is enabled or not
	EnableDiscovery bool `json:"enableDiscovery,omitempty"`

	// Whether semantic discovery is allowed to use external data sources or not
	EnableExternalDataSources bool `json:"enableExternalDataSources,omitempty"`
}
