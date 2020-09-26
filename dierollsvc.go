package dierollsvc

import (
	"fmt"

	"github.com/tsasser05/dieroll"
)

func main() {
	fmt.Println("dieroll service!")
	stat := dieroll.RollStat()
	fmt.Printf("stat = %d", stat)

}
