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

import (
	"time"
)

type VbMediaSummary struct {

	// Media format version. E.g. 3.0.1
	FormatVersion string `json:"formatVersion,omitempty"`

	// Media unique identifier.
	MediaId string `json:"mediaId,omitempty"`

	// Processing state.
	Status *VbStatusEnum `json:"status,omitempty"`

	// Creation timestamp
	DateCreated time.Time `json:"dateCreated,omitempty"`

	// User defined data associated with this record.
	Metadata *VbMetadata `json:"metadata,omitempty"`
}