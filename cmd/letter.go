package cmd

import (
	"fmt"
	"github.com/adamyordan/postbox/postbox"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var (
	letterCmd = &cobra.Command{
		Use:   "letter",
		Short: "manage dumped requests",
	}

	letterViewCmd = &cobra.Command{
		Use:   "view",
		Short: "view a request",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("error parsing ID: %v", err)
			}
			view(uint64(id))
		},
	}

	letterListCmd = &cobra.Command{
		Use:   "list",
		Short: "list all dumped requests",
		Run: func(cmd *cobra.Command, args []string) {
			list()
		},
	}

	letterClearCmd = &cobra.Command{
		Use:   "clear",
		Short: "remove all dumped requests",
		Run: func(cmd *cobra.Command, args []string) {
			clear()
		},
	}
)

func init() {
	rootCmd.AddCommand(letterCmd)
	letterCmd.AddCommand(letterViewCmd)
	letterCmd.AddCommand(letterListCmd)
	letterCmd.AddCommand(letterClearCmd)
}

func view(id uint64) {
	letter, err := postbox.Get(id)
	if err != nil {
		log.Fatalf("error getting request with id %d: %v", id, err)
	}
	displayLetter(letter)
}

func list() {
	letters, err := postbox.List()
	if err != nil {
		log.Fatalf("error listing requests: %v", err)
	}
	fmt.Printf("Number of dumped requests: %d\n\n", len(letters))
	for _, letter := range letters {
		displayLetterShort(letter)
	}
}

func clear() {
	err := postbox.Clear()
	if err != nil {
		log.Fatalf("error clearing requests from db: %v", err)
	}
	log.Info("All stored requests cleared")
}

func displayLetter(letter *postbox.Letter) {
	fmt.Printf("%-6s: %d\n%-6s: %s\n%-6s: %s\n\n%s\n", "id", letter.ID, "ipaddr", letter.Ipaddr, "time",
		time.Unix(letter.Time, 0).String(), letter.Value)
}

func displayLetterShort(letter *postbox.Letter) {
	fmt.Printf("[%d] %s (%s)\n", letter.ID, time.Unix(letter.Time, 0).String(), letter.Ipaddr)
}
