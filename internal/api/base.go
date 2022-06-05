package api

import (
	"context"
	"net/http"
	"omshub/core-api/internal/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

type Server struct {
	config     Config
	router     *gin.Engine
	httpServer *http.Server
}

// Dependencies provides a crude means of dependency injection.
type Dependencies struct {
	NewRelicApp *newrelic.Application // All operations on this are nil-safe. No mocks required for testing.
	DB          *gorm.DB
}

func NewServer(config Config, deps Dependencies) *Server {
	router := gin.Default()

	// - All origin allowed
	// - GET, POST, PUT, PATCH, DELETE, HEAD methods
	// - Credentials share enabled
	// - Preflight requests cached for 12 hours
	configCORS := cors.DefaultConfig()
	configCORS.AllowAllOrigins = true
	configCORS.AllowCredentials = true
	router.Use(cors.New(configCORS))

	if deps.NewRelicApp != nil {
		router.Use(nrgin.Middleware(deps.NewRelicApp))
	}

	if deps.DB != nil {
		v1 := router.Group("/api/v1")
		{
			reviews := v1.Group("/reviews")
			reviews.GET("", handlers.NewGetAllReviewsHandler(deps.DB))
			reviews.GET("/:id", handlers.NewGetOneReviewHandler(deps.DB))
			// reviews.POST("", handlers.NewAddReviewHandler(deps.DB))
			// reviews.PUT("/:id", handlers.NewUpdateReviewHandler(deps.DB))
			// reviews.DELETE(":id", handlers.NewDeleteReviewHandler(deps.DB))

			users := v1.Group("/users")
			users.GET("", handlers.NewGetAllUsersHandler(deps.DB))
			users.GET("/:id", handlers.NewGetOneUserHandler(deps.DB))
			// users.POST("", handlers.NewAddUserHandler(deps.DB))
			// users.PUT("/:id", handlers.NewUpdateUserHandler(deps.DB))
			// users.DELETE(":id", handlers.NewDeleteUserHandler(deps.DB))

			courses := v1.Group("/courses")
			courses.GET("", handlers.NewGetAllCoursesHandler(deps.DB))
			courses.GET("/:id", handlers.NewGetOneCourseHandler(deps.DB))
			courses.GET("/:id/reviews", handlers.NewGetAllCourseReviewsHandler(deps.DB))
			courses.GET("/stats", handlers.NewGetAllCourseStatReviewsHandler(deps.DB))
			// courses.POST("", handlers.NewAddCourseHandler(deps.DB))
			// courses.PUT("/:id", handlers.NewUpdateCourseHandler(deps.DB))
			// courses.DELETE(":id", handlers.NewDeleteCourseHandler(deps.DB))

			semesters := v1.Group("/semesters")
			semesters.GET("", handlers.NewGetAllSemestersHandler(deps.DB))
			semesters.GET("/:id", handlers.NewGetOneSemesterHandler(deps.DB))
			// semesters.POST("", handlers.NewAddSemesterHandler(deps.DB))
			// semesters.PUT("/:id", handlers.NewUpdateSemesterHandler(deps.DB))
			// semesters.DELETE(":id", handlers.NewDeleteSemesterHandler(deps.DB))

			specializations := v1.Group("/specializations")
			specializations.GET("", handlers.NewGetAllSpecializationsHandler(deps.DB))
			specializations.GET("/:id", handlers.NewGetOneSpecializationHandler(deps.DB))
			// specializations.POST("", handlers.NewAddSpecializationHandler(deps.DB))
			// specializations.PUT("/:id", handlers.NewUpdateSpecializationHandler(deps.DB))
			// specializations.DELETE(":id", handlers.NewDeleteSpecializationHandler(deps.DB))

			programs := v1.Group("/programs")
			programs.GET("", handlers.NewGetAllProgramsHandler(deps.DB))
			programs.GET("/:id", handlers.NewGetOneProgramHandler(deps.DB))
			// programs.POST("", handlers.NewAddProgramHandler(deps.DB))
			// programs.PUT("/:id", handlers.NewUpdateProgramHandler(deps.DB))
			// programs.DELETE(":id", handlers.NewDeleteProgramHandler(deps.DB))
		}
	}

	httpServer := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	server := &Server{
		config:     config,
		router:     router,
		httpServer: httpServer,
	}

	router.GET("/", server.Index)
	router.GET("/ping", server.Ping)

	return server
}

func (s *Server) Serve() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Index(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
