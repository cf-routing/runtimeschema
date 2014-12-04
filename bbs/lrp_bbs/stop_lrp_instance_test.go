package lrp_bbs_test

import (
	"errors"

	"github.com/cloudfoundry-incubator/runtime-schema/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StopLRPInstance", func() {
	var actualLRP models.ActualLRP
	var cellPresence models.CellPresence

	BeforeEach(func() {
		cellPresence = models.CellPresence{
			CellID:     "the-cell-id",
			Stack:      "the-stack",
			RepAddress: "cell.example.com",
		}

		registerCell(cellPresence)

		_, alrp, err := createAndClaim(models.NewActualLRP("some-process-guid", "some-instance-guid", cellPresence.CellID, "domain", 5678, models.ActualLRPStateClaimed))
		Ω(err).ShouldNot(HaveOccurred())
		actualLRP = *alrp
	})

	Describe("RequestStopLRPInstance", func() {
		Context("When the request is successful", func() {
			It("makes a stop instance request to the correct cell", func() {
				err := bbs.RequestStopLRPInstance(actualLRP)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(fakeCellClient.StopLRPInstanceCallCount()).Should(Equal(1))

				addr1, stop1 := fakeCellClient.StopLRPInstanceArgsForCall(0)
				Ω(addr1).Should(Equal(cellPresence.RepAddress))
				Ω(stop1).Should(Equal(actualLRP))
			})
		})

		Context("When the cell returns an error", func() {
			var expectedError = errors.New("cell go boom")
			BeforeEach(func() {
				fakeCellClient.StopLRPInstanceReturns(expectedError)
			})

			It("returns the error", func() {
				err := bbs.RequestStopLRPInstance(actualLRP)
				Ω(err).Should(Equal(expectedError))
			})
		})

		Context("when the store is out of commission", func() {
			itRetriesUntilStoreComesBack(func() error {
				return bbs.RequestStopLRPInstance(actualLRP)
			})
		})
	})

	Describe("RequestStopLRPInstances", func() {
		var anotherActualLRP models.ActualLRP

		BeforeEach(func() {
			_, alrp, err := createAndClaim(models.NewActualLRP(
				"some-other-process-guid", "some-other-instance-guid", cellPresence.CellID,
				"domain", 1234, models.ActualLRPStateClaimed))
			Ω(err).ShouldNot(HaveOccurred())
			anotherActualLRP = *alrp
		})

		It("stops the LRP instances on the correct cell", func() {
			err := bbs.RequestStopLRPInstances([]models.ActualLRP{actualLRP, anotherActualLRP})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(fakeCellClient.StopLRPInstanceCallCount()).Should(Equal(2))

			addr1, stop1 := fakeCellClient.StopLRPInstanceArgsForCall(0)
			Ω(addr1).Should(Equal(cellPresence.RepAddress))
			Ω(stop1).Should(Equal(actualLRP))

			addr2, stop2 := fakeCellClient.StopLRPInstanceArgsForCall(1)
			Ω(addr2).Should(Equal(cellPresence.RepAddress))
			Ω(stop2).Should(Equal(anotherActualLRP))
		})

		Context("when the store is out of commission", func() {
			itRetriesUntilStoreComesBack(func() error {
				return bbs.RequestStopLRPInstances([]models.ActualLRP{actualLRP})
			})
		})
	})
})
