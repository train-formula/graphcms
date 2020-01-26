package util

import "strings"

// Run strings.TrimSpace if input is not nil, or return nil.
func TrimSpaceNotNil(s *string) *string {
	if s != nil {
		t := strings.TrimSpace(*s)
		return &t
	}

	return nil
}
