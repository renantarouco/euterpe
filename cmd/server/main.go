package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/renantarouco/euterpe/internal/commands"
)

const (
	upgraderReadBufferSize  = 1024
	upgraderWriteBufferSize = 1024
	maxPlayers              = 5
)

func main() {
	fmt.Println("Euterpe - Server")

	upgrader := websocket.Upgrader{
		ReadBufferSize:  upgraderReadBufferSize,
		WriteBufferSize: upgraderWriteBufferSize,
	}

	players := map[uint]*websocket.Conn{}

	log.Println("starting with", len(players), "players")

	http.HandleFunc("/play", func(rw http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(rw, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		if len(players) > 5 {
			log.Println(errors.New("too many players"))
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		playerID := uint(len(players) + 1)
		players[playerID] = conn
		defer delete(players, playerID)

		log.Println("new player connected, now with", len(players), "players")

		err = conn.WriteJSON(commands.NewSetPlayerIDCommand(playerID))
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		for {
			var command commands.Command

			err := conn.ReadJSON(&command)
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
					log.Println(err)
					rw.WriteHeader(http.StatusBadRequest)
				}
				return
			}

			log.Println("received", command.Type, "command")
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
