//go:build wireinject
// +build wireinject

package di

import (
	db "github.com/Crud-application/db"
	h "github.com/Crud-application/pkg/api/handlers"
	svcInter "github.com/Crud-application/pkg/application/services"
	uApp "github.com/Crud-application/pkg/application/user"
	repoInter "github.com/Crud-application/pkg/domain/persistence"
	uRepo "github.com/Crud-application/pkg/infrastructure/persistence/user"
	"github.com/google/uuid"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func provideMongoDBclient() *mongo.Client {
	client, _, _ := db.GetMongoDB() // Use your GetMongoDB method to retrieve the client
	return client
}

func provideUserRepository() *uRepo.MongoUserRepository {
	wire.Build(uRepo.NewMongoUserRepository, provideMongoDBclient)
	return nil
}

var userRepoSet = wire.NewSet(
	provideUserRepository,
	wire.Bind(new(repoInter.IUserRepository), new(*uRepo.MongoUserRepository)),
)

func provideUUIDGenerator() uApp.UUIDGenerator {
	return func() string {
		return uuid.New().String() // Using the uuid package
	}
}

func provideUserService() *uApp.UserService {
	wire.Build(
		uApp.NewUserService,
		userRepoSet, // Injects the user repository
		provideUUIDGenerator,
	)
	return nil
}

var userSvcSet = wire.NewSet(
	provideUserService,
	wire.Bind(new(svcInter.IUserService), new(*uApp.UserService)),
	provideUUIDGenerator,
)

func provideUserHandler() *h.UserHandler {
	wire.Build(h.NewUserHandler, userSvcSet)
	return nil
}

func InjectHandler() *h.Handlers {
	wire.Build(h.NewHandlers, provideUserHandler)
	return nil
}
