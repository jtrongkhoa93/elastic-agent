// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package capabilities

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/elastic/elastic-agent/pkg/core/logger"
)

// Capability provides a way of applying predefined filter to object.
// It's up to capability to determine if capability is applicable on object.
type Capability interface {
	// Apply applies capabilities on input and returns true if input should be completely blocked
	// otherwise, false and updated input is returned
	Apply(interface{}) (interface{}, error)
}

var (
	// ErrBlocked is returned when capability is blocking.
	ErrBlocked = errors.New("capability blocked")
)

type capabilitiesManager struct {
	caps []Capability
}

type capabilityFactory func(*logger.Logger, *ruleDefinitions) (Capability, error)

// Load loads capabilities files and prepares manager.
func Load(capsFile string, log *logger.Logger) (Capability, error) {
	handlers := []capabilityFactory{
		newInputsCapability,
		newOutputsCapability,
		newUpgradesCapability,
	}

	cm := &capabilitiesManager{
		caps: make([]Capability, 0),
	}

	// load capabilities from file
	fd, err := os.Open(capsFile)
	if err != nil && !os.IsNotExist(err) {
		return cm, err
	}

	if os.IsNotExist(err) {
		log.Infof("Capabilities file not found in %s", capsFile)
		return cm, nil
	}
	defer fd.Close()

	definitions := &ruleDefinitions{Capabilities: make([]ruler, 0)}
	dec := yaml.NewDecoder(fd)
	if err := dec.Decode(&definitions); err != nil {
		return cm, err
	}

	// make list of handlers out of capabilities definition
	for _, h := range handlers {
		cap, err := h(log, definitions)
		if err != nil {
			return nil, err
		}

		if cap == nil {
			continue
		}

		cm.caps = append(cm.caps, cap)
	}

	return cm, nil
}

func (mgr *capabilitiesManager) Apply(in interface{}) (interface{}, error) {
	var err error
	for _, cap := range mgr.caps {
		in, err = cap.Apply(in)
		if err != nil {
			return in, err
		}
	}

	return in, nil
}
