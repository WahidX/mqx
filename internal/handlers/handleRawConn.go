package handlers

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"

	"go.uber.org/zap"
)

func HandleRawConn(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	ctx := context.Background()
	defer conn.Close()

	// Reading the first byte(command) and handling error (if connection closed or something)
	commandByte, err := reader.ReadByte()
	if err == io.EOF {
		zap.L().Warn("Client disconnected", zap.Error(err))
		return err
	}
	if err != nil {
		zap.L().Warn("Error reading command byte", zap.Error(err))
		return err
	}

	switch Command(commandByte) {
	case Ping:
		err = connHandler.Ping(ctx, conn)

	case Publish:
		err = connHandler.Publish(ctx, reader, conn)

	case Listen:
		err = connHandler.Listen(ctx, reader, conn)

	default:
		zap.L().Info("Unknown command", zap.Any("incoming command", commandByte))
		conn.Write([]byte("Unknown command"))

	}

	// Need to handle the error in a better way
	// Types of errors can occur:
	// 1. Connection closed by client
	// 2. Some server error
	// 3. Some client input error
	if err != nil {
		conn.Write([]byte("E01"))
	} else {
		fmt.Println("HEREHEHEHEHHER")
		conn.Write([]byte("1")) // 1 means success
	}

	return nil
}
