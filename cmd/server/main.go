package server

import (
	"diffme.dev/diffme-api/server/core/interfaces"
	"diffme.dev/diffme-api/server/core/middleware"
	ApiKeyDomain "diffme.dev/diffme-api/server/modules/api-keys"
	ApiKeysHTTP "diffme.dev/diffme-api/server/modules/api-keys/surfaces/http"
	AuthDomain "diffme.dev/diffme-api/server/modules/authentication"
	AuthenticationHTTP "diffme.dev/diffme-api/server/modules/authentication/surfaces/http"
	ChangeDomain "diffme.dev/diffme-api/server/modules/changes"
	EventHTTP "diffme.dev/diffme-api/server/modules/changes/surfaces/http"
	OrgDomain "diffme.dev/diffme-api/server/modules/organizations"
	OrganizationHTTP "diffme.dev/diffme-api/server/modules/organizations/surfaces/http"
	SnapshotDomain "diffme.dev/diffme-api/server/modules/snapshots"
	SnapshotHTTP "diffme.dev/diffme-api/server/modules/snapshots/surfaces/http"
	TeamMembersHTTP "diffme.dev/diffme-api/server/modules/team-members/surfaces/http"
	UserDomain "diffme.dev/diffme-api/server/modules/users"
	UserHTTP "diffme.dev/diffme-api/server/modules/users/surfaces/http"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	FiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type ServerDependencies struct {
	changeUseCases   ChangeDomain.ChangeUseCases
	orgUseCases      OrgDomain.OrganizationUseCases
	snapshotRepo     SnapshotDomain.SnapshotRepo
	userRepo         UserDomain.UserRepository
	userUseCases     UserDomain.UserUseCases
	snapshotUseCases SnapshotDomain.SnapshotUseCases
	searchChangeRepo ChangeDomain.SearchChangeRepository
	authUseCases     AuthDomain.UseCases
	producer         *kafka.Producer
	authProvider     interfaces.AuthProvider
	apiKeyRepo       ApiKeyDomain.ApiKeyRepository
}

func NewServerDependencies(
	changeUseCases ChangeDomain.ChangeUseCases,
	orgUseCases OrgDomain.OrganizationUseCases,
	snapshotRepo SnapshotDomain.SnapshotRepo,
	snapshotUseCases SnapshotDomain.SnapshotUseCases,
	searchChangeRepo ChangeDomain.SearchChangeRepository,
	userRepo UserDomain.UserRepository,
	userUseCases UserDomain.UserUseCases,
	authUseCases AuthDomain.UseCases,
	producer *kafka.Producer,
	authProvider interfaces.AuthProvider,
	apiKeyRepo ApiKeyDomain.ApiKeyRepository,
) ServerDependencies {
	return ServerDependencies{
		changeUseCases:   changeUseCases,
		orgUseCases:      orgUseCases,
		userRepo:         userRepo,
		userUseCases:     userUseCases,
		snapshotRepo:     snapshotRepo,
		snapshotUseCases: snapshotUseCases,
		searchChangeRepo: searchChangeRepo,
		producer:         producer,
		authProvider:     authProvider,
		authUseCases:     authUseCases,
		apiKeyRepo:       apiKeyRepo,
	}
}

func StartServer(deps ServerDependencies) {
	println("[starting server]")

	// Fiber instance
	app := fiber.New()

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New())

	// compress
	app.Use(compress.New())

	// cache TODO:
	//app.Use(cache.New())

	// request ID
	app.Use(requestid.New())

	// so that panics don't crash the server
	app.Use(FiberRecover.New())

	app.Use(middleware.AuthMiddleware(deps.authProvider, deps.userRepo, deps.apiKeyRepo))

	v1 := app.Group("/v1")

	addRoutes(v1, deps)

	println("server listening on :3001")

	// start server
	app.Listen(":3001")
	//defer func() {
	//	if err = mongoClient.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}

func addRoutes(app fiber.Router, deps ServerDependencies) {
	ApiKeysHTTP.ApiKeyRoutes(app, nil) // TODO:
	TeamMembersHTTP.TeamMemberRoutes(app)
	AuthenticationHTTP.AuthenticationRoutes(app, deps.authUseCases)
	UserHTTP.UserRoutes(app, deps.userRepo, deps.userUseCases, deps.authProvider)
	EventHTTP.ChangeRoutes(app, deps.changeUseCases)
	OrganizationHTTP.OrgRoutes(app, deps.orgUseCases)
	SnapshotHTTP.SnapshotRoutes(app, deps.snapshotRepo, deps.snapshotUseCases)
}
