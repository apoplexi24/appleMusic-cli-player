package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/andybrewer/mack"
	"github.com/qeesung/image2ascii/convert"
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

		// Run the AppleScript to get the album art
		newCmd := exec.Command("osascript", "scripts/get_album_art.scpt")
		err = newCmd.Run()
		if err != nil {
			log.Fatalf("Failed to run AppleScript: %v", err)
		}

		// Convert the image to ASCII
		convertOptions := convert.DefaultOptions
		converter := convert.NewImageConverter()
		convertOptions.FixedWidth = 40
		convertOptions.FixedHeight = 20
		asciiArt := converter.ImageFile2ASCIIString("scripts/tmp.jpg", &convertOptions)

		// Print the ASCII art
		fmt.Println(asciiArt)
	},
}

func init() {
	rootCmd.AddCommand(replayCmd)
}
