package pusher

import (
	"encoding/json"
	)

type eventPayload struct {
	Name     string   `json:"name"`
	Channels []string `json:"channels"`
	Data     string   `json:"data"`
	SocketId *string  `json:"socket_id,omitempty"`
}

type BufferedEvents struct {
	EventIds map[string]string `json:"event_ids,omitempty"`
}

func createTriggerPayload(channels []string, event string, data interface{}, socketId *string) ([]byte, error) {
	data2, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return json.Marshal(&eventPayload{
		Name:     event,
		Channels: channels,
		Data:     string(data2),
		SocketId: socketId,
	})
}
