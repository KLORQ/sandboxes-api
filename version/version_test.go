package version

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

func TestSourceSHA256(t *testing.T) {
	data, err := os.ReadFile("../proto/sandboxes/local/v1/sandbox.proto")
	if err != nil {
		t.Fatal(err)
	}
	if got := fmt.Sprintf("%x", sha256.Sum256(data)); got != SourceSHA256 {
		t.Fatalf("source hash = %s, want %s", got, SourceSHA256)
	}
}
