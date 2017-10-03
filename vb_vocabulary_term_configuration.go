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

type VbVocabularyTermConfiguration struct {
	Term string `json:"term,omitempty"`

	SoundsLike []string `json:"soundsLike,omitempty"`

	Weight int32 `json:"weight,omitempty"`
}
