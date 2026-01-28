package cli

import (
	"github.com/spf13/cobra"
	"strings"
)

func parseArgs(cmd *cobra.Command, args []string) string {
	msg := strings.Join(args, " ")
	return msg
}
