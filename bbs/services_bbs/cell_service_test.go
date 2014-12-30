package services_bbs_test

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"
	"github.com/tedsuo/ifrit"

	. "github.com/cloudfoundry-incubator/runtime-schema/bbs/services_bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/storeadapter"
)

var _ = Describe("Cell Service Registry", func() {
	const interval = time.Second
	var (
		bbs                *ServicesBBS
		heartbeat1         ifrit.Process
		heartbeat2         ifrit.Process
		firstCellPresence  models.CellPresence
		secondCellPresence models.CellPresence
	)

	BeforeEach(func() {
		bbs = New(etcdClient, lagertest.NewTestLogger("test"))

		firstCellPresence = models.NewCellPresence("first-rep", "lucid64", "1.2.3.4", "the-zone")
		secondCellPresence = models.NewCellPresence("second-rep", ".Net", "4.5.6.7", "the-zone")

		heartbeat1 = ifrit.Invoke(bbs.NewCellHeartbeat(firstCellPresence, interval))
		heartbeat2 = ifrit.Invoke(bbs.NewCellHeartbeat(secondCellPresence, interval))
	})

	AfterEach(func() {
		heartbeat1.Signal(os.Interrupt)
		heartbeat2.Signal(os.Interrupt)
		Eventually(heartbeat1.Wait()).Should(Receive(BeNil()))
		Eventually(heartbeat2.Wait()).Should(Receive(BeNil()))
	})

	Describe("MaintainCellPresence", func() {
		It("should put /cell/CELL_ID in the store with a TTL", func() {
			node, err := etcdClient.Get("/v1/cell/" + firstCellPresence.CellID)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(node.TTL).ShouldNot(BeZero())

			expectedJSON, err := models.ToJSON(firstCellPresence)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(node.Value).Should(MatchJSON(expectedJSON))
		})
	})

	Describe("CellById", func() {
		Context("when the cell exists", func() {
			It("returns the correct CellPresence", func() {
				cellPresence, err := bbs.CellById(firstCellPresence.CellID)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cellPresence).Should(Equal(firstCellPresence))
			})
		})

		Context("when the cell does not exist", func() {})
	})

	Describe("Cells", func() {
		Context("when there are available Cells", func() {
			It("should get from /v1/cell/", func() {
				cellPresences, err := bbs.Cells()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(cellPresences).Should(HaveLen(2))
				Ω(cellPresences).Should(ContainElement(firstCellPresence))
				Ω(cellPresences).Should(ContainElement(secondCellPresence))
			})

			Context("when there is unparsable JSON in there...", func() {
				BeforeEach(func() {
					etcdClient.Create(storeadapter.StoreNode{
						Key:   shared.CellSchemaPath("blah"),
						Value: []byte("ß"),
					})
				})

				It("should ignore the unparsable JSON and move on", func() {
					cellPresences, err := bbs.Cells()
					Ω(err).ShouldNot(HaveOccurred())
					Ω(cellPresences).Should(HaveLen(2))
					Ω(cellPresences).Should(ContainElement(firstCellPresence))
					Ω(cellPresences).Should(ContainElement(secondCellPresence))
				})
			})
		})

		Context("when there are none", func() {
			BeforeEach(func() {
				heartbeat1.Signal(os.Interrupt)
				heartbeat2.Signal(os.Interrupt)
				Eventually(heartbeat1.Wait()).Should(Receive(BeNil()))
				Eventually(heartbeat2.Wait()).Should(Receive(BeNil()))
			})

			It("should return empty", func() {
				reps, err := bbs.Cells()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(reps).Should(BeEmpty())
			})
		})
	})
})
