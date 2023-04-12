// This is derived from
// https://gist.github.com/staaldraad/d835126cd46969330a8fdadba62b9b69

package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/webdav"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var lock sync.RWMutex
var timeout time.Duration
var lastOp time.Time

func main() {

	baseUrlF := flag.String("b", "/", "Base URL of server [/].")
	dirF := flag.String("d", "./", "Root of served directory tree [CWD].")
	portF := flag.Int("p", 8000, "Port to serve on [8000].")
	timeoutF := flag.Int("t", 600, "Idle timeout in seconds [600].")

	flag.Parse()

	lastOp = time.Now()
	timeout = time.Duration(int(1e9) * *timeoutF)
	bindAddr := fmt.Sprintf(":%d", *portF)

	dir := *dirF
	baseUrl := *baseUrlF

	go reap()
	serve(baseUrl, dir, bindAddr)
}

func reap() {
	// We rely on the global lastOp being updated by the route handler
	log.Print("[REAPER] Starting.")
	for {
		lock.RLock()
		since := time.Since(lastOp)
		lock.RUnlock()
		if since > timeout {
			// This is our normal exit, hence the zero return code
			log.Printf("[REAPER] No operations in %v; shutting down.", timeout)
			os.Exit(0)
		}
		time.Sleep(time.Second)
	}
}

func serve(baseUrl string, dir string, bindAddr string) {
	log.Print("[SERVER] Starting.")
	dav := &webdav.Handler{
		FileSystem: webdav.Dir(dir),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			// We're totally abusing the logger here to update
			// the global lastOp, since it's called on every
			// request.
			lock.Lock()
			lastOp = time.Now()
			lock.Unlock()
			// But we might as well log the action too.
			logmsg := fmt.Sprintf("[SERVER] %s %s", r.Method, r.URL)
			if err != nil {
				logmsg += fmt.Sprintf(" -> ERROR: %s", err)
			}
			log.Printf(logmsg)
		},
	}
	http.Handle(baseUrl, dav)

	if err := http.ListenAndServe(bindAddr, nil); err != nil {
		log.Fatalf("[SERVER] ERROR: %w", err)
	}

}
