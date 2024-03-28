package handler

import "github.com/su-its/typing/typing-server/domain/repository/ent"

var entClient *ent.Client

func SetEntClient(client *ent.Client) {
	entClient = client
}
