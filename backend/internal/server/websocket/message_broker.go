package websocket

import "GoChat/internal/pb"

var (
	MessageChannel = make(chan *pb.Message, 100)
)
