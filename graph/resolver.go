package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// schema.resolvers.go:redeclared
// type queryResolver struct {
// 	*Resolver
// }
//
// type mutationResolver struct {
// 	*Resolver
// }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
