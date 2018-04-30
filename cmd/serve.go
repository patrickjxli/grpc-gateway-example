package cmd

import (
	"fmt"
	"net"

	"github.com/Stoakes/grpc-gateway-example/server"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on http://localhost:10000",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Server initialisation
func serve() {

	// Initialize listener
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Instanciate the server
	s := server.New(fmt.Sprintf("localhost:%d", server.Port), conn)
	s.Start()
}
