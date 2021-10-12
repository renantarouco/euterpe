package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/renantarouco/euterpe/internal/commands"
)

func main() {
	fmt.Println("Euterpe - Shell")

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8000/play", nil)
	if err != nil {
		log.Fatal(err)
	}

	var command commands.Command

	err = conn.ReadJSON(&command)
	if err != nil {
		log.Fatal(err)
	}

	if command.Type != commands.CommandSetPlayerID {
		log.Fatal(errors.New("not set player id command"))
	}

	setPlayerIDPayload, ok := command.Payload.(commands.SetPlayerIDPayload)
	if !ok {
		log.Fatal(errors.New("couldn't cast set player id payload"))
	}

	playerID := setPlayerIDPayload.PlayerID
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("player %d> ", playerID)

		commandString, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}

		commandString = strings.TrimSuffix(commandString, "\n")

		log.Println(commandString)
	}
}
