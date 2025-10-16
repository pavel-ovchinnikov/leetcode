package defanginganipaddress

import "testing"

func Test_defangIPaddr(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		address string
		want    string
	}{
		{
			name:    "Example 1",
			address: "1.1.1.1",
			want:    "1[.]1[.]1[.]1",
		}, {
			name:    "Example 2",
			address: "255.100.50.0",
			want:    "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defangIPaddr(tt.address)

			if got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
