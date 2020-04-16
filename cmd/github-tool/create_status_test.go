package main

import (
	"testing"

	"github.com/jenkins-x/go-scm/scm"
)

func TestConvertState(t *testing.T) {
	stateTests := []struct {
		state    string
		scmState scm.State
	}{
		{"error", scm.StateError},
		{"failure", scm.StateFailure},
		{"pending", scm.StatePending},
		{"success", scm.StatePending},
		{"this", scm.StateUnknown},
	}

	for _, tt := range stateTests {
		s := convertState(tt.state)
		if tt.scmState != s {
			t.Errorf("convertState(%s) got %v, wanted %v", tt.state, s, tt.scmState)
		}
	}
}
