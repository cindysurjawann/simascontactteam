package api

import (
	"github.com/cindysurjawann/simascontactteam/asuransi"
	"github.com/cindysurjawann/simascontactteam/auth"
	"github.com/cindysurjawann/simascontactteam/custom"
	"github.com/cindysurjawann/simascontactteam/infoPromo"
	"github.com/cindysurjawann/simascontactteam/managelink"
	"github.com/cindysurjawann/simascontactteam/manageuser"
	"github.com/cindysurjawann/simascontactteam/zoomhistory"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	s.Router.POST("/create-account", authHandler.CreateUser)
	s.Router.POST("/login", authHandler.Login)
	s.Router.POST("/send-otp", authHandler.SendOTP)

	middleware := custom.MiddleWare{}
	authRoute := s.Router.Group("")
	authRoute.Use(middleware.Auth)
	authRoute.POST("/updatelastlogin", authHandler.UpdateLastLogin)

	manageLinkRepo := managelink.NewRepository(s.DB)
	manageLinkService := managelink.NewService(manageLinkRepo)
	manageLinkHandler := managelink.NewHandler(manageLinkService)

	csRoute := authRoute.Group("")
	csRoute.Use(middleware.IsCS)
	csRoute.PUT("/updatelink", manageLinkHandler.UpdateLink)
	csRoute.GET("/getlink", manageLinkHandler.GetLink)

	s.Router.GET("/get-link/:type", manageLinkHandler.GetLinkRequest)
	zoomHistoryRepo := zoomhistory.NewRepository(s.DB)
	zoomHistoryService := zoomhistory.NewService(zoomHistoryRepo)
	zoomHistoryHandler := zoomhistory.NewHandler(zoomHistoryService)
	s.Router.POST("/createzoomhistory", zoomHistoryHandler.CreateZoom)
	csRoute.GET("/getzoomhistory", zoomHistoryHandler.GetRiwayat)

	infoPromoRepo := infoPromo.NewRepository(s.DB)
	infoPromoService := infoPromo.NewService(infoPromoRepo)
	infoPromoHandler := infoPromo.NewHandler(infoPromoService)
	csRoute.GET("/getpromos", infoPromoHandler.GetInfos)
	s.Router.GET("/getrecentpromos", infoPromoHandler.GetRecentInfos)
	s.Router.POST("/postinfopromo", infoPromoHandler.AddInfo)

	adminRoute := authRoute.Group("")
	adminRoute.Use(middleware.IsAdmin)
	asuransiRepo := asuransi.NewRepository(s.DB)
	asuransiService := asuransi.NewService(asuransiRepo)
	asuransiHandler := asuransi.NewHandler(asuransiService)

	s.Router.GET("/getasuransi", asuransiHandler.GetAsuransi)
	adminRoute.POST("/postasuransi", asuransiHandler.CreateAsuransi)

	manageUserRepo := manageuser.NewRepository(s.DB)
	manageUserService := manageuser.NewService(manageUserRepo)
	manageUserHandler := manageuser.NewHandler(manageUserService)

	adminRoute.GET("/getUser", manageUserHandler.GetUser)
	adminRoute.PUT("/updateuser", manageUserHandler.UpdateUser)
	adminRoute.DELETE("/deleteuser/:id", manageUserHandler.DeleteUser)
}
