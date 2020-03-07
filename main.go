package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

func main() {
	fmt.Println("start golang....")
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			db, err := gorm.Open("root:ronglong01@(127.0.0.1)/crawler?charset=utf8&parseTime=True&loc=Local")
			if err != nil {
				log.Fatal(err)
			}
			db.AutoMigrate(&models.Url{})
			return nil
		},
		Commands: []*cli.Command{
			{
			  Name:    "complete",
			  Aliases: []string{"c"},
			  Usage:   "complete a task on the list",
			  Action:  func(c *cli.Context) error {
				return nil
			  },
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
