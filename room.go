package main

type room struct {
	// forward is a channel that holds incomijng messages
	// that should be forwareded to the other clients
	forward chan []byte
}
