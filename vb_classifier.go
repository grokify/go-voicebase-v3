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

type VbClassifier struct {
	ClassifierId string `json:"classifierId,omitempty"`

	ClassifierName string `json:"classifierName,omitempty"`

	ClassifierDisplayName string `json:"classifierDisplayName,omitempty"`

	ClassifierType string `json:"classifierType,omitempty"`

	PredictedClassLabel string `json:"predictedClassLabel,omitempty"`

	PredictionScore float64 `json:"predictionScore,omitempty"`

	PredictedClass int32 `json:"predictedClass,omitempty"`
}
