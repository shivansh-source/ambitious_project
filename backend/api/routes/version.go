package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shivansh-source/ambitious_project/backend/controllers"
	"github.com/shivansh-source/ambitious_project/backend/middleware"
)

// SetupVersionRoutes configures all version control related routes
func SetupVersionRoutes(router *gin.RouterGroup) {
	// Apply authentication middleware
	auth := router.Group("")
	auth.Use(middleware.AuthRequired())
	{
		// Get commit history
		auth.GET("/projects/:projectId/commits", controllers.GetCommits)

		// Create new commit
		auth.POST("/projects/:projectId/commits", controllers.CreateCommit)

		// Get specific commit
		auth.GET("/projects/:projectId/commits/:commitId", controllers.GetCommit)

		// List branches
		auth.GET("/projects/:projectId/branches", controllers.GetBranches)

		// Create branch
		auth.POST("/projects/:projectId/branches", controllers.CreateBranch)

		// Switch branch (checkout)
		auth.POST("/projects/:projectId/checkout", controllers.Checkout)

		// Merge branches
		auth.POST("/projects/:projectId/merge", controllers.MergeBranches)

		// Get diff between commits
		auth.GET("/projects/:projectId/diff", controllers.GetDiff)

		// Rollback to specific commit
		auth.POST("/projects/:projectId/rollback", controllers.Rollback)
	}
}
