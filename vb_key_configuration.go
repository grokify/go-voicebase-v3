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

type VbKeyConfiguration struct {

	// Number of milliseconds that the key must be valid for
	TtlMillis int64 `json:"ttlMillis,omitempty"`

	// Expiration date expressed in ISO-8601 format
	ExpirationDate time.Time `json:"expirationDate,omitempty"`

	Ephemeral bool `json:"ephemeral,omitempty"`

	// Restriction of the scope of the token in the form of media:get?extendedFilter=speakers:agent
	Scope string `json:"scope,omitempty"`
}
