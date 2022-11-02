package database

import "Implementing_an_inmemory_database/interface/resp"
import "Implementing_an_inmemory_database/resp/reply"

func Ping(db *DB, args [][]byte) resp.Reply {
	return reply.MakePongReply()
}

// PING
func init() {
	RegisterCommand("ping", Ping, 1)
}
