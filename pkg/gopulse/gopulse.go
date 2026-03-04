package gopulse

import "github.com/prasanththumma/gopulse/internal/server"

func Start() {
	go server.StartServer()
}
