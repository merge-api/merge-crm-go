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

// LeadResponse struct for LeadResponse
type LeadResponse struct {
	Model Lead `json:"model"`
	Warnings []WarningValidationProblem `json:"warnings"`
	Errors []ErrorValidationProblem `json:"errors"`
	Logs *[]DebugModeLog `json:"logs,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLeadResponse instantiates a new LeadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeadResponse(model Lead, warnings []WarningValidationProblem, errors []ErrorValidationProblem) *LeadResponse {
	this := LeadResponse{}
	this.Model = model
	this.Warnings = warnings
	this.Errors = errors
	return &this
}

// NewLeadResponseWithDefaults instantiates a new LeadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeadResponseWithDefaults() *LeadResponse {
	this := LeadResponse{}
	return &this
}

// GetModel returns the Model field value
func (o *LeadResponse) GetModel() Lead {
	if o == nil {
		var ret Lead
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *LeadResponse) GetModelOk() (*Lead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *LeadResponse) SetModel(v Lead) {
	o.Model = v
}

// GetWarnings returns the Warnings field value
func (o *LeadResponse) GetWarnings() []WarningValidationProblem {
	if o == nil {
		var ret []WarningValidationProblem
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *LeadResponse) GetWarningsOk() (*[]WarningValidationProblem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value
func (o *LeadResponse) SetWarnings(v []WarningValidationProblem) {
	o.Warnings = v
}

// GetErrors returns the Errors field value
func (o *LeadResponse) GetErrors() []ErrorValidationProblem {
	if o == nil {
		var ret []ErrorValidationProblem
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *LeadResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *LeadResponse) SetErrors(v []ErrorValidationProblem) {
	o.Errors = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *LeadResponse) GetLogs() []DebugModeLog {
	if o == nil || o.Logs == nil {
		var ret []DebugModeLog
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeadResponse) GetLogsOk() (*[]DebugModeLog, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *LeadResponse) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []DebugModeLog and assigns it to the Logs field.
func (o *LeadResponse) SetLogs(v []DebugModeLog) {
	o.Logs = &v
}

func (o LeadResponse) MarshalJSON() ([]byte, error) {
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

func (v *LeadResponse) UnmarshalJSON(src []byte) error {
    type LeadResponseUnmarshalTarget LeadResponse

	var intermediateResult LeadResponseUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = LeadResponse(intermediateResult)
	return nil
}
type NullableLeadResponse struct {
	value *LeadResponse
	isSet bool
}

func (v NullableLeadResponse) Get() *LeadResponse {
	return v.value
}

func (v *NullableLeadResponse) Set(val *LeadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLeadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLeadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeadResponse(val *LeadResponse) *NullableLeadResponse {
	return &NullableLeadResponse{value: val, isSet: true}
}

func (v NullableLeadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


