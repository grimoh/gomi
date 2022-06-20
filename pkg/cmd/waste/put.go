package waste

import (
	"os"
	"fmt"
	"errors"

	"github.com/spf13/cobra"
)

var putSubCmd = &cobra.Command{
	Use: "put",
	Short: "Throw waste in the trash.",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected 1 arg.")
		}
		return convert(args[0])
	},
}

func init() {
	rootCmd.AddCommand(jstSubCmd)
}

func put(target string) error {
	if err := os.Rename(target, "~/.Trash/"); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("succeeded")
	return nil
}
