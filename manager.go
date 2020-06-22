// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package plugins

import (
	"path"
	"plugin"
)

// Manager is a manager that manage plugins.
type Manager struct {
	// The location of plugins.
	Path string
}

// New returns a plugins manager with the given path.
func New(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

// Open opens a Go plugin.
func (pm *Manager) Open(name string) (*plugin.Plugin, error) {
	return plugin.Open(path.Join(pm.Path, name))
}

// Lookup searchs for a symbol named symName in a plugin named pluginName.
func (pm *Manager) Lookup(pluginName, symName string) (plugin.Symbol, error) {
	p, err := pm.Open(pluginName)
	if err != nil {
		return nil, err
	}
	return p.Lookup(symName)
}
