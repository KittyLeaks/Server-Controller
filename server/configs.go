package main

const (
	SERVER_IP          string = "193.47.61.254"
	SERVER_PORT        string = "1337"
	BOT_PORT           string = "54321"
	SERVER_ADDRESS     string = SERVER_IP + ":" + SERVER_PORT
	SERVER_BOT_ADDRESS string = SERVER_IP + ":" + BOT_PORT
	DEBUG_MODE         bool   = true
	CLEAR_SCREEN       string = "\033[2J\033[1H"
	USERNAME           string = "admin" // I think I should use a database for this.
	PASSWORD           string = "botnet"
	VERIFY_MESSAGE     string = "0x00"
	PING_MESSAGE       string = "0x01"
	PONG_MESSAGE       string = "0x02"
	PING_DELAY         int    = 5
	ACTION_PREFIX      string = "!"
)
