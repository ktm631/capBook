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

func (r *mutationResolver) CreatePublisher(c context.Context, input NewPublisher) (Publisher, error) {
	var db = database_config.DbConn()
	var publisher Publisher
	publisher.Name = input.Name
	//c.BindJSON(&publisher)

	db.Create(&publisher)
	return publisher, nil
	//c.JSON(200, publisher)
	//panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, user_id string) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) Publishers(ctx context.Context) ([]Publisher, error) {
	panic("not implemented")
}
