package findtheindexofthefirstoccurrenceinastring

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "example 1",
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			name:     "example 2",
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			name:     "example 3",
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			name:     "xample 4",
			haystack: "abcde",
			needle:   "ab",
			want:     0,
		},
		{
			name:     "xample 5",
			haystack: "a",
			needle:   "a",
			want:     0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			if got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
