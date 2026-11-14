package version

import (
	"testing"

	localv1 "github.com/rheic/sandboxes-api/gen/go/sandboxes/local/v1"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestCloneContractFieldNumbersAndWireRoundTrip(t *testing.T) {
	requestDescriptor := (&localv1.CreateRequest{}).ProtoReflect().Descriptor()
	cloneRequestField := requestDescriptor.Fields().ByName("clone")
	if cloneRequestField == nil || cloneRequestField.Number() != 17 || cloneRequestField.Kind() != protoreflect.MessageKind {
		t.Fatal("CreateRequest.clone field 17 is missing from the generated descriptor")
	}
	sandboxDescriptor := (&localv1.Sandbox{}).ProtoReflect().Descriptor()
	cloneInfoField := sandboxDescriptor.Fields().ByName("clone_info")
	if cloneInfoField == nil || cloneInfoField.Number() != 22 || cloneInfoField.Kind() != protoreflect.MessageKind {
		t.Fatal("Sandbox.clone_info field 22 is missing from the generated descriptor")
	}
	guestWorkdirField := (&localv1.CloneOptions{}).ProtoReflect().Descriptor().Fields().ByName("guest_workdir")
	if guestWorkdirField == nil || !guestWorkdirField.HasPresence() || guestWorkdirField.Kind() != protoreflect.StringKind {
		t.Fatal("CloneOptions.guest_workdir must be an optional string with presence")
	}

	emptyWorkdir := ""
	want := &localv1.CreateRequest{Clone: &localv1.CloneOptions{
		SourceWorkspacePath: "/host/project", GuestWorkdir: &emptyWorkdir,
	}}
	encoded, err := proto.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got localv1.CreateRequest
	if err := proto.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(want, &got) || got.GetClone().GetSourceWorkspacePath() != "/host/project" {
		t.Fatalf("clone request round-trip mismatch: got %v, want %v", &got, want)
	}
	if got.GetClone().GuestWorkdir == nil || *got.GetClone().GuestWorkdir != "" {
		t.Fatal("explicit empty guest_workdir presence was not preserved")
	}

	info := &localv1.CloneInfo{GuestWorkdir: "/workspace/project", RemoteName: "origin"}
	withoutEndpoint, err := proto.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}
	var stoppedInfo localv1.CloneInfo
	if err := proto.Unmarshal(withoutEndpoint, &stoppedInfo); err != nil {
		t.Fatal(err)
	}
	if stoppedInfo.GitEndpoint != nil {
		t.Fatal("clone endpoint must remain absent when not provided")
	}

	wantSandbox := &localv1.Sandbox{CloneInfo: &localv1.CloneInfo{
		GuestWorkdir: "/workspace/project", RemoteName: "origin",
		GitEndpoint: &localv1.GitEndpoint{HostIp: "127.0.0.1", Port: 44121, Repository: "project.git"},
	}}
	encoded, err = proto.Marshal(wantSandbox)
	if err != nil {
		t.Fatal(err)
	}
	var gotSandbox localv1.Sandbox
	if err := proto.Unmarshal(encoded, &gotSandbox); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(wantSandbox, &gotSandbox) {
		t.Fatalf("sandbox clone metadata round-trip mismatch: got %v, want %v", &gotSandbox, wantSandbox)
	}
}
