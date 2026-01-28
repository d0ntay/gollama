package cli

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/d0ntay/gollama/internal/api"
	
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gollama",
	Short: "gollama is a cli based app to interface with ollama powered by go!",
	Long: "gollama is a cli based app to interface with ollama powered by go! Automaticcaly stores message history as well as allows web search.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		msg := parseArgs(cmd, args)

		resp, err := api.Chat(msg)
		if err != nil {
			return err
		}

		fmt.Println(resp)
		return nil
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
