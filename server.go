package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"code.google.com/p/go.net/websocket"
)

var (
	idSequence   = 0
	number       = 0
	clients      = make(map[int]*websocket.Conn)
	connectMutex sync.Mutex
	numberMutex  sync.Mutex
)

func main() {
	http.Handle("/data", websocket.Handler(handleSocket))
	log.Println("About to listen for HTTP traffic at localhost:9000")
	http.ListenAndServe("localhost:9000", nil)
}

func handleSocket(ws *websocket.Conn) {
	log.Println("Got connection")
	connectMutex.Lock()
	id := idSequence
	idSequence += 1
	clients[id] = ws
	connectMutex.Unlock()

	_, err := ws.Write([]byte(strconv.Itoa(number)))
	if err != nil {
		log.Printf("Unable to write: %s", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		readBuffer := make([]byte, 1024)
		for {
			n, err := ws.Read(readBuffer)
			if err != nil {
				connectMutex.Lock()
				defer connectMutex.Unlock()
				delete(clients, id)
				ws.Close()
				wg.Done()
				log.Println("Closed connection")
				return
			}
			numString := string(readBuffer[0:n])
			val, err := strconv.Atoi(numString)
			if err != nil {
				log.Printf("Unable to parse number string %s: %s", numString, err)
				continue
			}
			numberMutex.Lock()
			number += val
			numberMutex.Unlock()
			for _, peer := range clients {
				peer.Write([]byte(strconv.Itoa(number)))
			}
		}
	}()

	wg.Wait()
}
