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

package tosyslog

import (
	"github.com/gologs/log/context"
	"github.com/gologs/log/levels"
	"github.com/gologs/log/logger"
	"github.com/gologs/to-syslog/syslog"
)

var (
	levelMap = map[levels.Level]syslog.Priority{
		levels.Debug: syslog.Debug,
		levels.Info:  syslog.Info,
		levels.Warn:  syslog.Warning,
		levels.Error: syslog.Error,
		levels.Fatal: syslog.Critical,
		levels.Panic: syslog.Critical,
	}
	syslogger = logger.Func(func(ctx context.Context, m string, a ...interface{}) {
		l, ok := levels.FromContext(ctx)
		if ok {
			priority, ok := levelMap[l]
			if ok {
				syslog.Logf(priority, m, a...)
				return
			}
		}
		syslog.Logf(syslog.Info, m, a...)
	})
)

func Logger() logger.Logger {
	return syslogger
}
