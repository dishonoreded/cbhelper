package lib

import (
	"testing"
)

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
		wantErr  bool
	}{
		{
			name:     "simple json",
			input:    []byte(`{"name":"test","age":25}`),
			expected: "{\n \"age\": 25,\n \"name\": \"test\"\n}",
			wantErr:  false,
		},
		{
			name:     "nested json",
			input:    []byte(`{"user":{"name":"test","details":{"age":25,"city":"NYC"}}}`),
			expected: "{\n \"user\": {\n  \"details\": {\n   \"age\": 25,\n   \"city\": \"NYC\"\n  },\n  \"name\": \"test\"\n }\n}",
			wantErr:  false,
		},
		{
			name:     "array json",
			input:    []byte(`{"items":[1,2,3],"count":3}`),
			expected: "{\n \"count\": 3,\n \"items\": [\n  1,\n  2,\n  3\n ]\n}",
			wantErr:  false,
		},
		{
			name:     "empty object",
			input:    []byte(`{}`),
			expected: "{}",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("FormatJSON() panic = %v", r)
					}
				}
			}()

			result := FormatJSON(tt.input)
			if string(result) != tt.expected {
				t.Errorf("FormatJSON() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestFormatJSONInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("FormatJSON() should panic with invalid JSON")
		}
	}()

	FormatJSON([]byte(`invalid json`))
}