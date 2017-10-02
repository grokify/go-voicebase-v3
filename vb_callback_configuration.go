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

type VbCallbackConfiguration struct {

	Url string `json:"url,omitempty"`

	Method VbHttpMethodEnum `json:"method,omitempty"`

	Type_ VbCallbackTypeEnum `json:"type,omitempty"`

	Include []VbIncludeTypeEnum `json:"include,omitempty"`

	Stream VbCallbackStreamEnum `json:"stream,omitempty"`

	Format VbCallbackFormatEnum `json:"format,omitempty"`
}
