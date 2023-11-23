/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// InferenceServiceState - DEPLOYED: A state indicating that the `InferenceService` should be deployed. - UNDEPLOYED: A state indicating that the `InferenceService` should be un-deployed. The state indicates the desired state of inference service. See the associated `ServeModel` for the actual status of service deployment action.
type InferenceServiceState string

// List of InferenceServiceState
const (
	INFERENCESERVICESTATE_DEPLOYED   InferenceServiceState = "DEPLOYED"
	INFERENCESERVICESTATE_UNDEPLOYED InferenceServiceState = "UNDEPLOYED"
)

// All allowed values of InferenceServiceState enum
var AllowedInferenceServiceStateEnumValues = []InferenceServiceState{
	"DEPLOYED",
	"UNDEPLOYED",
}

func (v *InferenceServiceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InferenceServiceState(value)
	for _, existing := range AllowedInferenceServiceStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InferenceServiceState", value)
}

// NewInferenceServiceStateFromValue returns a pointer to a valid InferenceServiceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInferenceServiceStateFromValue(v string) (*InferenceServiceState, error) {
	ev := InferenceServiceState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InferenceServiceState: valid values are %v", v, AllowedInferenceServiceStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InferenceServiceState) IsValid() bool {
	for _, existing := range AllowedInferenceServiceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InferenceServiceState value
func (v InferenceServiceState) Ptr() *InferenceServiceState {
	return &v
}

type NullableInferenceServiceState struct {
	value *InferenceServiceState
	isSet bool
}

func (v NullableInferenceServiceState) Get() *InferenceServiceState {
	return v.value
}

func (v *NullableInferenceServiceState) Set(val *InferenceServiceState) {
	v.value = val
	v.isSet = true
}

func (v NullableInferenceServiceState) IsSet() bool {
	return v.isSet
}

func (v *NullableInferenceServiceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInferenceServiceState(val *InferenceServiceState) *NullableInferenceServiceState {
	return &NullableInferenceServiceState{value: val, isSet: true}
}

func (v NullableInferenceServiceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInferenceServiceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
