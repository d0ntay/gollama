package cli

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "gollama",
	Short: "gollama is a cli based app to interface with ollama powered by go!",
	Long: "gollama is a cli based app to interface with ollama powered by go! ",
}
