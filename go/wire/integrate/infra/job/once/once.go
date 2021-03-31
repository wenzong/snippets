package once

import (
	"context"

	"github.com/wenzong/demo/infra/job"
)

type Once struct {
	task job.Task

	interceptors []job.TaskInterceptor
}

func (o *Once) Run(ctx context.Context) error {
	handler := o.task

	for _, i := range o.interceptors {
		handler = i(handler)
	}

	return handler(ctx)
}

func NewOnce(t job.Task, is []job.TaskInterceptor) job.Job {
	return &Once{
		task:         t,
		interceptors: is,
	}
}
