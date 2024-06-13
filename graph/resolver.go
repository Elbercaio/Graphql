package graph

import "github.com/Elbercaio/gqlgen-todos/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserDb *database.User
}
