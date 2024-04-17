package cdir

import "testing"

func TestGetIPs(t *testing.T) {
	cidr := "192.168.0.0/30"
	expected := []string{"192.168.0.1", "192.168.0.2"}

	ips, err := GetIPs(cidr)
	if err != nil {
		t.Errorf("GetIPs(%s) returned error: %v", cidr, err)
	}

	if len(ips) != len(expected) {
		t.Errorf("GetIPs(%s) returned %d IPs, expected %d", cidr, len(ips), len(expected))
	}

	for i := range expected {
		if ips[i] != expected[i] {
			t.Errorf("GetIPs(%s) result %s, expected %s", cidr, ips[i], expected[i])
		}
	}
}
