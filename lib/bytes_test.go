package lib

import (
	"reflect"
	"testing"
)

func TestBytesArrayDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "empty array",
			input:    []byte("[]"),
			expected: []byte{},
			wantErr:  false,
		},
		{
			name:     "single byte",
			input:    []byte("[65]"),
			expected: []byte{65},
			wantErr:  false,
		},
		{
			name:     "multiple bytes with spaces",
			input:    []byte("[72 101 108 108 111]"),
			expected: []byte{72, 101, 108, 108, 111},
			wantErr:  false,
		},
		{
			name:     "multiple bytes with commas",
			input:    []byte("[72,101,108,108,111]"),
			expected: []byte{72, 101, 108, 108, 111},
			wantErr:  false,
		},
		{
			name:     "mixed separators",
			input:    []byte("[72, 101 108, 108 111]"),
			expected: []byte{72, 101, 108, 108, 111},
			wantErr:  false,
		},
		{
			name:     "with extra whitespace",
			input:    []byte("  [ 72  101   108 108 111 ]  "),
			expected: []byte{72, 101, 108, 108, 111},
			wantErr:  false,
		},
		{
			name:     "boundary values",
			input:    []byte("[0 255]"),
			expected: []byte{0, 255},
			wantErr:  false,
		},
		{
			name:     "missing opening bracket",
			input:    []byte("72 101 108]"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "missing closing bracket",
			input:    []byte("[72 101 108"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "no brackets",
			input:    []byte("72 101 108"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "invalid number",
			input:    []byte("[abc]"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "number too large",
			input:    []byte("[256]"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "negative number",
			input:    []byte("[-1]"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "mixed valid and invalid",
			input:    []byte("[72 abc 108]"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "empty string",
			input:    []byte(""),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "only whitespace",
			input:    []byte("   "),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BytesArrayDecode(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("BytesArrayDecode() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("BytesArrayDecode() error = %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BytesArrayDecode() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBytesArrayDecodeHelloWorld(t *testing.T) {
	input := []byte("[72 101 108 108 111 32 87 111 114 108 100]")
	expected := "Hello World"
	
	result, err := BytesArrayDecode(input)
	if err != nil {
		t.Fatalf("BytesArrayDecode() error = %v", err)
	}
	
	if string(result) != expected {
		t.Errorf("BytesArrayDecode() decoded string = %v, want %v", string(result), expected)
	}
}