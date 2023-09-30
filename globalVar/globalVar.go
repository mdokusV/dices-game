package globalVar

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/valyala/fastrand"
)

const NumberOfDices = 5

const NumberOfPlayers = 2

const NumberOfStates = 11

const IO_human = false

var RandomGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var FastRandomGenerator fastrand.RNG

var Reader *bufio.Reader = bufio.NewReader(os.Stdin)
