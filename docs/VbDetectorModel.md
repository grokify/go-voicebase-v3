# VbDetectorModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectorId** | **string** |  | [optional] [default to null]
**DetectorName** | **string** | Use this detector name when refering to this detector in a configuration for media processing | [optional] [default to null]
**DetectorVersion** | **string** | Use this version in conjuction with the detector name for referring to this detector in a configuration provided with the media for processing | [optional] [default to null]
**DetectorDescription** | **string** | Describes the function of this detector and its restrictions | [optional] [default to null]
**DetectorType** | **string** | Detector type, one of ( &#39;binary&#39;, &#39;multiclass&#39;).  Binary detectors only report positive cases. | [optional] [default to null]
**Classes** | [**[]VbClass**](VbClass.md) | Set of possible classes for segments identified by this detector | [optional] [default to null]
**Parameters** | [**[]VbParameterDefinition**](VbParameterDefinition.md) | Set of possible parameters for this detector | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


