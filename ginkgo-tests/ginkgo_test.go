package ginkgo

import (
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDebugEndpoint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testkube Debug Suite")

	g := NewGomegaWithT(t)
	Describe("Debug endpoints", func() {
		Context("gives information", func() {
			debug, err := queryDebugEndpoint()
			g.Expect(err).To(BeNil())
			It("contains clientVersion", func() {
				g.Expect(debug).To(ContainSubstring("clientVersion"))
			})
			It("contains serverVersion", func() {
				g.Expect(debug).To(ContainSubstring("serverVersion"))
			})
			It("contains apiLogs", func() {
				g.Expect(debug).To(ContainSubstring("apiLogs"))
			})
			It("contains operatorLogs", func() {
				g.Expect(debug).To(ContainSubstring("operatorLogs"))
			})
			It("contains executionLogs", func() {
				g.Expect(debug).To(ContainSubstring("executionLogs"))
			})
		})
	})
}

func queryDebugEndpoint() (string, error) {
	resp, err := http.Get("https://demo.testkube.io/results/v1/debug")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
