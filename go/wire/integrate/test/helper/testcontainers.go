package helper

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	RedisContainerReq = testcontainers.ContainerRequest{
		Image:        "redis:6-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	MySQLContainerReq = testcontainers.ContainerRequest{
		Image:        "mysql:5.6",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("port: 3306  MySQL Community Server").WithStartupTimeout(300 * time.Second),
		Cmd:          []string{"mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"},
		Env: map[string]string{
			"MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
			"MYSQL_ROOT_PASSWORD":        "pass",
			"TZ":                         "Asia/Shanghai",
		},
		Tmpfs: map[string]string{
			"/var/lib/mysql": "rw",
		},
	}

	RabbitMQContainerReq = testcontainers.ContainerRequest{
		Image:        "rabbitmq:3.8.9-alpine",
		ExposedPorts: []string{"15672/tcp", "5672/tcp"},
		WaitingFor:   wait.ForListeningPort("5672/tcp").WithStartupTimeout(3 * time.Minute),
		Env: map[string]string{
			"RABBITMQ_ERLANG_COOKIE": "abcdefghijkl",
			"RABBITMQ_DEFAULT_USER":  "user",
			"RABBITMQ_DEFAULT_PASS":  "pass",
		},
	}
)

func Container(ctx context.Context, req testcontainers.ContainerRequest) testcontainers.Container {
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	return c
}
