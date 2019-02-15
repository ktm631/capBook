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
	panic("not implemented")
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
	panic("not implemented")
}
func (r *mutationResolver) CreateBook(ctx context.Context, input NewBook) (Book, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLocation(ctx context.Context, input *NewLocation) (Location, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, user_id int) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAuthor(ctx context.Context, author_id int) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePublisher(ctx context.Context, publisher_id int) (int, error) {
	var db = database_config.DbConn()
	var publisher Publisher
	db.Where("publisher_id = ?", publisher_id).Delete(&publisher)
	return publisher_id, nil
}
func (r *mutationResolver) DeleteRental(ctx context.Context, rental_id int) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteBook(ctx context.Context, book_id int) (int, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteLocation(ctx context.Context, location_id int) (int, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) Authors(ctx context.Context) ([]Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Publishers(ctx context.Context) ([]Publisher, error) {
	var db = database_config.DbConn()
	var publishers []Publisher

	db.Find(&publishers)
	return publishers, nil
	//panic("not implemented")
}
func (r *queryResolver) Rentals(ctx context.Context) ([]Rental, error) {
	panic("not implemented")
}
func (r *queryResolver) Books(ctx context.Context) ([]Book, error) {
	panic("not implemented")
}
func (r *queryResolver) Locations(ctx context.Context) ([]Location, error) {
	panic("not implemented")
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
	panic("not implemented")
}
func (r *queryResolver) Book(ctx context.Context, book_id int) (*Book, error) {
	panic("not implemented")
}
func (r *queryResolver) Location(ctx context.Context, location_id int) (*Location, error) {
	panic("not implemented")
}
