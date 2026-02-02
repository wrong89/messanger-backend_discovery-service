package models

import "strings"

type Server struct {
	Title string
	Host  string
	Port  string
}

func NewServer(title, host, port string) Server {
	return Server{Title: title, Host: host, Port: port}
}

func (s Server) GetAddress() string {
	args := []string{s.Host, s.Port}

	return strings.Join(args, ":")
}
