package version

import (
	"testing"

	localv1 "github.com/rheic/sandboxes-api/gen/go/sandboxes/local/v1"

	"google.golang.org/protobuf/proto"
)

func TestMCPCallContractRoundTrip(t *testing.T) {
	want := &localv1.MCPCallRequest{
		SandboxId:     "sandbox-test",
		ServerName:    "fixture",
		ToolName:      "fixture-pong",
		ArgumentsJson: `{"message":"hello"}`,
	}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got localv1.MCPCallRequest
	if err := proto.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) {
		t.Fatalf("round-trip mismatch: got %v, want %v", &got, want)
	}

	service := localv1.File_sandboxes_local_v1_sandbox_proto.Services().ByName("SandboxService")
	if service == nil || service.Methods().ByName("MCPCall") == nil {
		t.Fatal("SandboxService.MCPCall RPC is missing from the generated descriptor")
	}
}
