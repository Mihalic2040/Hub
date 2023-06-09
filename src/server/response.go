package server

import "github.com/Mihalic2040/Hub/src/proto/api"

func Response(data string, status int) api.Response {
	return api.Response{
		Payload: data,
		Status:  int64(status),
	}
}
