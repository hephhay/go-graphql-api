package resolver

import (
	"github.com/hephhay/go-graphql/internal/graphql/generated"
	"github.com/hephhay/go-graphql/internal/service"
)

type Resolver struct {
	userService service.IUserService
}

func NewResolver(userService service.IUserService) *Resolver {
	return &Resolver{userService: userService}
}

// ========== Query ==========
type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// ========== Mutation ==========
type mutationResolver struct {
	*Resolver
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
