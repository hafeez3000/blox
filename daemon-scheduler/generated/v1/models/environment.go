// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Environment A representation of environment managed by scheduler via deployments
// swagger:model Environment
type Environment struct {

	// The token used to verify that the deployment is being kicked off on the correct version of the environment
	DeploymentToken string `json:"deploymentToken,omitempty"`

	// health
	// Required: true
	Health HealthStatus `json:"health"`

	// instance group
	// Required: true
	InstanceGroup *InstanceGroup `json:"instanceGroup"`

	// Name of the environment
	// Required: true
	Name *string `json:"name"`

	// TaskDefinition used to start tasks under this environment
	TaskDefinition string `json:"taskDefinition,omitempty"`
}

// Validate validates this environment
func (m *Environment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Environment) validateHealth(formats strfmt.Registry) error {

	if err := m.Health.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Environment) validateInstanceGroup(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroup", "body", m.InstanceGroup); err != nil {
		return err
	}

	if m.InstanceGroup != nil {

		if err := m.InstanceGroup.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Environment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
