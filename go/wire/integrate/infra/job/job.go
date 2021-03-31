package job

import (
	"context"
)

type Task func(context.Context) error

type TaskInterceptor func(Task) Task

type Job interface {
	Run(context.Context) error
}
