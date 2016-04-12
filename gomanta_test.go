//
// gomanta - Go library to interact with Joyent Manta
//
//
// Copyright (c) 2016 Joyent Inc.
//
// Written by Daniele Stroppa <daniele.stroppa@joyent.com>
//

package gomanta

import (
	gc "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	gc.TestingT(t)
}

type GoMantaTestSuite struct {
}

var _ = gc.Suite(&GoMantaTestSuite{})
