package database

import (
	"Implementing_an_inmemory_database/interface/database"
	"Implementing_an_inmemory_database/interface/resp"
	"Implementing_an_inmemory_database/resp/reply"
)

// GET k1
func execGet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		return reply.MakeNullBulkReply()
	}
	bytes := entity.Data.([]byte)
	return reply.MakeBulkReply(bytes)
}

// SET k1 v
func execSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity := &database.DataEntity{
		Data: value,
	}
	db.PutEntity(key, entity)
	return reply.MakeOkReply()
}

// SETNX k1 v1
func execSetnx(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity := &database.DataEntity{
		Data: value,
	}

	result := db.PutIfAbsent(key, entity)
	return reply.MakeIntReply(int64(result))
}

// GETSET k1 v1
func execGetSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity, exists := db.GetEntity(key)
	db.PutEntity(key, &database.DataEntity{
		Data: value,
	})
	if !exists {
		return reply.MakeNullBulkReply()
	}
	return reply.MakeBulkReply(entity.Data.([]byte))
}

// STRLEN k -> 'value'
func execStrLen(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		return reply.MakeNullBulkReply()
	}
	bytes := entity.Data.([]byte)
	return reply.MakeIntReply(int64(len(bytes)))
}

func init() {
	RegisterCommand("Get", execGet, 2) //get k1
	RegisterCommand("Set", execSet, 3) // set k1 v1
	RegisterCommand("SetNX", execSetnx, 3)
	RegisterCommand("GetSet", execGetSet, 3)
	RegisterCommand("StrLen", execStrLen, 2)
}
