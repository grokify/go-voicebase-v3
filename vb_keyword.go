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

type VbKeyword struct {
	Keyword string `json:"keyword,omitempty"`

	Relevance float64 `json:"relevance,omitempty"`

	Mentions []VbMention `json:"mentions,omitempty"`
}
