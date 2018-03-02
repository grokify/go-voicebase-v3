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

type VbWord struct {
	P int32 `json:"p,omitempty"`

	C float64 `json:"c,omitempty"`

	S int64 `json:"s,omitempty"`

	E int64 `json:"e,omitempty"`

	Frq []VbFrequency `json:"frq,omitempty"`

	M *VbWordTypeEnum `json:"m,omitempty"`

	V float64 `json:"v,omitempty"`

	W string `json:"w,omitempty"`
}