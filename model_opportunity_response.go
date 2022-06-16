/*
 * Merge CRM API
 *
 * The unified API for building rich integrations with multiple CRM platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_crm_client

import (
	"encoding/json"
)

// OpportunityResponse struct for OpportunityResponse
type OpportunityResponse struct {
	Model Opportunity `json:"model"`
	Warnings []WarningValidationProblem `json:"warnings"`
	Errors []ErrorValidationProblem `json:"errors"`
	Logs *[]DebugModeLog `json:"logs,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewOpportunityResponse instantiates a new OpportunityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpportunityResponse(model Opportunity, warnings []WarningValidationProblem, errors []ErrorValidationProblem) *OpportunityResponse {
	this := OpportunityResponse{}
	this.Model = model
	this.Warnings = warnings
	this.Errors = errors
	return &this
}

// NewOpportunityResponseWithDefaults instantiates a new OpportunityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpportunityResponseWithDefaults() *OpportunityResponse {
	this := OpportunityResponse{}
	return &this
}

// GetModel returns the Model field value
func (o *OpportunityResponse) GetModel() Opportunity {
	if o == nil {
		var ret Opportunity
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *OpportunityResponse) GetModelOk() (*Opportunity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *OpportunityResponse) SetModel(v Opportunity) {
	o.Model = v
}

// GetWarnings returns the Warnings field value
func (o *OpportunityResponse) GetWarnings() []WarningValidationProblem {
	if o == nil {
		var ret []WarningValidationProblem
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *OpportunityResponse) GetWarningsOk() (*[]WarningValidationProblem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value
func (o *OpportunityResponse) SetWarnings(v []WarningValidationProblem) {
	o.Warnings = v
}

// GetErrors returns the Errors field value
func (o *OpportunityResponse) GetErrors() []ErrorValidationProblem {
	if o == nil {
		var ret []ErrorValidationProblem
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *OpportunityResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *OpportunityResponse) SetErrors(v []ErrorValidationProblem) {
	o.Errors = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *OpportunityResponse) GetLogs() []DebugModeLog {
	if o == nil || o.Logs == nil {
		var ret []DebugModeLog
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpportunityResponse) GetLogsOk() (*[]DebugModeLog, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *OpportunityResponse) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []DebugModeLog and assigns it to the Logs field.
func (o *OpportunityResponse) SetLogs(v []DebugModeLog) {
	o.Logs = &v
}

func (o OpportunityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

func (v *OpportunityResponse) UnmarshalJSON(src []byte) error {
    type OpportunityResponseUnmarshalTarget OpportunityResponse

	var intermediateResult OpportunityResponseUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = OpportunityResponse(intermediateResult)
	return nil
}
type NullableOpportunityResponse struct {
	value *OpportunityResponse
	isSet bool
}

func (v NullableOpportunityResponse) Get() *OpportunityResponse {
	return v.value
}

func (v *NullableOpportunityResponse) Set(val *OpportunityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpportunityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpportunityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpportunityResponse(val *OpportunityResponse) *NullableOpportunityResponse {
	return &NullableOpportunityResponse{value: val, isSet: true}
}

func (v NullableOpportunityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpportunityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


