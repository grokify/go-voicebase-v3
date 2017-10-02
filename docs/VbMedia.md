# VbMedia

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatVersion** | **string** | Media format version. E.g. 3.0.1 | [optional] [default to null]
**MediaId** | **string** | Media unique identifier. | [optional] [default to null]
**Status** | [**VbStatusEnum**](VbStatusEnum.md) | Processing state. | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) | Creation timestamp | [optional] [default to null]
**Metadata** | [**VbMetadata**](VbMetadata.md) | User defined data associated with this record. | [optional] [default to null]
**MediaContentType** | **string** | The MIME type of the media submitted for processing. E.g. audio/x-wav, audio/mpeg | [optional] [default to null]
**Length** | **int64** | Duration of the audio/video submitted for processing expressed in milliseconds | [optional] [default to null]
**Knowledge** | [**VbKnowledge**](VbKnowledge.md) | Semantic knowledge discovery section. If knoweledge discovery was requested, this section contains the results. | [optional] [default to null]
**Spotting** | [**VbSpotting**](VbSpotting.md) | If kewyword spotting was requested, this section contains the results. | [optional] [default to null]
**Prediction** | [**VbPrediction**](VbPrediction.md) | If any predictions (classifiers, detectors) were requested, this section contains the results. | [optional] [default to null]
**Transcript** | [**VbTranscript**](VbTranscript.md) | This section contains the transcript in a variety of formats | [optional] [default to null]
**Warnings** | [**[]VbMessage**](VbMessage.md) | This section contains warnings about the media | [optional] [default to null]
**Streams** | [**[]VbStream**](VbStream.md) | Theaudio/video streams available. | [optional] [default to null]
**Job** | [**VbJob**](VbJob.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


