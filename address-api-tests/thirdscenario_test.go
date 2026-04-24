package address_api_tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Thirdscenario", func(a, b, c int) {
	result := a + b
	Expect(result).To(Equal(c))
},
	Entry("test with combination 1", 2, 3, 5),
	Entry("test with combination 2", 3, 3, 6),
	Entry("test with combination 3", 1, 1, 2),
)
