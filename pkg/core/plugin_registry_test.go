package core

import (
	"os"
	"testing"

	"github.com/danielmiessler/fabric/pkg/plugins/db/fsdb"
)

func TestSaveEnvFile(t *testing.T) {
	registry := NewPluginRegistry(fsdb.NewDb(os.TempDir()))

	err := registry.SaveEnvFile()
	if err != nil {
		t.Fatalf("SaveEnvFile() error = %v", err)
	}
}
