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

// ConditionTypeEnum * `BOOLEAN` - BOOLEAN * `DATE` - DATE * `DATE_TIME` - DATE_TIME * `INTEGER` - INTEGER * `FLOAT` - FLOAT * `STRING` - STRING * `LIST_OF_STRINGS` - LIST_OF_STRINGS
type ConditionTypeEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of ConditionTypeEnum
const (
    CONDITIONTYPEENUM_MERGE_NONSTANDARD_VALUE ConditionTypeEnum = "MERGE_NONSTANDARD_VALUE"
    
	CONDITIONTYPEENUM_BOOLEAN ConditionTypeEnum = "BOOLEAN"
	CONDITIONTYPEENUM_DATE ConditionTypeEnum = "DATE"
	CONDITIONTYPEENUM_DATE_TIME ConditionTypeEnum = "DATE_TIME"
	CONDITIONTYPEENUM_INTEGER ConditionTypeEnum = "INTEGER"
	CONDITIONTYPEENUM_FLOAT ConditionTypeEnum = "FLOAT"
	CONDITIONTYPEENUM_STRING ConditionTypeEnum = "STRING"
	CONDITIONTYPEENUM_LIST_OF_STRINGS ConditionTypeEnum = "LIST_OF_STRINGS"
)

var allowedConditionTypeEnumEnumValues = []ConditionTypeEnum{
	"BOOLEAN",
	"DATE",
	"DATE_TIME",
	"INTEGER",
	"FLOAT",
	"STRING",
	"LIST_OF_STRINGS",
}

func (v *ConditionTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConditionTypeEnum(value)
	for _, existing := range allowedConditionTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = CONDITIONTYPEENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewConditionTypeEnumFromValue returns a pointer to a valid ConditionTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionTypeEnumFromValue(v string) (*ConditionTypeEnum, error) {
	ev := ConditionTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := CONDITIONTYPEENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConditionTypeEnum) IsValid() bool {
	for _, existing := range allowedConditionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConditionTypeEnum value
func (v ConditionTypeEnum) Ptr() *ConditionTypeEnum {
	return &v
}

type NullableConditionTypeEnum struct {
	value *ConditionTypeEnum
	isSet bool
}

func (v NullableConditionTypeEnum) Get() *ConditionTypeEnum {
	return v.value
}

func (v *NullableConditionTypeEnum) Set(val *ConditionTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionTypeEnum(val *ConditionTypeEnum) *NullableConditionTypeEnum {
	return &NullableConditionTypeEnum{value: val, isSet: true}
}

func (v NullableConditionTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

