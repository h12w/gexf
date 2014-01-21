// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gexf

import (
	"encoding/gob"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"strings"
	"time"
)

func p(v ...interface{}) {
	fmt.Println(v...)
}

var printer = spew.ConfigState{Indent: "    "}

func pp(v ...interface{}) {
	for _, item := range v {
		printer.Dump(item)
	}
}

func urlToFileName(s string) string {
	return strings.Replace(s, "/", "_", -1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkErrorf(cond bool, format string, a ...interface{}) {
	if cond {
		panic(fmt.Errorf(format, a...))
	}
}

func now() *time.Time {
	t := time.Now()
	return &t
}

func FromGobFile(file string, e interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(e)
	if err != nil {
		return err
	}
	return nil
}

func ToGobFile(file string, e interface{}) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = gob.NewEncoder(f).Encode(e)
	if err != nil {
		return err
	}
	return nil
}
