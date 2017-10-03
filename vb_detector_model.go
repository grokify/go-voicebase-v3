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

type VbDetectorModel struct {
	DetectorId string `json:"detectorId,omitempty"`

	// Use this detector name when refering to this detector in a configuration for media processing
	DetectorName string `json:"detectorName,omitempty"`

	// Use this version in conjuction with the detector name for referring to this detector in a configuration provided with the media for processing
	DetectorVersion string `json:"detectorVersion,omitempty"`

	// Describes the function of this detector and its restrictions
	DetectorDescription string `json:"detectorDescription,omitempty"`

	// Detector type, one of ( 'binary', 'multiclass').  Binary detectors only report positive cases.
	DetectorType string `json:"detectorType,omitempty"`

	// Set of possible classes for segments identified by this detector
	Classes []VbClass `json:"classes,omitempty"`

	// Set of possible parameters for this detector
	Parameters []VbParameterDefinition `json:"parameters,omitempty"`
}
