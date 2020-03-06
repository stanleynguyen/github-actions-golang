// Copyright (c) 2019, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package actions

import (
	"fmt"
	"runtime"
	nonce "github.com/LarryBattle/nonce-golang"
	gpv "github.com/go-playground/validator/v10"
)

func Demo() {
	gpv.New()
	fmt.Printf("Nonce token: %s\n", nonce.NewToken())
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
}
