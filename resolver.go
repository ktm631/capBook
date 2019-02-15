package capBook

import (
	"capBook/database_config"
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateAuthor(ctx context.Context, input NewAuthor) (Author, error) {
	var db = database_config.DbConn()
	var author Author
	author.Name = input.Name
	author.Surname = input.Surname
	db.Create(&author)
	return author, nil
}
func (r *mutationResolver) CreatePublisher(ctx context.Context, input NewPublisher) (Publisher, error) {
	var db = database_config.DbConn()
	var publisher Publisher
	publisher.Name = input.Name
	//c.BindJSON(&publisher)

	db.Create(&publisher)
	return publisher, nil
}
func (r *mutationResolver) CreateRental(ctx context.Context, input NewRental) (Rental, error) {
	var db = database_config.DbConn()
	var rental Rental
	rental.StartDate = input.StartDate
	rental.EndDate = input.EndDate
	rental.ExpectedEndDate = input.ExpectedEndDate
	rental.BookID = input.BookID
	rental.UserID = input.UserID

	db.Create(&rental)
	return rental, nil
}
func (r *mutationResolver) CreateBook(ctx context.Context, input NewBook) (Book, error) {
	var db = database_config.DbConn()
	var book Book
	book.Title = input.Title
	book.IsFree = input.IsFree
	book.Isbn = input.Isbn
	book.Edition = input.Edition
	book.DescriptionURL = input.DescriptionURL
	book.AuthorID = input.AuthorID
	book.LocationID = input.LocationID
	book.OwnerID = input.OwnerID
	book.PublisherID = input.PublisherID

	db.Create(&book)
	return book, nil
}
func (r *mutationResolver) CreateLocation(ctx context.Context, input *NewLocation) (Location, error) {
	var db = database_config.DbConn()
	var location Location
	location.Building = input.Building
	location.Room = input.Room

	db.Create(&location)
	return location, nil
}
func (r *mutationResolver) DeleteUser(ctx context.Context, user_id int) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAuthor(ctx context.Context, author_id int) (int, error) {
	var db = database_config.DbConn()
	var author Author
	db.Where("author_id = ?", author_id).Delete(&author)
	return author_id, nil
}
func (r *mutationResolver) DeletePublisher(ctx context.Context, publisher_id int) (int, error) {
	var db = database_config.DbConn()
	var publisher Publisher
	db.Where("publisher_id = ?", publisher_id).Delete(&publisher)
	return publisher_id, nil
}
func (r *mutationResolver) DeleteRental(ctx context.Context, rental_id int) (int, error) {
	var db = database_config.DbConn()
	var rental Rental
	db.Where("rental_id = ?", rental_id).Delete(&rental)
	return rental_id, nil
}
func (r *mutationResolver) DeleteBook(ctx context.Context, book_id int) (int, error) {
	var db = database_config.DbConn()
	var book Book
	db.Where("rental_id = ?", book_id).Delete(&book)
	return book_id, nil
}
func (r *mutationResolver) DeleteLocation(ctx context.Context, location_id int) (int, error) {
	var db = database_config.DbConn()
	var location Location
	db.Where("rental_id = ?", location_id).Delete(&location)
	return location_id, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input NewUser, user_id int) (User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAuthor(ctx context.Context, input NewAuthor, author_id int) (Author, error) {
	var db = database_config.DbConn()
	var author Author
	db.Where("author_id = ?", author_id).First(&author)
	author.Surname = input.Surname
	author.Name = input.Name
	db.Save(&author)
	return author, nil
}
func (r *mutationResolver) UpdatePublisher(ctx context.Context, input NewPublisher, publisher_id int) (Publisher, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRental(ctx context.Context, input NewRental, rental_id int) (Rental, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateBook(ctx context.Context, input NewBook, book_id int) (Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateLocation(ctx context.Context, input NewLocation, location_id int) (Location, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) Authors(ctx context.Context) ([]Author, error) {
	var db = database_config.DbConn()
	var authors []Author

	db.Find(&authors)
	return authors, nil
}
func (r *queryResolver) Publishers(ctx context.Context) ([]Publisher, error) {
	var db = database_config.DbConn()
	var publishers []Publisher

	db.Find(&publishers)
	return publishers, nil
	//panic("not implemented")
}
func (r *queryResolver) Rentals(ctx context.Context) ([]Rental, error) {
	var db = database_config.DbConn()
	var rentals []Rental

	db.Find(&rentals)
	return rentals, nil
}
func (r *queryResolver) Books(ctx context.Context) ([]Book, error) {
	var db = database_config.DbConn()
	var books []Book

	db.Find(&books)
	return books, nil
}
func (r *queryResolver) Locations(ctx context.Context) ([]Location, error) {
	var db = database_config.DbConn()
	var locations []Location

	db.Find(&locations)
	return locations, nil
}
func (r *queryResolver) User(ctx context.Context, user_id int) (*User, error) {
	panic("not implemented")
}
func (r *queryResolver) Author(ctx context.Context, author_id int) (*Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Publisher(ctx context.Context, publisher_id int) (*Publisher, error) {
	var db = database_config.DbConn()
	var publisher Publisher
	db.Where("publisher_id = ?", publisher_id).First(&publisher)
	return &publisher, nil
}
func (r *queryResolver) Rental(ctx context.Context, rental_id int) (*Rental, error) {
	var db = database_config.DbConn()
	var rental Rental
	db.Where("publisher_id = ?", rental_id).First(&rental)
	return &rental, nil
}
func (r *queryResolver) Book(ctx context.Context, book_id int) (*Book, error) {
	var db = database_config.DbConn()
	var book Book
	db.Where("publisher_id = ?", book_id).First(&book)
	return &book, nil
}
func (r *queryResolver) Location(ctx context.Context, location_id int) (*Location, error) {
	var db = database_config.DbConn()
	var location Location
	db.Where("publisher_id = ?", location_id).First(&location)
	return &location, nil
}
