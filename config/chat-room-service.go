package config

type ChatRoomServiceConfig struct {
	CodecKey       string `json:"codecKey"`
	BroadcastTopic string `json:"broadcastTopic"`
	EventTopic     string `json:"eventTopic"`
}
