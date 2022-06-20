package waste

import (
	"os"
	"fmt"
	"runtime"
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
		return put(args[0])
	},
}

func init() {
	rootCmd.AddCommand(putSubCmd)
}

func put(target string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var trash string
	switch runtime.GOOS {
		case "darwin":
			trash = "/.Trash/"
		default:
			return errors.New("unsupported OS") 
	}
	if err = os.Rename(target, home + trash + target); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("succeeded")
	return nil
}
