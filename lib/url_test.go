package lib

import (
	"testing"
)

func TestURLEscape(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "simple text",
			input:    []byte("hello world"),
			expected: "hello+world",
		},
		{
			name:     "special characters",
			input:    []byte("hello & world @ 2023"),
			expected: "hello+%26+world+%40+2023",
		},
		{
			name:     "url with parameters",
			input:    []byte("http://example.com?param=value&other=test"),
			expected: "http%3A%2F%2Fexample.com%3Fparam%3Dvalue%26other%3Dtest",
		},
		{
			name:     "chinese characters",
			input:    []byte("你好世界"),
			expected: "%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := URLEscape(tt.input)
			if err != nil {
				t.Fatalf("URLEscape() error = %v", err)
			}
			if string(result) != tt.expected {
				t.Errorf("URLEscape() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestURLUnescape(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
		wantErr  bool
	}{
		{
			name:     "empty input",
			input:    []byte{},
			expected: "",
			wantErr:  false,
		},
		{
			name:     "simple text",
			input:    []byte("hello+world"),
			expected: "hello world",
			wantErr:  false,
		},
		{
			name:     "special characters",
			input:    []byte("hello+%26+world+%40+2023"),
			expected: "hello & world @ 2023",
			wantErr:  false,
		},
		{
			name:     "url with parameters",
			input:    []byte("http%3A%2F%2Fexample.com%3Fparam%3Dvalue%26other%3Dtest"),
			expected: "http://example.com?param=value&other=test",
			wantErr:  false,
		},
		{
			name:     "chinese characters",
			input:    []byte("%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C"),
			expected: "你好世界",
			wantErr:  false,
		},
		{
			name:     "invalid escape sequence",
			input:    []byte("hello%ZZ"),
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := URLUnescape(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("URLUnescape() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("URLUnescape() error = %v", err)
			}
			if string(result) != tt.expected {
				t.Errorf("URLUnescape() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestURLEscapeUnescape(t *testing.T) {
	testCases := [][]byte{
		[]byte("hello world"),
		[]byte("test@example.com"),
		[]byte("http://example.com?param=value"),
		[]byte("你好世界"),
		[]byte("special!@#$%^&*()chars"),
	}

	for _, original := range testCases {
		escaped, err := URLEscape(original)
		if err != nil {
			t.Fatalf("URLEscape() error = %v", err)
		}

		unescaped, err := URLUnescape(escaped)
		if err != nil {
			t.Fatalf("URLUnescape() error = %v", err)
		}

		if string(unescaped) != string(original) {
			t.Errorf("Round trip failed: original = %v, final = %v", string(original), string(unescaped))
		}
	}
}