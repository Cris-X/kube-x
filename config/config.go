package config

// Server is a struct that contains the configuration of the server
type Server struct {
	System System `json:"system"`
	Logger Logger `json:"logger"`
	K8S    K8S    `json:"k8s"`
}
