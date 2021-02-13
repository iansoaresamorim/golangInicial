package app

import "github.com/urfave/cli"

// Gerar vai retonar a aplicacao de linha de comando prona para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca Ipps e Nomes de Servidor na internet"

	return app
}
