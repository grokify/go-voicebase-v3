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

type VbConfiguration struct {
	SpeechModel *VbSpeechModelConfiguration `json:"speechModel,omitempty"`

	// Allows to specify other languages to enable recognition when conversations occur in multiple languages
	AdditionalSpeechModels []VbSpeechModelConfiguration `json:"additionalSpeechModels,omitempty"`

	Priority *VbPriorityEnum `json:"priority,omitempty"`

	Ingest *VbIngestConfiguration `json:"ingest,omitempty"`

	Prediction *VbPredictionConfiguration `json:"prediction,omitempty"`

	Spotting *VbSpottingConfiguration `json:"spotting,omitempty"`

	Knowledge *VbKnowledgeConfiguration `json:"knowledge,omitempty"`

	Transcript *VbTranscriptConfiguration `json:"transcript,omitempty"`

	Vocabularies []VbVocabularyConfiguration `json:"vocabularies,omitempty"`

	Publish *VbPublishConfiguration `json:"publish,omitempty"`
}
