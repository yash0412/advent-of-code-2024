package main

import (
	"adventofcode/daynine"
	"fmt"
	"time"
)

func main() {
	// dayone.Solve()
	// daytwo.Solve()
	// daytwo.Solve2()
	// daythree.Solve()
	// daythree.Solve2()
	// dayfour.Solve()
	// dayfour.Solve2()
	// dayfive.Solve()
	// daysix.Solve()
	// daysix.Solve2()
	// dayseven.Solve()
	// dayseven.Solve2()
	// dayeight.Solve()
	// dayeight.Solve2()
	// dayfourteen.Solve()
	// dayfourteen.Solve2()
	// daythirteen.Solve()
	// daythirteen.Solve2()
	// daytwelve.Solve()
	// daytwelve.Solve2()
	// dayfifteen.Solve()
	// dayfifteen.Solve2()
	// dayeleven.Solve()
	// dayeleven.Solve2()
	// dayten.Solve()
	// dayten.Solve2()
	startTime := time.Now()
	daynine.Solve()
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
	daynine.Solve2()
	elapsed = time.Since(startTime)
	fmt.Println(elapsed)
}
