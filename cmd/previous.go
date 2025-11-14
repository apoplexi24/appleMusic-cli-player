package cmd

import (
	"fmt"
	"time"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var previousCmd = &cobra.Command{
	Use:     "previous",
	Aliases: []string{"prev", "back"},
	Short:   "Go back to the previous track in Apple Music",
	Long: `Go back to the previous track in Apple Music. This command will check if Apple Music is open,
and if it is, it will skip to the previous track. After skipping, it will display the 
current song information.

Usage example:

music previous
This will go back to the actual previous song and display its information.`,
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

		// First, restart current song (set position to 0)
		if _, err := mack.Tell("Music", "set player position to 0"); err != nil {
			fmt.Println("Error setting player position:", err)
			return
		}

		// Small delay
		time.Sleep(100 * time.Millisecond)

		// Now go to previous track
		if _, err := mack.Tell("Music", "previous track"); err != nil {
			fmt.Println("Error going to previous song:", err)
			return
		}

		time.Sleep(500 * time.Millisecond)

		info, err := getCurrentSongInfo()
		if err != nil {
			fmt.Printf("Error getting current song info: %v", err)
			return
		}

		if info.trackName == "" {
			fmt.Println("Went to Previous Song")
			return
		}
		fmt.Printf("Now Playing: %s\nBy: %s\n", info.trackName, info.artistName)
	},
}

func init() {
	rootCmd.AddCommand(previousCmd)
}
