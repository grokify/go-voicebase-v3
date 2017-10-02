/* 
 * Voicebase V3 API
 *
 * APIs for speech recognition and speech analytics, powering insights every business needs.
 *
 * OpenAPI spec version: 3.0.1
 * Contact: support@voicebase.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type VbClassifierModel struct {

	ClassifierId string `json:"classifierId,omitempty"`

	// Use this name for referring to the classifier in a configuration provided with the media for processing
	ClassifierName string `json:"classifierName,omitempty"`

	// Use this version in conjuction with the classifier name for referring to this classifier in a configuration provided with the media for processing
	ClassifierVersion string `json:"classifierVersion,omitempty"`

	ClassifierDescription string `json:"classifierDescription,omitempty"`

	ClassifierType string `json:"classifierType,omitempty"`

	// Set of possible classes identified by the classifier
	Classes []VbClass `json:"classes,omitempty"`

	// Set of possible parameters for this classifier
	Parameters []VbParameterDefinition `json:"parameters,omitempty"`
}
