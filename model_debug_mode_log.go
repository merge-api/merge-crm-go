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

// DebugModeLog struct for DebugModeLog
type DebugModeLog struct {
	LogId string `json:"log_id"`
	DashboardView string `json:"dashboard_view"`
	LogSummary DebugModelLogSummary `json:"log_summary"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewDebugModeLog instantiates a new DebugModeLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebugModeLog(logId string, dashboardView string, logSummary DebugModelLogSummary) *DebugModeLog {
	this := DebugModeLog{}
	this.LogId = logId
	this.DashboardView = dashboardView
	this.LogSummary = logSummary
	return &this
}

// NewDebugModeLogWithDefaults instantiates a new DebugModeLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebugModeLogWithDefaults() *DebugModeLog {
	this := DebugModeLog{}
	return &this
}

// GetLogId returns the LogId field value
func (o *DebugModeLog) GetLogId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value
// and a boolean to check if the value has been set.
func (o *DebugModeLog) GetLogIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogId, true
}

// SetLogId sets field value
func (o *DebugModeLog) SetLogId(v string) {
	o.LogId = v
}

// GetDashboardView returns the DashboardView field value
func (o *DebugModeLog) GetDashboardView() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardView
}

// GetDashboardViewOk returns a tuple with the DashboardView field value
// and a boolean to check if the value has been set.
func (o *DebugModeLog) GetDashboardViewOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DashboardView, true
}

// SetDashboardView sets field value
func (o *DebugModeLog) SetDashboardView(v string) {
	o.DashboardView = v
}

// GetLogSummary returns the LogSummary field value
func (o *DebugModeLog) GetLogSummary() DebugModelLogSummary {
	if o == nil {
		var ret DebugModelLogSummary
		return ret
	}

	return o.LogSummary
}

// GetLogSummaryOk returns a tuple with the LogSummary field value
// and a boolean to check if the value has been set.
func (o *DebugModeLog) GetLogSummaryOk() (*DebugModelLogSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LogSummary, true
}

// SetLogSummary sets field value
func (o *DebugModeLog) SetLogSummary(v DebugModelLogSummary) {
	o.LogSummary = v
}

func (o DebugModeLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["log_id"] = o.LogId
	}
	if true {
		toSerialize["dashboard_view"] = o.DashboardView
	}
	if true {
		toSerialize["log_summary"] = o.LogSummary
	}
	return json.Marshal(toSerialize)
}

func (v *DebugModeLog) UnmarshalJSON(src []byte) error {
    type DebugModeLogUnmarshalTarget DebugModeLog

	var intermediateResult DebugModeLogUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = DebugModeLog(intermediateResult)
	return nil
}
type NullableDebugModeLog struct {
	value *DebugModeLog
	isSet bool
}

func (v NullableDebugModeLog) Get() *DebugModeLog {
	return v.value
}

func (v *NullableDebugModeLog) Set(val *DebugModeLog) {
	v.value = val
	v.isSet = true
}

func (v NullableDebugModeLog) IsSet() bool {
	return v.isSet
}

func (v *NullableDebugModeLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebugModeLog(val *DebugModeLog) *NullableDebugModeLog {
	return &NullableDebugModeLog{value: val, isSet: true}
}

func (v NullableDebugModeLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebugModeLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


