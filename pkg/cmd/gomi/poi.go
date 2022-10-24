package gomi

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var poiSubCmd = &cobra.Command{
	Use:   "poi",
	Short: "Trash in the trash box.",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected 1 arg")
		}
		return poi(args[0])
	},
}

func init() {
	rootCmd.AddCommand(poiSubCmd)
}

func poi(target string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var trash string
	switch runtime.GOOS {
	case "darwin":
		trash = "/.Trash/"
	default:
		return fmt.Errorf("err: %d, %s", http.StatusNotImplemented, "Unsupported OS")
	}
	if err = os.Rename(target, home+trash+target); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Succeeded")
	return nil
}
