package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"server/server/game"
	"server/server/game/tree"
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

	log.Printf("reqeust to %s", r.URL)
	r.ParseForm()
	fmt.Println(r.PostForm.Get("id"),r.FormValue("id"))
	id := r.FormValue("id")
	log.Printf("reqeust id %s", id)
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
		bts, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Printf("read message error: %v", err)
			return
		}
		fmt.Println(string(bts))

		t := tree.NewTree()
		resp := t.ToMap()
		resp["msgId"] = 1
		res, err := json.Marshal(resp)
		fmt.Println(string(res))
		err = wsutil.WriteServerMessage(conn, op, res)
		if err != nil {
			log.Printf("write message error: %v", err)
			return
		}
	}

}
