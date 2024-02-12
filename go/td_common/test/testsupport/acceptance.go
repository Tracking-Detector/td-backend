package testsupport

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport/containers"
)

type AcceptanceTest struct {
	Ctx      context.Context
	cancel   context.CancelFunc
	mongo    *containers.MongoDBContainer
	minio    *containers.MinIOContainer
	rabbitmq *containers.RabbitMQContainer
}

func (t *AcceptanceTest) SetupIntegration() {
	t.Ctx, t.cancel = context.WithCancel(context.Background())
	t.mongo, _ = containers.NewMongoContainer(t.Ctx)
	t.minio, _ = containers.NewMinIOContainer(t.Ctx)
	t.rabbitmq, _ = containers.NewRabbitMQContainer(t.Ctx)
	t.mongo.Setenvs()
	t.minio.Setenvs()
	t.rabbitmq.Setenvs()
}

func (t *AcceptanceTest) TeardownIntegration() {
	t.mongo.Terminate(t.Ctx)
	t.minio.Terminate(t.Ctx)
	t.rabbitmq.Terminate(t.Ctx)
	t.cancel()
}
