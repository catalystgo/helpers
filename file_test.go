package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		file      string
		data      []byte
		cfg       *SaveFileOpt
		expectErr bool
	}{
		{
			name:      "create_new_file",
			file:      "test_data/testfile1.txt",
			data:      []byte("hello world"),
			cfg:       &SaveFileOpt{Override: false},
			expectErr: false,
		},
		{
			name:      "skip_existing_file_without_override",
			file:      "test_data/testfile2.txt",
			data:      []byte("initial data"),
			cfg:       &SaveFileOpt{Override: false},
			expectErr: false,
		},
		{
			name:      "override_existing_file",
			file:      "test_data/testfile3.txt",
			data:      []byte("new content"),
			cfg:       &SaveFileOpt{Override: true},
			expectErr: false,
		},
		{
			name:      "empty_data_should_not_create_file",
			file:      "test_data/testfile4.txt",
			data:      []byte{},
			cfg:       &SaveFileOpt{Override: false},
			expectErr: false,
		},
	}

	require.NoError(t, os.MkdirAll("test_data", os.ModePerm))
	require.NoError(t, os.WriteFile("test_data/testfile2.txt", []byte("initial data"), 0644))
	require.NoError(t, os.WriteFile("test_data/testfile3.txt", []byte("old content"), 0644))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := SaveFile(tt.file, tt.data, tt.cfg)
			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if len(tt.data) > 0 && !tt.expectErr {
				storedData, err := os.ReadFile(tt.file)
				require.NoError(t, err)
				require.Equal(t, string(tt.data), string(storedData))
			}
		})
	}

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll("test_data"))
	})
}
