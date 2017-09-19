package book_route

import (
	"github.com/gin-gonic/gin"
	"yellowroad_library/services/auth_serv"
	"yellowroad_library/utils/app_error"
	"yellowroad_library/database/entities"
	"yellowroad_library/services/book_serv"
	"yellowroad_library/utils/api_response"
	"strconv"
	"net/http"
	"yellowroad_library/database/repo/book_repo"
)

func FetchSingleBook(bookRepo book_repo.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		book_id, convErr := strconv.Atoi(c.Param("book_id"))
		if (convErr != nil){
			appErr := app_error.New(http.StatusUnprocessableEntity,"","ID must be a valid integer value!")
			c.JSON(api_response.ConvertErrWithCode(appErr))
			return
		}

		book, findErr := bookRepo.FindById(book_id)
		if findErr != nil {
			c.JSON(api_response.ConvertErrWithCode(findErr))
			return
		}

		c.JSON(api_response.SuccessWithCode(book))
		return
	}
}

func CreateBook(authService auth_serv.AuthService, bookService book_serv.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User
		var formData createBookForm

		//Get logged in user
		user, err := authService.GetLoggedInUser(c.Copy());
		if err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		//Get form data to create book with
		if err := c.BindJSON(&formData); err != nil {
			var err app_error.AppError = app_error.Wrap(err)
			c.JSON( api_response.ConvertErrWithCode(err) )
			return
		}

		//Create the book
		book := entities.Book {
			CreatorId: user.ID,
			Title: formData.Title,
			Description: formData.Description,
		}
		if err := bookService.CreateBook(user, &book); err != nil {
			c.JSON(api_response.ConvertErrWithCode(err))
			return
		}

		c.JSON(api_response.SuccessWithCode(
			gin.H{"book": book},
		))
	}
}
type createBookForm struct {
	Title string
	Description string
}