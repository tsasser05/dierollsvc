package dierollsvc

import (
	"fmt"

	"github.com/tsasser05/dieroll"
)

type Roll struct {
	Roll int `json:"roll"`
}

// Dieroll rolls 4d6 for a stat
func Example() {
	dieroll.Init()
	fmt.Println("dieroll 4d6 for a stat!")
	stat := dieroll.RollStat()
	fmt.Printf("stat = %d", stat)

}

func Init() {
	dieroll.Init()
}

/* Endpoints:
 *
 *  Roll(max int, num int) (returns either total of 1 die or total of all dice rolled)
 *	RollStat() int
 *	RollDice() []int (returns a list containing the result for each roll)
 *	RandomKey(m map[string]string) string  (returns a random key in a map)
 *
 */
