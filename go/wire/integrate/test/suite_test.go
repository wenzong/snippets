package test

import (
	"context"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/wenzong/demo/test/helper"
	_ "github.com/wenzong/demo/test/user"
)

func TestSuite(t *testing.T) {
	app, cleanup := helper.App(t)
	defer cleanup()

	ctx, cancel := context.WithCancel(context.Background())
	go app.Run(ctx)
	defer cancel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Demo Test Suite")
}

var _ = ginkgo.BeforeSuite(func() {
})

var _ = ginkgo.AfterSuite(func() {
})
