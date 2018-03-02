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

type VbProgress struct {
	Finish string `json:"finish,omitempty"`

	JobId string `json:"jobId,omitempty"`

	Start string `json:"start,omitempty"`

	Status *VbTaskStatusEnum `json:"status,omitempty"`

	Tasks []VbProgressTask `json:"tasks,omitempty"`
}