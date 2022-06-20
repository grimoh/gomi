package rubbish

import (
	"os"
	"fmt"
	"runtime"
	"errors"

	"github.com/spf13/cobra"
)

var chuckoutSubCmd = &cobra.Command{
	Use: "chuckout",
	Short: "chuck out rubbish in the trash box",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected 1 arg")
		}
		return chuckout(args[0])
	},
}

func init() {
	rootCmd.AddCommand(chuckoutSubCmd)
}

func chuckout(target string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var rubbish string
	switch runtime.GOOS {
		case "darwin":
			rubbish = "/.Trash/"
		default:
			return errors.New("unsupported OS") 
	}
	if err = os.Rename(target, home + rubbish + target); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("succeeded")
	return nil
}
