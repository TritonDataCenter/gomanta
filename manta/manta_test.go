//
// gomanta - Go library to interact with Joyent Manta
//
//
// Copyright (c) 2013 Joyent Inc.
//
// Written by Daniele Stroppa <daniele.stroppa@joyent.com>
//

package manta_test

import (
	"flag"
	"github.com/joyent/gocommon/jpc"
	gc "launchpad.net/gocheck"
	"testing"
)

var live = flag.Bool("live", false, "Include live Manta tests")
var keyName = flag.String("key.name", "", "Specify the full path to the private key, defaults to ~/.ssh/id_rsa")

func Test(t *testing.T) {
	if *live {
		creds, err := jpc.CompleteCredentialsFromEnv(*keyName)
		if err != nil {
			t.Fatalf("Error setting up test suite: %s", err.Error())
		}
		registerMantaTests(creds)
	}
	registerLocalTests(*keyName)
	gc.TestingT(t)
}
