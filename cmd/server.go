package cmd

import (
	"github.com/adamyordan/postbox/daemon"
	"github.com/adamyordan/postbox/postbox"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	daemonFlag bool
	addressFlag string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "manage request endpoint listener",
	}

	serverUpCmd = &cobra.Command{
		Use:   "up",
		Short: "start server",
		Run: func(cmd *cobra.Command, args []string) {
			up(daemonFlag, addressFlag)
		},
	}

	serverDownCmd = &cobra.Command{
		Use:   "down",
		Short: "stop server",
		Run: func(cmd *cobra.Command, args []string) {
			down()
		},
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverUpCmd)
	serverCmd.AddCommand(serverDownCmd)

	serverUpCmd.Flags().BoolVarP(&daemonFlag, "daemon", "d", false, "run server in background")
	serverUpCmd.Flags().StringVarP(&addressFlag, "address", "a", ":8000", "specify local address for server")
}

func up(isDaemon bool, addr string) {
	if isDaemon {
		daemon.StartDaemon(func() error {
			return postbox.ServeHttp(addr)
		})
	} else {
		log.Fatal(postbox.ServeHttp(addr))
	}
}

func down() {
	daemon.StopDaemon()
}
