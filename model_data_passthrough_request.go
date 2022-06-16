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

// DataPassthroughRequest # The DataPassthrough Object ### Description The `DataPassthrough` object is used to send information to an otherwise-unsupported third-party endpoint.  ### Usage Example Create a `DataPassthrough` to get team hierarchies from your Rippling integration.
type DataPassthroughRequest struct {
	Method MethodEnum `json:"method"`
	Path string `json:"path"`
	BaseUrlOverride NullableString `json:"base_url_override,omitempty"`
	Data NullableString `json:"data,omitempty"`
	// Pass an array of `MultipartFormField` objects in here instead of using the `data` param if `request_format` is set to `MULTIPART`.
	MultipartFormData []MultipartFormFieldRequest `json:"multipart_form_data,omitempty"`
	// The headers to use for the request (Merge will handle the account's authorization headers). `Content-Type` header is required for passthrough. Choose content type corresponding to expected format of receiving server.
	Headers map[string]interface{} `json:"headers,omitempty"`
	RequestFormat NullableRequestFormatEnum `json:"request_format,omitempty"`
	NormalizeResponse *bool `json:"normalize_response,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewDataPassthroughRequest instantiates a new DataPassthroughRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPassthroughRequest(method MethodEnum, path string) *DataPassthroughRequest {
	this := DataPassthroughRequest{}
	this.Method = method
	this.Path = path
	return &this
}

// NewDataPassthroughRequestWithDefaults instantiates a new DataPassthroughRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPassthroughRequestWithDefaults() *DataPassthroughRequest {
	this := DataPassthroughRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *DataPassthroughRequest) GetMethod() MethodEnum {
	if o == nil {
		var ret MethodEnum
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *DataPassthroughRequest) GetMethodOk() (*MethodEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *DataPassthroughRequest) SetMethod(v MethodEnum) {
	o.Method = v
}

// GetPath returns the Path field value
func (o *DataPassthroughRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DataPassthroughRequest) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DataPassthroughRequest) SetPath(v string) {
	o.Path = v
}

// GetBaseUrlOverride returns the BaseUrlOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPassthroughRequest) GetBaseUrlOverride() string {
	if o == nil || o.BaseUrlOverride.Get() == nil {
		var ret string
		return ret
	}
	return *o.BaseUrlOverride.Get()
}

// GetBaseUrlOverrideOk returns a tuple with the BaseUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPassthroughRequest) GetBaseUrlOverrideOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BaseUrlOverride.Get(), o.BaseUrlOverride.IsSet()
}

// HasBaseUrlOverride returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasBaseUrlOverride() bool {
	if o != nil && o.BaseUrlOverride.IsSet() {
		return true
	}

	return false
}

// SetBaseUrlOverride gets a reference to the given NullableString and assigns it to the BaseUrlOverride field.
func (o *DataPassthroughRequest) SetBaseUrlOverride(v string) {
	o.BaseUrlOverride.Set(&v)
}
// SetBaseUrlOverrideNil sets the value for BaseUrlOverride to be an explicit nil
func (o *DataPassthroughRequest) SetBaseUrlOverrideNil() {
	o.BaseUrlOverride.Set(nil)
}

// UnsetBaseUrlOverride ensures that no value is present for BaseUrlOverride, not even an explicit nil
func (o *DataPassthroughRequest) UnsetBaseUrlOverride() {
	o.BaseUrlOverride.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPassthroughRequest) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPassthroughRequest) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *DataPassthroughRequest) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *DataPassthroughRequest) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *DataPassthroughRequest) UnsetData() {
	o.Data.Unset()
}

// GetMultipartFormData returns the MultipartFormData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPassthroughRequest) GetMultipartFormData() []MultipartFormFieldRequest {
	if o == nil  {
		var ret []MultipartFormFieldRequest
		return ret
	}
	return o.MultipartFormData
}

// GetMultipartFormDataOk returns a tuple with the MultipartFormData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPassthroughRequest) GetMultipartFormDataOk() (*[]MultipartFormFieldRequest, bool) {
	if o == nil || o.MultipartFormData == nil {
		return nil, false
	}
	return &o.MultipartFormData, true
}

// HasMultipartFormData returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasMultipartFormData() bool {
	if o != nil && o.MultipartFormData != nil {
		return true
	}

	return false
}

// SetMultipartFormData gets a reference to the given []MultipartFormFieldRequest and assigns it to the MultipartFormData field.
func (o *DataPassthroughRequest) SetMultipartFormData(v []MultipartFormFieldRequest) {
	o.MultipartFormData = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPassthroughRequest) GetHeaders() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPassthroughRequest) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *DataPassthroughRequest) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetRequestFormat returns the RequestFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataPassthroughRequest) GetRequestFormat() RequestFormatEnum {
	if o == nil || o.RequestFormat.Get() == nil {
		var ret RequestFormatEnum
		return ret
	}
	return *o.RequestFormat.Get()
}

// GetRequestFormatOk returns a tuple with the RequestFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataPassthroughRequest) GetRequestFormatOk() (*RequestFormatEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequestFormat.Get(), o.RequestFormat.IsSet()
}

// HasRequestFormat returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasRequestFormat() bool {
	if o != nil && o.RequestFormat.IsSet() {
		return true
	}

	return false
}

// SetRequestFormat gets a reference to the given NullableRequestFormatEnum and assigns it to the RequestFormat field.
func (o *DataPassthroughRequest) SetRequestFormat(v RequestFormatEnum) {
	o.RequestFormat.Set(&v)
}
// SetRequestFormatNil sets the value for RequestFormat to be an explicit nil
func (o *DataPassthroughRequest) SetRequestFormatNil() {
	o.RequestFormat.Set(nil)
}

// UnsetRequestFormat ensures that no value is present for RequestFormat, not even an explicit nil
func (o *DataPassthroughRequest) UnsetRequestFormat() {
	o.RequestFormat.Unset()
}

// GetNormalizeResponse returns the NormalizeResponse field value if set, zero value otherwise.
func (o *DataPassthroughRequest) GetNormalizeResponse() bool {
	if o == nil || o.NormalizeResponse == nil {
		var ret bool
		return ret
	}
	return *o.NormalizeResponse
}

// GetNormalizeResponseOk returns a tuple with the NormalizeResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPassthroughRequest) GetNormalizeResponseOk() (*bool, bool) {
	if o == nil || o.NormalizeResponse == nil {
		return nil, false
	}
	return o.NormalizeResponse, true
}

// HasNormalizeResponse returns a boolean if a field has been set.
func (o *DataPassthroughRequest) HasNormalizeResponse() bool {
	if o != nil && o.NormalizeResponse != nil {
		return true
	}

	return false
}

// SetNormalizeResponse gets a reference to the given bool and assigns it to the NormalizeResponse field.
func (o *DataPassthroughRequest) SetNormalizeResponse(v bool) {
	o.NormalizeResponse = &v
}

func (o DataPassthroughRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.BaseUrlOverride.IsSet() {
		toSerialize["base_url_override"] = o.BaseUrlOverride.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.MultipartFormData != nil {
		toSerialize["multipart_form_data"] = o.MultipartFormData
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.RequestFormat.IsSet() {
		toSerialize["request_format"] = o.RequestFormat.Get()
	}
	if o.NormalizeResponse != nil {
		toSerialize["normalize_response"] = o.NormalizeResponse
	}
	return json.Marshal(toSerialize)
}

func (v *DataPassthroughRequest) UnmarshalJSON(src []byte) error {
    type DataPassthroughRequestUnmarshalTarget DataPassthroughRequest

	var intermediateResult DataPassthroughRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = DataPassthroughRequest(intermediateResult)
	return nil
}
type NullableDataPassthroughRequest struct {
	value *DataPassthroughRequest
	isSet bool
}

func (v NullableDataPassthroughRequest) Get() *DataPassthroughRequest {
	return v.value
}

func (v *NullableDataPassthroughRequest) Set(val *DataPassthroughRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPassthroughRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPassthroughRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPassthroughRequest(val *DataPassthroughRequest) *NullableDataPassthroughRequest {
	return &NullableDataPassthroughRequest{value: val, isSet: true}
}

func (v NullableDataPassthroughRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPassthroughRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


