package Defer

import (
	"fmt"
	"math/rand"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func ExampleDefer() {
	nums := []int{rand.Intn(999), rand.Intn(999), rand.Intn(999), rand.Intn(999), rand.Intn(999)}
	largest(nums)
}
