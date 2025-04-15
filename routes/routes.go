package routes

import (
	"auction-system/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// User routes
	user := router.Group("/users")
	{
		user.POST("/", controllers.CreateUser)
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	// Item routes
	item := router.Group("/items")
	{
		item.POST("/", controllers.CreateItem)
		item.GET("/", controllers.GetItems)
		item.GET("/:id", controllers.GetItemByID)
		item.PUT("/:id", controllers.UpdateItem)
		item.DELETE("/:id", controllers.DeleteItem)
	}

	// Auction routes
	auction := router.Group("/auctions")
	{
		auction.POST("/", controllers.CreateAuction)
		auction.GET("/", controllers.GetAuctions)
		auction.GET("/:id", controllers.GetAuctionByID)
		auction.PUT("/:id", controllers.UpdateAuction)
		auction.DELETE("/:id", controllers.DeleteAuction)
	}

	// Bid routes
	bid := router.Group("/bids")
	{
		bid.POST("/", controllers.PlaceBid)
		bid.GET("/", controllers.GetBids)
		bid.GET("/auction/:auction_id", controllers.GetBidsForAuction)
	}

	// Wallet routes
	wallet := router.Group("/wallets")
	{
		wallet.POST("/", controllers.CreateWallet)
		wallet.GET("/", controllers.GetWallets)
		wallet.GET("/:id", controllers.GetWalletByID)
		wallet.PUT("/:id", controllers.UpdateWallet)
	}

	// Transaction routes
	transaction := router.Group("/transactions")
	{
		transaction.POST("/", controllers.CreateTransaction)
		transaction.GET("/", controllers.GetTransactions)
		transaction.GET("/:id", controllers.GetTransactionByID)
	}

	// Complaint routes
	complaint := router.Group("/complaints")
	{
		complaint.POST("/", controllers.CreateComplaint)
		complaint.GET("/", controllers.GetComplaints)
		complaint.PUT("/:id/status", controllers.UpdateComplaintStatus)
	}

	// Notification routes
	notification := router.Group("/notifications")
	{
		notification.POST("/", controllers.CreateNotification)
		notification.GET("/", controllers.GetNotifications)
		notification.GET("/:user_id", controllers.GetNotificationsForUser)
	}
}

// package routes

// import (
//     "github.com/gin-gonic/gin"
//     "auction-system/controllers"
// )

// func RegisterRoutes(router *gin.Engine) {
//     user := router.Group("/users")
//     {
//         user.POST("/", controllers.CreateUser)
//         user.GET("/", controllers.GetUsers)
//     }
// }
