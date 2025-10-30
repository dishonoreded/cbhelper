package lib

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
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
			input:    []byte("hello"),
			expected: "aGVsbG8=",
		},
		{
			name:     "special characters",
			input:    []byte("hello world!@#"),
			expected: "aGVsbG8gd29ybGQhQCM=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base64Encode(tt.input)
			if err != nil {
				t.Fatalf("Base64Encode() error = %v", err)
			}
			if string(result) != tt.expected {
				t.Errorf("Base64Encode() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
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
			input:    []byte("aGVsbG8="),
			expected: "hello",
			wantErr:  false,
		},
		{
			name:     "special characters",
			input:    []byte("aGVsbG8gd29ybGQhQCM="),
			expected: "hello world!@#",
			wantErr:  false,
		},
		{
			name:     "invalid base64",
			input:    []byte("invalid!"),
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base64Decode(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Base64Decode() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Base64Decode() error = %v", err)
			}
			if string(result) != tt.expected {
				t.Errorf("Base64Decode() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestUnGzipWithBase64Decoding(t *testing.T) {
	// First create a gzipped and base64 encoded test data
	originalData := []byte("test data for gzip and base64")
	
	// Gzip the data
	gzippedData, err := Gzip(originalData)
	if err != nil {
		t.Fatalf("Failed to gzip test data: %v", err)
	}
	
	// Base64 encode the gzipped data
	base64Data, err := Base64Encode(gzippedData)
	if err != nil {
		t.Fatalf("Failed to base64 encode test data: %v", err)
	}

	tests := []struct {
		name     string
		input    []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "valid gzipped base64 data",
			input:    base64Data,
			expected: originalData,
			wantErr:  false,
		},
		{
			name:     "invalid base64",
			input:    []byte("invalid!"),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnGzipWithBase64Decoding(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("UnGzipWithBase64Decoding() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("UnGzipWithBase64Decoding() error = %v", err)
			}
			if string(result) != string(tt.expected) {
				t.Errorf("UnGzipWithBase64Decoding() = %v, want %v", string(result), string(tt.expected))
			}
		})
	}
}