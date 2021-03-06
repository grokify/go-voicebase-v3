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

type VbCallbackStreamEnum string

// List of VbCallbackStreamEnum
const (
	ORIGINAL       VbCallbackStreamEnum = "original"
	REDACTED_AUDIO VbCallbackStreamEnum = "redacted-audio"
)
