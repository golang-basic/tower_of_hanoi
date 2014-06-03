package tower_of_hanoi

import (
	"fmt"
	"strconv"
	"math"
)

var SourcePeg *[]int
var TargetPeg *[]int
var HelperPeg *[]int

func NewTowerOfHanoi(numOfDisks int) {
	SourcePeg = getNsortedNum(numOfDisks)
	TargetPeg = &([]int {})
	HelperPeg = &([]int {})
	return
}

func getNsortedNum(num int) (*[]int) {
	ser := []int {num}
	for i := 1; i < num; i++ {
		ser = append(ser, (num-i))
	}
	return &ser
}

func Print() {
	fmt.Print("Source >>>")
	fmt.Println(*(SourcePeg))
	fmt.Print("Helper >>>")
	fmt.Println(*(HelperPeg))
	fmt.Print("Target >>>")
	fmt.Println(*(TargetPeg))
}
var recur float64

func SimulateGame() {
	recur = math.Exp2(float64(len(*(SourcePeg)))) -1
	fmt.Println("Order of the algorithm" + strconv.Itoa(int(recur)))
	hanoi(len(*(SourcePeg)), SourcePeg, HelperPeg, TargetPeg)
}


func hanoi(n int, source *[]int, helper *[]int, target *[]int) {
	if n > 0 {
		recur = recur -1
		fmt.Println("**************" + strconv.Itoa(int(recur)))

		//		# move tower of size n - 1 to helper:
		//		hanoi(n - 1, source, target, helper)
		//		# move disk from source peg to target peg
		//		if source[0]:
		//		disk = source[0].pop()
		//		print "moving " + str(disk) + " from " + source[1] + " to " + target[1]
		//		target[0].append(disk)
		//		# move tower of size n-1 from helper to target
		//		hanoi(n - 1, helper, source, target)
		hanoi(n-1, source, target, helper)
		if len(*source) > 0 {
			disk := (*source)[len(*source) - 1]
			*source = (*source)[:(len(*source)-1)]
			*target = append((*target), disk)
			fmt.Println("target")
			fmt.Println(*target)
		}
		hanoi(n-1, helper, source, target)
	}
}
