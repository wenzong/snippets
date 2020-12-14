package user

import (
	"context"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/wenzong/demo/api/pb"
	"github.com/wenzong/demo/test/helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("User gRPC", func() {
	var (
		client pb.UserServiceClient
	)

	BeforeEach(func() {
		client = pb.NewUserServiceClient(helper.ClientConn())
	})

	AfterEach(func() {
	})

	It("gRPC 请求 UserService.Get", func() {
		_, err := client.Get(context.Background(), &pb.GetRequest{UserId: 1})
		if assert.NotNil(GinkgoT(), err) {
			assert.EqualValues(GinkgoT(), status.Convert(err).Code(), codes.NotFound)
		}
	})
})
