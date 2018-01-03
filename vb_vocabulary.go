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

type VbVocabulary struct {
	VocabularyName string `json:"vocabularyName,omitempty"`

	VocabularyType *VbVocabularyType `json:"vocabularyType,omitempty"`

	Terms []VbVocabularyTerm `json:"terms,omitempty"`

	Scripts []VbVocabularyScript `json:"scripts,omitempty"`
}
