package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	json "github.com/json-iterator/go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"server/server/game"
	"syscall"
	"time"
)

var addr = flag.String("listen", ":8000", "addr to listen")

func main() {
	log.SetFlags(0)
	flag.Parse()
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	http.HandleFunc("/hello", handlerIndex)
	http.HandleFunc("/ws", wsHandler)

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen %q error: %v", *addr, err)
	}
	log.Printf("listening %s (%q)", ln.Addr(), *addr)

	var (
		s     = new(http.Server)
		serve = make(chan error, 1)
		sig   = make(chan os.Signal, 1)
	)
	signal.Notify(sig, syscall.SIGTERM)
	go func() { serve <- s.Serve(ln) }()

	select {
	case err := <-serve:
		log.Fatal(err)
	case sig := <-sig:
		const timeout = 5 * time.Second

		log.Printf("signal %q received; shutting down with %s timeout", sig, timeout)

		ctx, _ := context.WithTimeout(context.Background(), timeout)
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	log.Printf("--------- %s --reqeust to %s --reqeust id %s", time.Now().Format("2006-01-02 15:04:05"), r.URL, id)
	game.RenderHtml(w, id, nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	defer conn.Close()

	for {
		msg, _, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Printf("read message error: %v", err)
			return
		}
		fmt.Println(string(msg))
		param := make(map[string]interface{})
		err = json.Unmarshal(msg, &param)
		if err != nil {
			log.Printf("msg Unmarshal failed,err %v", err)
		}
		game.Dispatcher(param, conn)
	}

}
