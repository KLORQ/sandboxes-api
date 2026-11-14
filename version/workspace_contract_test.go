package version

import (
	"testing"

	localv1 "github.com/rheic/sandboxes-api/gen/go/sandboxes/local/v1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestWorkspaceGuestPathContractRoundTrip(t *testing.T) {
	want := &localv1.Workspace{Path: "/host/repo", ReadOnly: true, GuestPath: "/run/sandbox/source"}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got localv1.Workspace
	if err := proto.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) {
		t.Fatalf("round-trip mismatch: got %v, want %v", &got, want)
	}

	field := (&localv1.Workspace{}).ProtoReflect().Descriptor().Fields().ByName("guest_path")
	if field == nil || field.Number() != 3 || field.Kind() != protoreflect.StringKind {
		t.Fatal("Workspace.guest_path field 3 is missing from the generated descriptor")
	}
	if got.GuestPath != "/run/sandbox/source" {
		t.Fatalf("guest path = %q, want explicit destination", got.GuestPath)
	}
	if (&localv1.Workspace{}).GetGuestPath() != "" {
		t.Fatal("empty guest_path must remain the legacy same-path default")
	}
}
