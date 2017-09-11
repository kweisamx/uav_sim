package Strategy
import (
	"math/rand"
	"time"
)
func RandomStrategy()string{
	var Strategy = [...]string{"STILL","FORWARD","BACKWARD","RIGHT","LEFT"}
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return Strategy[rand.Intn(len(Strategy))]
}