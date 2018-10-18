package repo

import (
	"database/sql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"yellowroad_library/test_utils"
)

func TestGormBookRepository(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given a live SQL connection", t, test_utils.WithRealSqlDBTransaction(func(db *sql.Tx){

		Convey("Insert should work", func () {
			repo := BookRepository{db}
			book, err := repo.Insert(BookInsertParams{
				Title: "New Book",
				Description: "New Book's Description",
			})
			So(err, ShouldBeNil)

			So(book.ID, ShouldNotEqual, 0)
			So(book.Description, ShouldEqual,"New Book's Description")
			So(book.Title, ShouldEqual, "New Book")

			Convey("Delete should work", func () {
				repo := BookRepository{db}

				err := repo.Delete(book.ID)

				So(err, ShouldBeNil)
			})
		})

		Convey("Deleting by an unknown ID should throw an error", func () {
			repo := BookRepository{db}

			err := repo.Delete(-100)

			So(err, ShouldNotBeNil)
			So(err.EndpointMessage(), ShouldEqual, "Book with id of -100 was not found")
		})

	}))
}