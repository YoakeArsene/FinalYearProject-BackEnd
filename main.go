package main

import (
	db "GDN-delivery-management/db/sql"
	handle "GDN-delivery-management/delivery/http"
	"GDN-delivery-management/repository"
	"GDN-delivery-management/router"
	"database/sql"
	"github.com/labstack/echo/v4"
	//migrate "github.com/rubenv/sql-migrate"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("./app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	psqlInfo := os.Getenv("DBSOURCE")
	driver, err := sql.Open("postgres", psqlInfo)
	//Migrate(driver)
	if err != nil {
		log.Println(err)
		return
	}

	//_, err = driver.Exec(`INSERT INTO roles (role_name, ticker)
	//							VALUES ('System Admin', 'SAD') ON CONFLICT DO NOTHING`)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//_, err = driver.Exec(`INSERT INTO roles (role_name, ticker)
	//							VALUES ('User', 'USR') ON CONFLICT DO NOTHING`)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//_, err = driver.Exec(`INSERT INTO users (id, username, email, password, avatar, role_ticker)
	//							VALUES ('123e4567-e89b-12d3-a456-426614174001', 'John Doe', 'johndoe@gmail.com', '7fe8babbd1346dbbd1861e12d9c70ac42771d039ea257be82f02ad81079bbc60', 'http://localhost:3000/images/miku.jpg', 'USR') ON CONFLICT DO NOTHING`)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	queries := db.New(driver)
	userRepo := repository.NewUserRepo(queries)
	sessionRepo := repository.NewSessionRepo(queries)
	roleRepo := repository.NewRoleRepo(queries)
	gameRepo := repository.NewGameRepo(queries)
	userHandle := handle.UserHandler{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	}
	roleHandler := handle.RoleHandler{
		RoleRepo: roleRepo,
	}
	gameHandler := handle.GameHandler{
		GameRepo: gameRepo,
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	routerSetup := router.Router{
		Echo:        e,
		UserHandler: userHandle,
		RoleHandler: roleHandler,
		GameHandler: gameHandler,
	}
	routerSetup.SetupRouter()
	e.Logger.Fatal(e.Start(":1313"))
}

//func Migrate(db *sql.DB) {
//	migrations := &migrate.FileMigrationSource{
//		Dir: "migrations",
//	}
//	d, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
//	if err != nil {
//		log.Println(err)
//	}
//
//	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Printf("Applied %d & %d migrations!\n", d, n)
//}
