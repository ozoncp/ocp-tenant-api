package flusher_test

import (
	"context"
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
		sizeOfPkg   uint64
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
		result = f.Flush(context.TODO(), tenants)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo is ok", func() {
		BeforeEach(func() {
			mockStorage.EXPECT().AddTenants(context.TODO(), gomock.Any()).Return(sizeOfPkg, nil).MinTimes(2)
		})

		It("repo AddTenants OK", func() {
			Expect(result).Should(BeNil())
		})
	})

	Context("repo is`not ok", func() {
		BeforeEach(func() {
			mockStorage.EXPECT().AddTenants(context.TODO(), gomock.Any()).Return(uint64(0), mockErr)
		})

		It("repo AddTenants Fail", func() {
			Expect(len(result)).Should(BeEquivalentTo(len(tenants)))
			Expect(result).Should(BeEquivalentTo(tenants))
		})
	})

	Context("repo runs from time to time", func() {
		BeforeEach(func() {
			mockStorage.EXPECT().AddTenants(context.TODO(), gomock.Any()).Return(sizeOfPkg, nil).Times(1)
			mockStorage.EXPECT().AddTenants(context.TODO(), gomock.Any()).Return(uint64(0), mockErr)
		})

		It("repo AddTenants Ok and Fail", func() {
			Expect(result).Should(BeEquivalentTo(tenants[sizeOfPkg:]))
		})
	})
})
