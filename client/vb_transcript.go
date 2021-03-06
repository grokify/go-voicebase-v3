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

type VbTranscript struct {
	Confidence float64 `json:"confidence,omitempty"`

	Words []VbWord `json:"words,omitempty"`

	AlternateFormats []VbTranscriptFormat `json:"alternateFormats,omitempty"`
}
