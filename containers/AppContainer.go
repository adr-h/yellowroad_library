package containers

import (
	"fmt"

	"yellowroad_library/config"
	db "yellowroad_library/database"
	"yellowroad_library/database/repo/user_repo"
	"yellowroad_library/http/middleware/auth_middleware"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/services/token_serv"
	"yellowroad_library/database/repo/user_repo/gorm_user_repo"

	"github.com/jinzhu/gorm"
	"yellowroad_library/database/repo/book_repo"
	"yellowroad_library/database/repo/book_repo/gorm_book_repo"
	"yellowroad_library/services/story_serv"
	"yellowroad_library/database/repo/chapter_repo"
	"yellowroad_library/database/repo/chapter_repo/gorm_chapter_repo"
	"yellowroad_library/database/repo/chapterpath_repo"
	"yellowroad_library/database/repo/chapterpath_repo/gorm_chapterpath_repo"
	"yellowroad_library/database/repo/uow"
	"yellowroad_library/services/auth_serv/user_registration_serv"
	"yellowroad_library/services/book_serv/book_create"
	"yellowroad_library/services/book_serv/book_delete"
)

type AppContainer struct {
	dbConn        *gorm.DB
	tokenService  *token_serv.TokenService
	authService   *auth_serv.AuthService
	storyService  *story_serv.StoryService
	configuration config.Configuration
}
//ensure interface implementation
var _ Container = AppContainer{}

func NewAppContainer(config config.Configuration) AppContainer {
	return AppContainer{
		configuration: config,
	}
}

/***********************************************************************************************/
/***********************************************************************************************/
//Non-interface methods

func (ac AppContainer) GetDbConn() *gorm.DB {
	var dbSettings = ac.configuration.Database

	var dbType = dbSettings.Driver
	var connectionString = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbSettings.Host,
		dbSettings.Username,
		dbSettings.Database,
		dbSettings.SSLMode,
		dbSettings.Password,
	)

	if ac.dbConn == nil {
		ac.dbConn = db.Conn(dbType, connectionString)
	}

	return ac.dbConn
}

/***********************************************************************************************/
/***********************************************************************************************/
//Configuration

func (ac AppContainer) GetConfiguration() config.Configuration {
	return ac.configuration
}

/***********************************************************************************************/
/***********************************************************************************************/
//Services

func (ac AppContainer) GetAuthService() auth_serv.AuthService {
	if ac.authService == nil {
		var AuthService auth_serv.AuthService = auth_serv.Default(ac.GetUserRepository(), ac.GetTokenService())
		ac.authService = &AuthService
	}

	return *ac.authService
}
func (ac AppContainer) UserRegistrationService(work *uow.UnitOfWork, autocommit bool) user_registration_serv.UserRegistrationService {
	var workVal uow.UnitOfWork
	if (work == nil){
		workVal = ac.UnitOfWork()
	}else {
		workVal = *work
	}

	userRegistrationService := user_registration_serv.Default(workVal,autocommit)
	return userRegistrationService
}

func (ac AppContainer) GetTokenService() token_serv.TokenService {
	if ac.tokenService == nil {
		var tokenService token_serv.TokenService = token_serv.Default()
		ac.tokenService = &tokenService
	}

	return *ac.tokenService
}

func (ac AppContainer) BookCreateService(work uow.UnitOfWork, autocommit bool) book_create.BookCreateService {
	var workVal uow.UnitOfWork
	if (work == nil){
		workVal = ac.UnitOfWork()
	}else {
		workVal = work
	}

	service := book_create.Default(workVal,autocommit)
	return service
}

func (ac AppContainer) BookDeleteService(work uow.UnitOfWork, autocommit bool) book_delete.BookDeleteService {
	var workVal uow.UnitOfWork
	if (work == nil){
		workVal = ac.UnitOfWork()
	}else {
		workVal = work
	}

	service := book_delete.Default(workVal,autocommit)
	return service
}

func (ac AppContainer) GetStoryService() story_serv.StoryService {
	if ac.storyService == nil {
		var storyService story_serv.StoryService = story_serv.Default(ac.GetChapterRepository())
		ac.storyService = &storyService
	}

	return *ac.storyService
}

/***********************************************************************************************/
/***********************************************************************************************/
//Repositories

func (ac AppContainer) UnitOfWork() uow.UnitOfWork {
	return uow.NewAppUnitOfWork(ac.GetDbConn())
}

func (ac AppContainer) GetUserRepository() user_repo.UserRepository {
	return gorm_user_repo.New(ac.GetDbConn())
}

func (ac AppContainer) GetBookRepository() book_repo.BookRepository {
	return gorm_book_repo.New(ac.GetDbConn())
}

func (ac AppContainer) GetChapterRepository() chapter_repo.ChapterRepository {
	return gorm_chapter_repo.New(ac.GetDbConn())
}

func (ac AppContainer) GetChapterPathRepository() chapterpath_repo.ChapterPathRepository {
	return gorm_chapterpath_repo.New(ac.GetDbConn())
}



/***********************************************************************************************/
/***********************************************************************************************/
//Middleware

func (ac AppContainer) GetAuthMiddleware() auth_middleware.AuthMiddleware {
	return auth_middleware.New(ac.GetTokenService())
}

/***********************************************************************************************/
/***********************************************************************************************/
