// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package plugins

import (
	"path"
	"plugin"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	for _, path := range []string{"foo", "bar"} {
		m := New(path)
		assert.Equal(t, path, m.Path)
	}
}

func TestManagerOpen(t *testing.T) {
	dir := "testdata"
	m := New(dir)

	for _, name := range []string{"foo.so", "bar.so"} {
		expectedP, expectedErr := plugin.Open(path.Join(dir, name))
		p, err := m.Open(name)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, expectedP, p)
	}
}

func TestManagerLookup(t *testing.T) {
	dir := "testdata"
	m := New(dir)

	_, err := m.Lookup("nil", "Foo")
	assert.NotNil(t, err)

	_, err = m.Lookup("foo.so", "Bar")
	assert.NotNil(t, err)

	sym, err := m.Lookup("foo.so", "Foo")
	assert.Nil(t, err)
	f := sym.(func() string)
	assert.Equal(t, "Bar", f())
}
