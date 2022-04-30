package config

import (
	"fmt"
	"log"
	"net/http"
	"winartodev/coba-graphql/delivery"
	"winartodev/coba-graphql/entity"
	"winartodev/coba-graphql/handler"
	"winartodev/coba-graphql/repository"
	"winartodev/coba-graphql/schema"
	"winartodev/coba-graphql/schema/mutation"
	"winartodev/coba-graphql/schema/query"
	"winartodev/coba-graphql/schema/resolver"
	"winartodev/coba-graphql/usecase"

	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

func NewConfig() Config {
	var cfg Config

	gotenv.Load(".env")
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

func Serve() {
	cfg := NewConfig()
	db, err := NewPostgresConfig(&cfg)
	if err != nil {
		panic(err)
	}
	// database migration
	db.AutoMigrate(entity.User{})

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userResolver := resolver.NewUserResolver(userUsecase)
	userQuery := query.NewUserQuery(userResolver)
	userMutation := mutation.NewUserMutation(userResolver)
	userSchema := schema.NewUserSchema()
	userHandler := delivery.NewUserHandler(userUsecase, userQuery, userMutation, userSchema)

	h := handler.NewHandler(&userHandler)

	s := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		Handler: h,
	}

	fmt.Printf("serve on http://%v", s.Addr)
	if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
		log.Fatal(serr)
	}
}
