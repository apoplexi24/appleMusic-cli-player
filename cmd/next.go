package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/andybrewer/mack"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Skip to the next track in Apple Music",
	Long: `Skip to the next track in Apple Music. This command will check if Apple Music is open,
and if it is, it will skip to the next track. After skipping, it will display the 
current song information.

Usage example:

music next
This will skip the currently playing song and display the new song's information.`,
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

		if _, err := mack.Tell("Music", "next track"); err != nil {
			fmt.Println("Error skipping the song:", err)
		}

		info, err := getCurrentSongInfo()
		if err != nil {
			fmt.Printf("Error getting current song info: %v", err)
			return
		}

		if info.trackName == "" {
			fmt.Println("Song Skipped")
			return
		}
		fmt.Printf("Now Playing: %s\nBy: %s\n", info.trackName, info.artistName)

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
	rootCmd.AddCommand(nextCmd)
}
