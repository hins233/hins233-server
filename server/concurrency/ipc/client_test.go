package ipc

import (
	"reflect"
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{
		Code: "OK",
		Body: "Echo: " + method + "~" + params,
	}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpcClient_Call(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	session := server.Connect()
	type fields struct {
		conn chan string
	}
	type args struct {
		method string
		params string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *Response
		wantErr  bool
	}{
		{
			name:   "echo client",
			fields: fields{conn:session},
			args:args{
				method: "create room",
				params: "id: frank,room: 10001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &IpcClient{
				conn: tt.fields.conn,
			}
			gotResp, err := client.Call(tt.args.method, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Call() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestIpcClient_close(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	type fields struct {
		conn chan string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:"close",
			fields:fields{conn:client1.conn},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client1.close()
		})
	}
}

func TestIpcServer_Connect(t *testing.T) {
	type fields struct {
		Server Server
	}
	server := &EchoServer{}
	session := make(chan string, 0)
	tests := []struct {
		name   string
		fields fields
		want   chan string
	}{
		{
			name:"test1",
			fields:fields{Server:server},
			want:session,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &IpcServer{
				Server: tt.fields.Server,
			}
			if got := server.Connect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIpcClient(t *testing.T) {
	type args struct {
		server *IpcServer
	}
	tests := []struct {
		name string
		args args
		want *IpcClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIpcClient(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIpcClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIpcServer(t *testing.T) {
	type args struct {
		server Server
	}
	tests := []struct {
		name string
		args args
		want *IpcServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIpcServer(tt.args.server); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIpcServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
