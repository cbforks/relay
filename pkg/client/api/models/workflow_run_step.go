// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowRunStep workflow run step
// swagger:model WorkflowRunStep
type WorkflowRunStep struct {
	WorkflowStep

	// Time at which the step execution ended
	// Format: date-time
	EndedAt *strfmt.DateTime `json:"ended_at,omitempty"`

	// Time at which step execution started
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	// Workflow run step status
	// Required: true
	// Enum: [success failure in-progress pending]
	Status *string `json:"status"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowRunStep) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowStep
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowStep = aO0

	// AO1
	var dataAO1 struct {
		EndedAt *strfmt.DateTime `json:"ended_at,omitempty"`

		StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

		Status *string `json:"status"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EndedAt = dataAO1.EndedAt

	m.StartedAt = dataAO1.StartedAt

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowRunStep) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowStep)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EndedAt *strfmt.DateTime `json:"ended_at,omitempty"`

		StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

		Status *string `json:"status"`
	}

	dataAO1.EndedAt = m.EndedAt

	dataAO1.StartedAt = m.StartedAt

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow run step
func (m *WorkflowRunStep) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowStep
	if err := m.WorkflowStep.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowRunStep) validateEndedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EndedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ended_at", "body", "date-time", m.EndedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowRunStep) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var workflowRunStepTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success","failure","in-progress","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowRunStepTypeStatusPropEnum = append(workflowRunStepTypeStatusPropEnum, v)
	}
}

// property enum
func (m *WorkflowRunStep) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowRunStepTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowRunStep) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRunStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRunStep) UnmarshalBinary(b []byte) error {
	var res WorkflowRunStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
