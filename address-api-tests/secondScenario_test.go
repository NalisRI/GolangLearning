package address_api_tests_test

import (
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecondScenario", func() {

	BeforeEach(func() {
		log.Println("we are inside before each")
	})

	Context("When ", func() {

		It("addition should be as expected", func() {
			a := 2 + 3
			Expect(a).To(Equal(5))
		})

		It("multiplication should be as expected", func() {
			m := 3 * 3
			Expect(m).To(Equal(9))
		})

	})

	AfterEach(func() {
		log.Println("we are after each")
	})

})
