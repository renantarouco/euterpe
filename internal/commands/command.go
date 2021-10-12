package commands

const (
	CommandSetPlayerID = "setplayerid"
	CommandStart       = "start"
	CommandNext        = "next"
)

type Command struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type SetPlayerIDPayload struct {
	PlayerID uint `json:"playerID"`
}

func NewSetPlayerIDCommand(playerID uint) Command {
	return Command{
		Type: CommandSetPlayerID,
		Payload: SetPlayerIDPayload{
			PlayerID: playerID,
		},
	}
}

func NewStartCommand() Command {
	return Command{
		Type: CommandStart,
	}
}

func NewNextCommand() Command {
	return Command{
		Type: CommandNext,
	}
}
