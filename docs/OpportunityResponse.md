# OpportunityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**Opportunity**](Opportunity.md) |  | 
**Warnings** | [**[]WarningValidationProblem**](WarningValidationProblem.md) |  | 
**Errors** | [**[]ErrorValidationProblem**](ErrorValidationProblem.md) |  | 
**Logs** | Pointer to [**[]DebugModeLog**](DebugModeLog.md) |  | [optional] 

## Methods

### NewOpportunityResponse

`func NewOpportunityResponse(model Opportunity, warnings []WarningValidationProblem, errors []ErrorValidationProblem, ) *OpportunityResponse`

NewOpportunityResponse instantiates a new OpportunityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityResponseWithDefaults

`func NewOpportunityResponseWithDefaults() *OpportunityResponse`

NewOpportunityResponseWithDefaults instantiates a new OpportunityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *OpportunityResponse) GetModel() Opportunity`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpportunityResponse) GetModelOk() (*Opportunity, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpportunityResponse) SetModel(v Opportunity)`

SetModel sets Model field to given value.


### GetWarnings

`func (o *OpportunityResponse) GetWarnings() []WarningValidationProblem`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *OpportunityResponse) GetWarningsOk() (*[]WarningValidationProblem, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *OpportunityResponse) SetWarnings(v []WarningValidationProblem)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *OpportunityResponse) GetErrors() []ErrorValidationProblem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OpportunityResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OpportunityResponse) SetErrors(v []ErrorValidationProblem)`

SetErrors sets Errors field to given value.


### GetLogs

`func (o *OpportunityResponse) GetLogs() []DebugModeLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *OpportunityResponse) GetLogsOk() (*[]DebugModeLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *OpportunityResponse) SetLogs(v []DebugModeLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *OpportunityResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


