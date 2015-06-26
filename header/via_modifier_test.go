// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package header

import (
	"net/http"
	"testing"

	"github.com/google/martian"
)

func TestViaModifier(t *testing.T) {
	m := NewViaModifier("1.1 martian")
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(): got %v, want no error", err)
	}

	if err := m.ModifyRequest(martian.NewContext(), req); err != nil {
		t.Fatalf("ModifyRequest(): got %v, want no error", err)
	}
	if got, want := req.Header.Get("Via"), "1.1 martian"; got != want {
		t.Errorf("req.Header.Get(%q): got %q, want %q", "Via", got, want)
	}

	req.Header.Set("Via", "1.0 alpha")
	if err := m.ModifyRequest(martian.NewContext(), req); err != nil {
		t.Fatalf("ModifyRequest(): got %v, want no error", err)
	}
	if got, want := req.Header.Get("Via"), "1.0 alpha, 1.1 martian"; got != want {
		t.Errorf("req.Header.Get(%q): got %q, want %q", "Via", got, want)
	}
}
