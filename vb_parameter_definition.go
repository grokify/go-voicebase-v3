/* 
 * Voicebase V3 API
 *
 * APIs for speech recognition and speech analytics, powering insights every business needs.
 *
 * OpenAPI spec version: 3.0.1
 * Contact: support@voicebase.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package voicebase

type VbParameterDefinition struct {

	Name string `json:"name,omitempty"`

	// whether the parameter is optional or not
	Optional bool `json:"optional,omitempty"`

	// Parameter type, one of ( \"String\", \"Integer\", \"Float\", \"Boolean\" )
	Type_ string `json:"type,omitempty"`

	// List of possible values for this parameter, only provided if there is a finite number of values accepted.
	ValueSet []string `json:"valueSet,omitempty"`

	DefaultValue string `json:"defaultValue,omitempty"`
}
