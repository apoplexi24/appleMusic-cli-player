package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var replayCmd = &cobra.Command{
	Use:     "replay",
	Aliases: []string{"restart", "again"},
	Short:   "Restart the current track in Apple Music",
	Long: `Restart the current track in Apple Music from the beginning. This command will check 
if Apple Music is open, and if it is, it will restart the currently playing track.
After restarting, it will display the current song information.

Usage example:

music replay
This will restart the currently playing song from the beginning and display its information.`,
	Run: func(cmd *cobra.Command, args []string) {
		isOpen, err := isMusicOpen()
		if err != nil {
			fmt.Println("Error checking if Apple Music is open:", err)
			return
		}

		if !isOpen {
			fmt.Println("Apple Music is not open!")
			return
		}

		// Set player position to 0 to restart the song
		if _, err := mack.Tell("Music", "set player position to 0"); err != nil {
			fmt.Println("Error restarting the song:", err)
			return
		}

		info, err := getCurrentSongInfo()
		if err != nil {
			fmt.Printf("Error getting current song info: %v", err)
			return
		}

		if info.trackName == "" {
			fmt.Println("Song Restarted")
			return
		}
		fmt.Printf("Restarting: %s\nBy: %s\n", info.trackName, info.artistName)
	},
}

func init() {
	rootCmd.AddCommand(replayCmd)
}
