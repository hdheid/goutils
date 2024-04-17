package goutils

import "testing"

func TestGetIP(t *testing.T) {
	ip := "127.0.0.1"
	t.Logf(GetIP(ip))
}
