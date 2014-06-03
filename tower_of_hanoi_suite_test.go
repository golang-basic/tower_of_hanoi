package tower_of_hanoi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TowerOfHanoi Suite")
}
