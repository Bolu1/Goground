package main

import(
	"fmt"
	"github.com/Bolu1/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/Bolu1/go-fiber-crm/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRoutes(app *fiber.App){

	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead",  lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBconn, err = gorm.Open("mysql", "guest:1234@/nest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to databse")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){

	app := fiber.New()

	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}