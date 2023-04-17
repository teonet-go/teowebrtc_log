// Copyright 2021-2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Webrtc log contains log definition used in Teonet webrts packages
package teowebrtc_log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Teolog
type teolog struct {
	logFlags int
	packages packagesMap
}
type packagesMap map[string]*packageData
type packageData struct {
	show   bool
	prefix string
	log    *log.Logger
}

// Teonet webrtc modules name
const (
	Package_main                    = "main"
	Package_teowebrtc_log           = "teowebrtc_log"
	Package_teowebrtc_server        = "teowebrtc_server"
	Package_teowebrtc_client        = "teowebrtc_client"
	Package_teowebrtc_signal        = "teowebrtc_signal"
	Package_teowebrtc_signal_client = "teowebrtc_signal_client"
)

var logt *log.Logger

// Initialize Teonet Logger and set default parameters
var teologPtr *teolog = func() (t *teolog) {
	t = &teolog{
		// Logger flags
		logFlags: log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix,
		// Application packages
		packages: packagesMap{
			Package_main:                    &packageData{show: true, prefix: "[ASUZS ] "},
			Package_teowebrtc_log:           &packageData{show: true, prefix: "[TEOLOG] "},
			Package_teowebrtc_server:        &packageData{show: true, prefix: "[WEBRTC] "},
			Package_teowebrtc_client:        &packageData{show: true, prefix: "[WEBCLI] "},
			Package_teowebrtc_signal:        &packageData{show: true, prefix: "[SIGNAL] "},
			Package_teowebrtc_signal_client: &packageData{show: true, prefix: "[SIGCLI] "},
		},
	}
	for packageName, pac := range t.packages {
		pac.log, _ = t.getLog(packageName)
	}
	logt = t.packages[Package_teowebrtc_log].log
	logt.Println("teolog created")
	return
}()

// getLog returns teowebrtc logger depend of package name
func (teolog *teolog) getLog(packageName string) (logt *log.Logger, err error) {

	pac, ok := teolog.packages[packageName]
	if !ok {
		err = fmt.Errorf("wrong package %s", packageName)
	}

	// Log output
	var out io.Writer = os.Stdout
	if !pac.show {
		out = io.Discard
	}

	// Create new loger
	logt = log.New(out, pac.prefix, teolog.logFlags)
	pac.log = logt

	return
}

// GetLog returns teowebrtc logger depend of package name
func GetLog(packageName string) (logger *log.Logger) {

	logt.Printf("get log %s", packageName)

	pac, ok := teologPtr.packages[packageName]
	if !ok {
		logt.Printf("error: wrong package %s", packageName)
		return
	}
	logger = pac.log

	return
}

// SetVisibility sets the visibility of package log
func SetVisibility(packageName string, show bool) (err error) {

	pac, ok := teologPtr.packages[packageName]
	if !ok {
		err = fmt.Errorf("wrong package %s", packageName)
		return
	}

	pac.show = show

	var out io.Writer = os.Stdout
	if !pac.show {
		out = io.Discard
	}

	pac.log.SetOutput(out)

	return
}
