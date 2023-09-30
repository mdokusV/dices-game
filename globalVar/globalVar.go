package globalVar

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const NumberOfDices = 5

const NumberOfPlayers = 2

const IO_human = true

var RandomGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
