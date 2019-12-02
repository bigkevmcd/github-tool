package main

import "testing"

func TestSplitFullName(t *testing.T) {
	nameTests := []struct {
		fullname string
		owner    string
		repo     string
		fail     bool
	}{
		{"testing/testing", "testing", "testing", false},
		{"testing", "", "", true},
	}

	for _, tt := range nameTests {
		owner, repo, err := splitFullname(tt.fullname)
		if err != nil && !tt.fail {
			t.Errorf("splitFullname(%s) failed %v", tt.fullname, err)
			continue
		}
		if tt.owner != owner {
			t.Errorf("splitFullname(%s) owner got %s, wanted %s\n", tt.fullname, owner, tt.owner)
		}
		if tt.repo != repo {
			t.Errorf("splitFullname(%s) repo got %s, wanted %s\n", tt.fullname, repo, tt.repo)
		}
	}
}
