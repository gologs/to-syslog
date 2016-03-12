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

package main

import (
	"github.com/gologs/log"
	"github.com/gologs/log/config"
	"github.com/gologs/to-syslog"
	"github.com/gologs/to-syslog/syslog"
)

func main() {
	syslog.Configure("", syslog.ToStderr, syslog.User)
	config.Default, _ = config.DefaultConfig.With(
		config.Logger(tosyslog.Logger()),
	)

	log.Logf("hello %% world")
	log.Debugf("this is debug")
	log.Infof("this is info")
	log.Warnf("this is warn")
	log.Errorf("this is error")
	log.Fatalf("this is fatal")
}
