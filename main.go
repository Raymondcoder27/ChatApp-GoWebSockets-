package main

import "fmt"

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}
func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client.", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func readLoop(ws *websocket.Conn) {

}

func main() {

}
