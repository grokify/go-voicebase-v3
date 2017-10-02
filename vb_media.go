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

import (
	"time"
)

type VbMedia struct {

	// Media format version. E.g. 3.0.1
	FormatVersion string `json:"formatVersion,omitempty"`

	// Media unique identifier.
	MediaId string `json:"mediaId,omitempty"`

	// Processing state.
	Status VbStatusEnum `json:"status,omitempty"`

	// Creation timestamp
	DateCreated time.Time `json:"dateCreated,omitempty"`

	// User defined data associated with this record.
	Metadata VbMetadata `json:"metadata,omitempty"`

	// The MIME type of the media submitted for processing. E.g. audio/x-wav, audio/mpeg
	MediaContentType string `json:"mediaContentType,omitempty"`

	// Duration of the audio/video submitted for processing expressed in milliseconds
	Length int64 `json:"length,omitempty"`

	// Semantic knowledge discovery section. If knoweledge discovery was requested, this section contains the results.
	Knowledge VbKnowledge `json:"knowledge,omitempty"`

	// If kewyword spotting was requested, this section contains the results.
	Spotting VbSpotting `json:"spotting,omitempty"`

	// If any predictions (classifiers, detectors) were requested, this section contains the results.
	Prediction VbPrediction `json:"prediction,omitempty"`

	// This section contains the transcript in a variety of formats
	Transcript VbTranscript `json:"transcript,omitempty"`

	// This section contains warnings about the media
	Warnings []VbMessage `json:"warnings,omitempty"`

	// Theaudio/video streams available.
	Streams []VbStream `json:"streams,omitempty"`

	Job VbJob `json:"_job,omitempty"`
}
