package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-tenant-api/internal/flusher"
	"github.com/ozoncp/ocp-tenant-api/internal/mocks"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
)

var mockErr = errors.New("mock error")

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mockStorage *mocks.MockRepo
		tenants     []tenant.Tenant
		result      []tenant.Tenant
		f           flusher.Flusher
		sizeOfPkg   uint
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockStorage = mocks.NewMockRepo(ctrl)

		tenants = []tenant.Tenant{
			tenant.New(1, "name1", 1),
			tenant.New(2, "name2", 1),
			tenant.New(3, "name3", 1)}
		sizeOfPkg = 2
	})

	JustBeforeEach(func() {
		f = flusher.New(mockStorage, sizeOfPkg)
		result = f.Flush(tenants)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo is ok", func() {
		BeforeEach(func() {
			mockStorage.EXPECT().AddTenants(gomock.Any()).Return(nil).MinTimes(2)
		})

		It("repo AddTenants OK", func() {
			Expect(result).Should(BeNil())
		})
	})

	Context("repo is`not ok", func() {
		BeforeEach(func() {
			mockStorage.EXPECT().AddTenants(gomock.Any()).Return(mockErr)
		})

		It("repo AddTenants Fail", func() {
			Expect(len(result)).Should(BeEquivalentTo(len(tenants)))
			Expect(result).Should(BeEquivalentTo(tenants))
		})
	})
})
