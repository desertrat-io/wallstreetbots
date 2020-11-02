package main

import (
    "github.com/joho/godotenv"
    "github.com/turnage/graw"
    "github.com/turnage/graw/reddit"
    "log"
    "os"
)

type wsBot struct {
    wsBot reddit.Bot
}

func boot() {
    log.SetPrefix("wsb:")
    log.Print("Bot startup, let's get stupid")
}

func main() {
    boot()
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Unable to load environment configuration", err)
    }
    log.Print("Configuration loaded")
    if bot, err := reddit.NewBot(agentFromEnv()); err != nil {
        log.Fatal("Unable to create bot instance", err)
    } else {
        log.Print("Building bot from agent")
        cfg := graw.Config{Subreddits: []string{"wallstreetbets"}}
        handler := &wsBot{wsBot: bot}

    }

}

func agentFromEnv() reddit.BotConfig {
    return reddit.BotConfig{
        Agent: os.Getenv(""),
        App: reddit.App{
            Username: os.Getenv(""),
            Password: os.Getenv(""),
            ID:       os.Getenv(""),
            Secret:   os.Getenv(""),
        }}
}
