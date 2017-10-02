/* 
 * Voicebase V3 API
 *
 * APIs for speech recognition and speech analytics, powering insights every business needs.
 *
 * OpenAPI spec version: 3.0.1
 * Contact: support@voicebase.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type VbDetectedSegment struct {

	SpeakerName string `json:"speakerName,omitempty"`

	Occurrences []VbOccurrence `json:"occurrences,omitempty"`
}
