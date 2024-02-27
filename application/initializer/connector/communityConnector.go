// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package connector

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/repository"
	"clove-api/domain/community/presentation"
	"clove-api/domain/community/service"
)

// The PostContainerConnection function injects the dependence of services related to Post objects.
func PostContainerConnection(db *ent.Client) *presentation.PostController {
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := presentation.NewPostController(postService)
	return postController
}

// The CommentContainerConnection function injects the dependence of services related to Comment objects.
func CommentContainerConnection(db *ent.Client) *presentation.CommentController {
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := presentation.NewCommentController(commentService)
	return commentController
}

// The ProfileContainerConnection function injects the dependence of services related to Profile objects.
func ProfileContainerConnection(db *ent.Client) *presentation.ProfileController {
	profileRepository := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepository)
	profileController := presentation.NewProfileController(profileService)
	return profileController
}
