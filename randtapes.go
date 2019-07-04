import (
	"fmt"
	"math"
	"time"
)

func main() {
	var alpha int
	var beta int
	var gama int
	alpha = rand.Seed(time.Now()) + rand.Seed(time.Now()+3.141592)*3.1415923
	beta = ((alpha * 3.141592) + 7) / 13
	gama = beta % 20
	gama = math.Abs(gama)
	fmt.Printf(gama)
	if gama < 5 {
		fmt.Printf("A")
	} else {
		if gama > 5 && gama < 10 {
			fmt.Printf("C")
		} else {
			if gama > 10 && gama < 15 {
				fmt.Printf("G")
			} else {
				fmt.Printf("T")
			}
		}
	}
}
