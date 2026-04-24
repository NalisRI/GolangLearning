package address_api_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAddressApiTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AddressApiTests Suite")
}
