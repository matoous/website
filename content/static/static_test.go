// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package static

import (
	"bytes"
	"io/ioutil"
	"testing"

	"golang.org/x/website/content/static/internal/gen"
)

func TestStaticIsUpToDate(t *testing.T) {
	oldBuf, err := ioutil.ReadFile("static.go")
	if err != nil {
		t.Errorf("error while reading static.go: %v\n", err)
	}

	newBuf, err := gen.Generate()
	if err != nil {
		t.Errorf("error while generating static.go: %v\n", err)
	}

	if bytes.Compare(oldBuf, newBuf) != 0 {
		t.Error(`static.go is stale.  Run:
  $ go generate golang.org/x/website/content/static
  $ git diff
to see the differences.`)

	}
}
