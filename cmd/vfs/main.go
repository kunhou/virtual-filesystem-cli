package main

import (
	"github/kunhou/virtual-filesystem-cli/internal/deliver/cli"
	"github/kunhou/virtual-filesystem-cli/internal/repository"
	"github/kunhou/virtual-filesystem-cli/internal/usecase"
)

func main() {
	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	cliSrv := cli.NewCLIServer(uc)

	cliSrv.Run()
}
