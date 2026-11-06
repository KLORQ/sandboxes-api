package version

import (
	"testing"

	localv1 "github.com/rheic/sandboxes-api/gen/go/sandboxes/local/v1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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

func TestListMCPServersContractRoundTrip(t *testing.T) {
	want := &localv1.ListMCPServersResponse{Servers: []*localv1.MCPServerSummary{{
		Name: "fixture", Gateway: "local", Transport: "stdio", Loaded: true,
		Usable: true, AuthRequired: true, AuthConfigured: false,
		HasCustomHeaders: true, CustomHeadersSupported: true,
		CustomHeaderSecretsConfigured: false,
	}}}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got localv1.ListMCPServersResponse
	if err := proto.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) {
		t.Fatalf("round-trip mismatch: got %v, want %v", &got, want)
	}

	service := localv1.File_sandboxes_local_v1_sandbox_proto.Services().ByName("SandboxService")
	if service == nil || service.Methods().ByName("ListMCPServers") == nil {
		t.Fatal("SandboxService.ListMCPServers RPC is missing from the generated descriptor")
	}
}

func TestMCPApprovalAndRegistrationContracts(t *testing.T) {
	service := localv1.File_sandboxes_local_v1_sandbox_proto.Services().ByName("SandboxService")
	for _, method := range []string{"MCPRegisterServer", "MCPApproveRegistration", "MCPRemoveServer", "MCPCall", "MCPApproveCall"} {
		if service == nil || service.Methods().ByName(protoreflect.Name(method)) == nil {
			t.Fatalf("SandboxService.%s RPC is missing from the generated descriptor", method)
		}
	}

	want := &localv1.MCPCallResponse{ApprovalRequired: true, ApprovalId: "challenge-test", ApprovalReason: "MCP tool execution"}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got localv1.MCPCallResponse
	if err := proto.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) {
		t.Fatalf("round-trip mismatch: got %v, want %v", &got, want)
	}

	registration := &localv1.MCPRegisterServerRequest{Name: "fixture", Command: "/tmp/fixture", Arguments: []string{"--stdio"}, WorkingDirectory: "/tmp"}
	encoded, err = proto.Marshal(registration)
	if err != nil {
		t.Fatal(err)
	}
	var decoded localv1.MCPRegisterServerRequest
	if err := proto.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(registration, &decoded) {
		t.Fatalf("registration round-trip mismatch: got %v, want %v", &decoded, registration)
	}
}
