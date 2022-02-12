package dhcp

import (
	"path"
	"testing"

	"github.com/hpcng/warewulf/internal/pkg/buildconfig"
)

func TestDhcpTemplateFile(t *testing.T) {
	tests := []struct{
		parameter string
		expected string
	}{
		{"", path.Join(buildconfig.SYSCONFDIR(), "warewulf/dhcp/default-dhcpd.conf")},
		{"default", path.Join(buildconfig.SYSCONFDIR(),  "warewulf/dhcp/default-dhcpd.conf")},
		{"static", path.Join(buildconfig.SYSCONFDIR(), "warewulf/dhcp/static-dhcpd.conf")},
		{"/test/absolute/path.conf", "/test/absolute/path.conf"},
	}
	for _, tt := range tests {
		actual := dhcpTemplateFile(tt.parameter)
		if actual != tt.expected {
			t.Errorf("dhcpTemplateFile(%v) expected: %v, actual: %v",
				tt.parameter, tt.expected, actual)
		}
	}
}
