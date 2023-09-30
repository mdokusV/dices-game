package main

import (
	"bufio"
	"os"
	"runtime/pprof"

	"github.com/mdokusV/dices-game/game"
	globalStructs "github.com/mdokusV/dices-game/globalStructs"
	"github.com/mdokusV/dices-game/globalVar"
)

func main() {
	//profile game speed
	defer bufio.NewWriter(os.Stdout).Flush()
	f, _ := os.Create("cpu_profile.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	playerGroup := globalStructs.NewPlayerGroup(globalVar.NumberOfPlayers)
	for i := 0; i < 1_000_000; i++ {
		globalStructs.DeletePlayerGroup(playerGroup)
		//start game
		game.Start(playerGroup)
		// fmt.Println(playerGroup[0])
		// fmt.Println(playerGroup[1])
		// fmt.Println()

	}
}
