package main

import (
	"bufio"
	"os"
	"runtime/pprof"
	"time"

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

	//RNG setup
	globalVar.FastRandomGenerator.Seed(uint32(time.Now().UnixNano()))

	playerGroup := globalStructs.NewPlayerGroup(globalVar.NumberOfPlayers)
	for i := 0; i < 500_000; i++ {
		globalStructs.ResetPlayerGroup(playerGroup)
		//start game
		game.Start(playerGroup)
		// fmt.Println(playerGroup[0])
		// fmt.Println(playerGroup[1])
		// fmt.Println()

	}
}
