/*
Copyright © 2020 Fernando Silva <audacius@pm.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package logging wraps and setup the application logger
package logging

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/viper"
)

func init() {
	log.SetHandler(cli.Default)

	viper.SetDefault("LOG_LEVEL", "info")
}

// New returns a new named logger
func New(name string) *log.Entry {
	// Can't be in `init` as it runs before configuration (viper) reads
	log.SetLevel(log.MustParseLevel(viper.GetString("LOG_LEVEL")))

	return log.WithFields(log.Fields{
		"name": name,
	})
}
