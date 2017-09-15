package model

import (
	"errors"
	"reflect"
	"runtime"

	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
	"github.com/axsh/openvdc/model/backend"
	"golang.org/x/net/context"
)

var ErrBackendNotInContext = errors.New("Given context does not have the backend object")

func Connect(ctx context.Context, dest backend.ConnectionAddress) (context.Context, error) {
	bk := backend.NewZkBackend()
	err := bk.Connect(dest)
	if err != nil {
		return nil, err
	}
	return withBackendCtx(ctx, bk), nil
}

func Close(ctx context.Context) error {
	bk := GetBackendCtx(ctx)
	if bk == nil {
		return ErrBackendNotInContext
	}
	return bk.Close()
}

var schemaKeys []string

func InstallSchemas(bk backend.ModelSchema) error {
	schema := bk.Schema()
	return schema.Install(schemaKeys)
}

type ctxKey string

const ctxBackendKey ctxKey = "model.backend"

func withBackendCtx(ctx context.Context, bk backend.ModelBackend) context.Context {
	return context.WithValue(ctx, ctxBackendKey, bk)
}

func GetBackendCtx(ctx context.Context) backend.ModelBackend {
	if ctx == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Fatalf("GetBackendCtx() does not accept nil.: %s:%d", file, line)
	}
	bk, ok := ctx.Value(ctxBackendKey).(backend.ModelBackend)
	// Assert returned type from ctx.
	if !ok && bk != nil {
		log.Fatalf("Unexpected type to '%s' context value: %v", ctxBackendKey, reflect.TypeOf(bk))
	}
	return bk
}

func GrpcInterceptor(modelAddr backend.ConnectionAddress, clusterCtx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, err := Connect(ctx, modelAddr)
		if err != nil {
			log.WithError(err).Errorf("Failed to connect to model backend: %s", modelAddr)
			return nil, err
		}
		defer func() {
			err := Close(ctx)
			if err != nil {
				log.WithError(err).Error("Failed to close connection to model backend.")
			}
		}()
		ctx = withClusterBackendCtx(ctx, GetClusterBackendCtx(clusterCtx))
		return handler(ctx, req)
	}
}

// https://gist.github.com/shaxbee/a87e2c028a21c60e5aace593a23b27a1
type serverStreamWithContext struct {
	grpc.ServerStream
	ctx context.Context
}

func (ss *serverStreamWithContext) Context() context.Context {
	return ss.ctx
}

func GrpcStreamInterceptor(modelAddr backend.ConnectionAddress, clusterCtx context.Context) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := Connect(ss.Context(), modelAddr)
		if err != nil {
			log.WithError(err).Errorf("Failed to connect to model backend: %s", modelAddr)
			return err
		}
		defer func() {
			err := Close(ctx)
			if err != nil {
				log.WithError(err).Error("Failed to close connection to model backend.")
			}
		}()
		ctx = withClusterBackendCtx(ctx, GetClusterBackendCtx(clusterCtx))
		return handler(srv, &serverStreamWithContext{ss, ctx})
	}
}

const ctxClusterBackendKey ctxKey = "cluster.backend"

func withClusterBackendCtx(ctx context.Context, bk backend.ClusterBackend) context.Context {
	return context.WithValue(ctx, ctxClusterBackendKey, bk)
}

func WithMockClusterBackendCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxClusterBackendKey, &backend.MockClusterBackend{})
}

func ClusterConnect(ctx context.Context, dest backend.ConnectionAddress) (context.Context, error) {
	bk := backend.NewZkClusterBackend()
	err := bk.Connect(dest)
	if err != nil {
		return nil, err
	}
	return withClusterBackendCtx(ctx, bk), nil
}

func ClusterClose(ctx context.Context) error {
	bk := GetClusterBackendCtx(ctx)
	if bk == nil {
		return ErrBackendNotInContext
	}
	return bk.Close()
}

func GetClusterBackendCtx(ctx context.Context) backend.ClusterBackend {
	if ctx == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Fatalf("BUGON: GetClusterBackendCtx() does not accept nil.: %s:%d", file, line)
	}
	bk, ok := ctx.Value(ctxClusterBackendKey).(backend.ClusterBackend)
	// Assert returned type from ctx.
	if !ok && bk != nil {
		log.Fatalf("BUGON: Unexpected type to '%s' context value: %v", ctxClusterBackendKey, reflect.TypeOf(bk))
	}
	return bk
}
