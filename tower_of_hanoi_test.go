package tower_of_hanoi_test

import (
	. "github.com/gobasic/game/tower_of_hanoi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TowerOfHanoi", func() {
	Describe("Initialization", func() {
		Context("fills up SourcePeg with decreasing order of disks ", func() {
			NewTowerOfHanoi(3)

			Ω(*(SourcePeg)).To(HaveLen(3))
			Ω(*(TargetPeg)).To(HaveLen(0))
			Ω(*(HelperPeg)).To(HaveLen(0))
		})
	})
	Describe("Simulation", func(){
		Context("playes the tower of hanoi game", func() {
			NewTowerOfHanoi(16)
			SimulateGame()
		})
	})
})
