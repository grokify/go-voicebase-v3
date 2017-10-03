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

type VbTopic struct {
	TopicName string `json:"topicName,omitempty"`

	Relevance float64 `json:"relevance,omitempty"`

	SubTopics []VbReference `json:"subTopics,omitempty"`

	Keywords []VbKeyword `json:"keywords,omitempty"`
}
