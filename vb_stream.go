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

type VbStream struct {
	Status int32 `json:"status,omitempty"`

	StreamName string `json:"streamName,omitempty"`

	StreamLocation string `json:"streamLocation,omitempty"`

	Comment string `json:"comment,omitempty"`
}
