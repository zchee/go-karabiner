// Copyright 2017 The go-karabiner Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	karabiner "github.com/zchee/go-karabiner"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need json file")
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	v := new(karabiner.Karabiner)
	if err := json.Unmarshal(data, v); err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
