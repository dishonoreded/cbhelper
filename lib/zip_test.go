package lib

import (
	"testing"
)

func TestGzip(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantErr  bool
	}{
		{
			name:     "empty input",
			input:    []byte{},
			wantErr:  false,
		},
		{
			name:     "simple text",
			input:    []byte("hello world"),
			wantErr:  false,
		},
		{
			name:     "longer text",
			input:    []byte("This is a longer text that should compress well with gzip compression algorithm"),
			wantErr:  false,
		},
		{
			name:     "binary data",
			input:    []byte{0, 1, 2, 3, 4, 5, 255, 254, 253},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Gzip(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Gzip() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("Gzip() error = %v", err)
			}
			if result == nil {
				t.Errorf("Gzip() returned nil result")
			}
		})
	}
}

func TestUnGzip(t *testing.T) {
	// First create gzipped test data
	originalData := []byte("test data for ungzip")
	gzippedData, err := Gzip(originalData)
	if err != nil {
		t.Fatalf("Failed to create test gzipped data: %v", err)
	}

	tests := []struct {
		name     string
		input    []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "valid gzipped data",
			input:    gzippedData,
			expected: originalData,
			wantErr:  false,
		},
		{
			name:     "invalid gzipped data",
			input:    []byte("not gzipped data"),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "empty input",
			input:    []byte{},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := UnGzip(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("UnGzip() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("UnGzip() error = %v", err)
			}
			if string(result) != string(tt.expected) {
				t.Errorf("UnGzip() = %v, want %v", string(result), string(tt.expected))
			}
		})
	}
}

func TestGzipUnGzipRoundTrip(t *testing.T) {
	testCases := [][]byte{
		[]byte("hello world"),
		[]byte(""),
		[]byte("This is a longer text that should test the compression and decompression thoroughly"),
		[]byte{0, 1, 2, 3, 255, 254, 253},
		[]byte("Special characters: !@#$%^&*()_+-=[]{}|;':\",./<>?"),
		[]byte("Unicode: 你好世界 🌍"),
	}

	for i, original := range testCases {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			// Compress
			compressed, err := Gzip(original)
			if err != nil {
				t.Fatalf("Gzip() error = %v", err)
			}

			// Decompress
			decompressed, err := UnGzip(compressed)
			if err != nil {
				t.Fatalf("UnGzip() error = %v", err)
			}

			// Compare
			if string(decompressed) != string(original) {
				t.Errorf("Round trip failed: original = %v, final = %v", string(original), string(decompressed))
			}
		})
	}
}