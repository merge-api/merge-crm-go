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

// RemoteResponse # The RemoteResponse Object ### Description The `RemoteResponse` object is used to represent information returned from a third-party endpoint.  ### Usage Example View the `RemoteResponse` returned from your `DataPassthrough`.
type RemoteResponse struct {
	Method string `json:"method"`
	Path string `json:"path"`
	Status int32 `json:"status"`
	Response map[string]interface{} `json:"response"`
	ResponseHeaders *map[string]interface{} `json:"response_headers,omitempty"`
	Headers *map[string]interface{} `json:"headers,omitempty"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewRemoteResponse instantiates a new RemoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteResponse(method string, path string, status int32, response map[string]interface{}) *RemoteResponse {
	this := RemoteResponse{}
	this.Method = method
	this.Path = path
	this.Status = status
	this.Response = response
	return &this
}

// NewRemoteResponseWithDefaults instantiates a new RemoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteResponseWithDefaults() *RemoteResponse {
	this := RemoteResponse{}
	return &this
}

// GetMethod returns the Method field value
func (o *RemoteResponse) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *RemoteResponse) SetMethod(v string) {
	o.Method = v
}

// GetPath returns the Path field value
func (o *RemoteResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RemoteResponse) SetPath(v string) {
	o.Path = v
}

// GetStatus returns the Status field value
func (o *RemoteResponse) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetStatusOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoteResponse) SetStatus(v int32) {
	o.Status = v
}

// GetResponse returns the Response field value
func (o *RemoteResponse) GetResponse() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetResponseOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *RemoteResponse) SetResponse(v map[string]interface{}) {
	o.Response = v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *RemoteResponse) GetResponseHeaders() map[string]interface{} {
	if o == nil || o.ResponseHeaders == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetResponseHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResponseHeaders == nil {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *RemoteResponse) HasResponseHeaders() bool {
	if o != nil && o.ResponseHeaders != nil {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given map[string]interface{} and assigns it to the ResponseHeaders field.
func (o *RemoteResponse) SetResponseHeaders(v map[string]interface{}) {
	o.ResponseHeaders = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *RemoteResponse) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteResponse) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RemoteResponse) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *RemoteResponse) SetHeaders(v map[string]interface{}) {
	o.Headers = &v
}

func (o RemoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["response"] = o.Response
	}
	if o.ResponseHeaders != nil {
		toSerialize["response_headers"] = o.ResponseHeaders
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	return json.Marshal(toSerialize)
}

func (v *RemoteResponse) UnmarshalJSON(src []byte) error {
    type RemoteResponseUnmarshalTarget RemoteResponse

	var intermediateResult RemoteResponseUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = RemoteResponse(intermediateResult)
	return nil
}
type NullableRemoteResponse struct {
	value *RemoteResponse
	isSet bool
}

func (v NullableRemoteResponse) Get() *RemoteResponse {
	return v.value
}

func (v *NullableRemoteResponse) Set(val *RemoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteResponse(val *RemoteResponse) *NullableRemoteResponse {
	return &NullableRemoteResponse{value: val, isSet: true}
}

func (v NullableRemoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


