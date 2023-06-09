// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build darwin
// +build darwin

package paths

const (
	// BinaryName is the name of the installed binary.
	BinaryName = "elastic-agent"

	// InstallPath is the installation path using for install command.
	InstallPath = "/Library/Elastic/Agent"

	// ControlSocketPath is the control socket path used when installed.
	ControlSocketPath = "unix:///var/run/elastic-agent.sock"

	// ShipperSocketPipePattern is the socket path used when installed for a shipper pipe.
	ShipperSocketPipePattern = "unix:///var/run/elastic-agent-%s-pipe.sock"

	// ServiceName is the service name when installed.
	ServiceName = "co.elastic.elastic-agent"

	// ShellWrapperPath is the path to the installed shell wrapper.
	ShellWrapperPath = "/usr/local/bin/elastic-agent"

	// ShellWrapper is the wrapper that is installed.
	ShellWrapper = `#!/bin/sh
exec /Library/Elastic/Agent/elastic-agent $@
`

	// defaultAgentVaultName is keychain item name for mac
	defaultAgentVaultName = "co.elastic.elastic-agent"
)

// ArePathsEqual determines whether paths are equal taking case sensitivity of os into account.
func ArePathsEqual(expected, actual string) bool {
	return expected == actual
}

// AgentVaultPath is keychain name on Mac OS
func AgentVaultPath() string {
	return defaultAgentVaultName
}
