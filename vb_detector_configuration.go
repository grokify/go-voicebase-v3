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

type VbDetectorConfiguration struct {

	// Detector identifier, a UUID
	DetectorId string `json:"detectorId,omitempty"`

	// A name identifying a detector offered by Voicebase to all customers
	DetectorName string `json:"detectorName,omitempty"`

	// Provide a specific version (e.g. \"3.1.3\"), specify a minimum (e.g. \"3.1+\") or get the \"latest\". Defaults to \"latest\" if not specified.
	Version string `json:"version,omitempty"`

	// List of parameter values to the detector
	Parameters []VbParameter `json:"parameters,omitempty"`

	Redactor VbRedactorConfiguration `json:"redactor,omitempty"`
}
