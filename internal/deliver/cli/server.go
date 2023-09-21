package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github/kunhou/virtual-filesystem-cli/internal/usecase"
	"github/kunhou/virtual-filesystem-cli/pkg/log"
	"github/kunhou/virtual-filesystem-cli/pkg/parser"
)

type CLIServer struct {
	handlers map[string]CmdHandler
	usecase  *usecase.Usecase
}

type CmdHandler func(args []string)

func NewCLIServer(u *usecase.Usecase) *CLIServer {
	srv := CLIServer{
		handlers: make(map[string]CmdHandler),
		usecase:  u,
	}

	srv.registerHandler("register", srv.RegisterUserHandler)
	srv.registerHandler("create-folder", srv.CreateFolderHandler)
	srv.registerHandler("delete-folder", srv.DeleteFolderHandler)
	srv.registerHandler("list-folders", srv.ListFoldersHandler)

	return &srv
}

func (s *CLIServer) registerHandler(cmd string, handler CmdHandler) {
	s.handlers[cmd] = handler
}

func (s *CLIServer) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("# ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			continue
		}

		args := parser.ParseInput(input)

		if handler, exists := s.handlers[args[0]]; exists {
			handler(args[1:])
		} else {
			log.Error("Unknown command %s", args[0])
		}
	}
}
