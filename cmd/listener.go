package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/idprm/go-three-direct/internal/domain/repository"
	"github.com/idprm/go-three-direct/internal/handler"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/gorm"
	loggerDb "gorm.io/gorm/logger"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Listener Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		l := logger.NewLogger()

		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MO_EXCHANGE, true, RMQ_MO_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_DR_EXCHANGE, true, RMQ_DR_QUEUE)

		r := routeListenerUrl(db, rds, rmq, l)

		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeListenerUrl(db *gorm.DB, rds *redis.Client, rmq rmqp.AMQP, l *logger.Logger) *fiber.App {
	engine := html.New("./src/presenter/views", ".html")

	/**
	 * Init Fiber
	 */
	r := fiber.New(fiber.Config{
		Views: engine,
	})

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	/**
	 * Access log on browser
	 */
	r.Use("/logs", filesystem.New(filesystem.Config{
		Root:         http.Dir(LOG_PATH),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))

	r.Static("/static", path+"/public")

	// Default config
	r.Use(cors.New())

	// Config for customization
	r.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	blacklistRepo := repository.NewBlacklistRepository(db)
	blacklistService := services.NewBlacklistService(blacklistRepo)

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)

	contentRepo := repository.NewContentRepository(db)
	contentService := services.NewContentService(contentRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)

	h := handler.NewIncomingHandler(
		rds,
		rmq,
		l,
		blacklistService,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
	)

	r.Get("/moh3i", h.MobileOriginated)
	r.Get("/camph3i", h.MobileOriginated)
	r.Get("/drh3i", h.DeliveryReport)

	/**
	 * Reports
	 */
	// report := r.Group("report")
	// report.Get("mo", handlerIncoming.ReportMO)
	// report.Get("renewal", handlerIncoming.ReportRenewal)
	// report.Get("firstpush", handlerIncoming.ReportFirstpush)

	/**
	 * Landing Page
	 */
	// r.Get("gamren", handlerIncoming.GamrenIndex)
	// r.Get("gamren/term", handlerIncoming.GamrenTerm)

	return r
}
