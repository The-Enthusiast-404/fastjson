package fastjson

import (
	"encoding/json"
	"testing"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
		wantErr bool
	} {
		{
			name: "simple string",
			input: `"hello"`,
			expected:"hello",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("ParseString() = %v, want %v", got, tt.expected)
			}
		})
	}


}

func BenchmarkParseString_FastJSON(b *testing.B) {
	input := `"hello world"`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			_, _ = ParseString(input)
	}
}

func BenchmarkParseString_StdLib(b *testing.B) {
	input := `"hello world"`
	var result string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			_ = json.Unmarshal([]byte(input), &result)
	}
}

// Let's also add benchmarks for different string sizes
func BenchmarkParseString_FastJSON_Long(b *testing.B) {
	input := `"hello world, this is a longer string to parse with more characters to test performance"`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			_, _ = ParseString(input)
	}
}

func BenchmarkParseString_StdLib_Long(b *testing.B) {
	input := `"hello world, this is a longer string to parse with more characters to test performance"`
	var result string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
			_ = json.Unmarshal([]byte(input), &result)
	}
}