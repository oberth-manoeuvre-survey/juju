// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package muxhttpserver_test

import (
	"testing"

	"github.com/juju/juju/pki"
	pki_test "github.com/juju/juju/pki/test"
	gc "gopkg.in/check.v1"
)

func TestSuite(t *testing.T) { gc.TestingT(t) }

func init() {
	// Use full strength key profile
	pki.DefaultKeyProfile = pki_test.OriginalDefaultKeyProfile
}
