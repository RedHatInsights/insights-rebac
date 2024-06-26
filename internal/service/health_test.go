package service

import (
	"context"
	"testing"

	pb "github.com/project-kessel/relations-api/api/health/v1"
	"github.com/project-kessel/relations-api/internal/biz"
	"github.com/project-kessel/relations-api/internal/data"

	"github.com/stretchr/testify/assert"
)

func TestHealthService_GetLivez(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	spicedb, err := container.CreateSpiceDbRepository()
	assert.NoError(t, err)

	service := createHealthService(spicedb)
	resp, err := service.GetLivez(ctx, &pb.GetLivezRequest{})

	assert.NoError(t, err)
	assert.Equal(t, resp, &pb.GetLivezReply{Status: "OK", Code: 200})
}

func TestHealthService_GetReadyz_SpiceDBAvailable(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	spicedb, err := container.CreateSpiceDbRepository()
	assert.NoError(t, err)

	service := createHealthService(spicedb)
	resp, err := service.GetReadyz(ctx, &pb.GetReadyzRequest{})

	assert.NoError(t, err)
	assert.Equal(t, resp, &pb.GetReadyzReply{Status: "OK", Code: 200})
}

func createHealthService(spicedb *data.SpiceDbRepository) *HealthService {
	return NewHealthService(biz.NewIsBackendAvailableUsecase(spicedb))
}