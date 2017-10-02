# VbTrackConfiguration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackIndex** | **int32** | Track index. | [optional] [default to null]
**SpeakerName** | **string** | Speaker name. Defaults to \&quot;Speaker\&quot;. If this attribute is specified, all channels are mixed into a single one. | [optional] [default to null]
**SpeakerRole** | **string** | A description of the speaker role. Only used as metadata. | [optional] [default to null]
**Channels** | [**[]VbChannelConfiguration**](VbChannelConfiguration.md) | List of channels to process, mutually exclusive with \&quot;speakerName\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


