package reverseonlyletters

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "Test Case 1",
			s:    "ab-cd",
			want: "dc-ba",
		}, {
			name: "Test Case 2",
			s:    "a-bC-dEf-ghIj",
			want: "j-Ih-gfE-dCba",
		}, {
			name: "Test Case 3",
			s:    "Test1ng-Leet=code-Q!",
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseOnlyLetters(tt.s)
			if got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
