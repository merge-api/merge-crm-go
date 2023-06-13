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
	"fmt"
)

// TaskStatusEnum * `OPEN` - OPEN * `CLOSED` - CLOSED
type TaskStatusEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of TaskStatusEnum
const (
    TASKSTATUSENUM_MERGE_NONSTANDARD_VALUE TaskStatusEnum = "MERGE_NONSTANDARD_VALUE"
    
	TASKSTATUSENUM_OPEN TaskStatusEnum = "OPEN"
	TASKSTATUSENUM_CLOSED TaskStatusEnum = "CLOSED"
)

var allowedTaskStatusEnumEnumValues = []TaskStatusEnum{
	"OPEN",
	"CLOSED",
}

func (v *TaskStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskStatusEnum(value)
	for _, existing := range allowedTaskStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TASKSTATUSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewTaskStatusEnumFromValue returns a pointer to a valid TaskStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskStatusEnumFromValue(v string) (*TaskStatusEnum, error) {
	ev := TaskStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := TASKSTATUSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskStatusEnum) IsValid() bool {
	for _, existing := range allowedTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskStatusEnum value
func (v TaskStatusEnum) Ptr() *TaskStatusEnum {
	return &v
}

type NullableTaskStatusEnum struct {
	value *TaskStatusEnum
	isSet bool
}

func (v NullableTaskStatusEnum) Get() *TaskStatusEnum {
	return v.value
}

func (v *NullableTaskStatusEnum) Set(val *TaskStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatusEnum(val *TaskStatusEnum) *NullableTaskStatusEnum {
	return &NullableTaskStatusEnum{value: val, isSet: true}
}

func (v NullableTaskStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
