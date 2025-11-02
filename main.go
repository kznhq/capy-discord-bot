package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	"time"
	"database/sql"

	"github.com/kznhq/capyDiscordBot/handlers"
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
	"github.com/go-sql-driver/mysql"
)

func initRoleDb() error {
	var err error

	// get a DB handler for the database storing information used for react for role messages/assigninment
	// this connection is used in multiple other files so that's why we're initializing here
	// we also prepare statements here for similar reasons since we only do a few unique queries and just change values
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")
	
	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPasswd,
		Net:    "tcp",
		Addr:   dbAddr,
		DBName: dbName,
	}

	utils.RoleDb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("Error when trying to get handler for role database: %w", err)
	}

	utils.RoleDb.SetConnMaxLifetime(time.Minute) // we want to implement a cache so shouldn't need a db connection lifetime to be open too long
	// our MySQL instance is hosted on Aiven, Aiven docs for MySQL say max connections should be 75 * RAM + 1 and we have 1 GB = 0.93 GiB of RAM in free tier
	utils.RoleDb.SetMaxOpenConns(70)
	utils.RoleDb.SetMaxIdleConns(10) // idle connections are a subset of open ones

	utils.GetRoleStatement, err = utils.RoleDb.Prepare("SELECT roleid FROM roleassigningmessagestable WHERE rolename = ? AND guildid = ?")
	if err != nil {
		return fmt.Errorf("Error when preparing SQL statement for deleting roles: %w", err)
	}

	utils.GetRoleFromMsgStatement, err = utils.RoleDb.Prepare("SELECT roleid FROM roleassigningmessagestable WHERE messageid = ?")
	if err != nil {
		return fmt.Errorf("Error when preparing SQL statement for deleting roles: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("Loading .env")

	// need the token to get the client for the bot
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env: ", err) 
		return
	} else {
		fmt.Println(".env loaded")
	}
	discordToken := os.Getenv("DISCORD_TOKEN")

	// make a client
	client, err := discordgo.New("Bot " + discordToken)

	// seeding for the random numbers used to make a random color for new roles
	rand.Seed(time.Now().UnixNano())

	// add the handlers (newMessage takes care of commands and calls the correct response function)
	client.AddHandler(handlers.NewMessageHandler)
	client.AddHandler(handlers.StrawberryHandler) 
	client.AddHandler(handlers.React4roleReactionAddHandler)
	client.AddHandler(handlers.React4roleReactionRemoveHandler)

	err = initRoleDb()
	if err != nil {
		fmt.Println(err) // the error is already formatted to say which part had the error + the actual error itself
		return
	}

	// open the Discord client
	err = client.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err) 
	} else {
		fmt.Println("Connection opened")
	}
	defer client.Close()

	// cleanup if ctrl c signal is sent.
	// if the program terminates abruptly (not ctrl c), MySQL detects
	// the connection terminated and cleans up those connections by itself
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	utils.RoleDb.Close()
}
