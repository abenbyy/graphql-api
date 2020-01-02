package query

import (
	res "github.com/abenbyy/graphql-api/graphql/query/resolver"
	typ "github.com/abenbyy/graphql-api/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootQuery",
		Fields: graphql.Fields{
			"users":{
				Type:	graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
				Description: "Get All Users",
			},
			"userbyemail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByEmail,
				Description: "Get User Based on Email",
			},
			"userbyphone":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByPhone,
				Description: "Get User Based on Phone",
			},
			"userbyphoneoremail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args:graphql.FieldConfigArgument{
					"arg": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByPhoneOrEmail,
				Description: "Find User to Validate Login",
			},
		},
	})
}