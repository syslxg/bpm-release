// Copyright (C) 2018-Present CloudFoundry.org Foundation, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package exitstatus

import (
	"errors"
	"testing"
)

func TestMessage(t *testing.T) {
	msg := "disaster"
	err := &Error{
		Status: 34,
		Err:    errors.New(msg),
	}

	if got, want := err.Error(), "disaster (exit status 34)"; got != want {
		t.Errorf("err.Error(..) = %q; want: %q", got, want)
	}
}

func TestStatus(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want int
	}{
		{
			name: "nil",
			err:  nil,
			want: 0,
		},
		{
			name: "status error",
			err:  &Error{Status: 41, Err: errors.New("oops")},
			want: 41,
		},
		{
			name: "other error",
			err:  errors.New("other"),
			want: 1,
		},
	}

	for _, tc := range cases {
		tc := tc // capture variable
		t.Run(tc.name, func(t *testing.T) {
			if got, want := FromError(tc.err), tc.want; got != want {
				t.Errorf("FromError(..) = %d; want: %d", got, want)
			}
		})
	}
}
