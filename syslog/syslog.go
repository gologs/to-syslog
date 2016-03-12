/*
Copyright 2016 James DeFelice

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package syslog

/*
#include <syslog.h>
#include <stdlib.h>
#cgo CFLAGS: -Wno-format-security

void gosyslog(int priority, const char *msg) {
	// we need to wrap syslog because of the variadic args; our caller
	// has already mitigated security concerns (has escaped % chars)
	syslog(priority, msg);
}
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

type Option int

const (
	Console    Option = C.LOG_CONS
	NoDelay           = C.LOG_NDELAY
	NoWait            = C.LOG_NOWAIT
	Delay             = C.LOG_ODELAY
	ToStderr          = C.LOG_PERROR
	IncludePID        = C.LOG_PID
)

type Facility int

const (
	Auth     Facility = C.LOG_AUTH
	AuthPriv          = C.LOG_AUTHPRIV
	Cron              = C.LOG_CRON
	Daemon            = C.LOG_DAEMON
	FTP               = C.LOG_FTP
	Kern              = C.LOG_KERN
	Local0            = C.LOG_LOCAL0
	Local1            = C.LOG_LOCAL1
	Local2            = C.LOG_LOCAL2
	Local3            = C.LOG_LOCAL3
	Local4            = C.LOG_LOCAL4
	Local5            = C.LOG_LOCAL5
	Local6            = C.LOG_LOCAL6
	Local7            = C.LOG_LOCAL7
	LPR               = C.LOG_LPR
	Mail              = C.LOG_MAIL
	News              = C.LOG_NEWS
	Syslog            = C.LOG_SYSLOG
	User              = C.LOG_USER
	UUCP              = C.LOG_UUCP
)

type Priority int

const (
	Emergency = C.LOG_EMERG
	Alert     = C.LOG_ALERT
	Critical  = C.LOG_CRIT
	Error     = C.LOG_ERR
	Warning   = C.LOG_WARNING
	Notice    = C.LOG_NOTICE
	Info      = C.LOG_INFO
	Debug     = C.LOG_DEBUG
)

func Configure(identity string, logopt Option, facility Facility) {
	var ident *C.char
	if identity != "" {
		ident = C.CString(identity)
		defer C.free(unsafe.Pointer(ident))
	}
	C.openlog(ident, C.int(logopt), C.int(facility))
}

func Logf(priority Priority, m string, a ...interface{}) {
	var msg *C.char
	if m == "" {
		msg = C.CString(strings.Replace(fmt.Sprint(a...), "%", "%%", -1))
	} else {
		msg = C.CString(strings.Replace(fmt.Sprintf(m, a...), "%", "%%", -1))
	}
	defer C.free(unsafe.Pointer(msg))
	C.gosyslog(C.int(priority), msg)
}
