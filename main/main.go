package main

import (
	"flag"
	"fmt"
	"time"

	figure "github.com/common-nighthawk/go-figure"
)

func main() {
	seconds := flag.Int("s", 0, "The seconds for the Pomodoro timer")
	minutes := flag.Int("m", 0, "The minutes for the Pomodoro timer")
	hours := flag.Int("h", 0, "The hours for the Pomodoro timer")

	flag.Parse()

	if *seconds == 0 && *minutes == 0 && *hours == 0 {
		fmt.Println("Please provide at least one flag for the timer")
		return
	}

	remainingSeconds := *seconds + (*minutes * 60) + (*hours * 3600)

	pomodoroTicker := time.NewTicker(1 * time.Second)

	for rs := remainingSeconds; rs >= 0; rs-- {
		<-pomodoroTicker.C
		clearScreen()
		printFigure(formatTime(rs%60, (rs/60)%60, (rs/3600)%60))
	}
	
	clearScreen()
	printFigure("Time's up!")
	pomodoroTicker.Stop()
}

func formatTime(s, m, h int) string {
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func printFigure(str string) {
	figure.NewFigure(str, "", true).Print()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}