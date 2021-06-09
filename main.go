package main

import (
    "github.com/joho/godotenv"
    "github.com/turnage/graw"
    "github.com/turnage/graw/reddit"
    "log"
    "os"
    "wallstreet-bots/pkg"
)


func boot() WSBType.WSBot {
    log.SetPrefix("wsb:")

    log.Print("Loading current configuration")
    err := godotenv.Load()
    log.Print("Configuration loaded")
    if err != nil {
        log.Fatal("Unable to load environment configuration", err)
    }
    if bot, err := reddit.NewBot(agentFromEnv()); err != nil {
        log.Fatal("Unable to create bot instance", err)
    } else {
        _ = graw.Config{Subreddits: []string{"wallstreetbets"}}
        log.Print("learned something")
        return WSBType.WSBot{WsBot: bot}
    }
    return WSBType.WSBot{}
}

func main() {
    _ = boot()
    log.Print("Bot startup successful, let's get stupid")
}

func agentFromEnv() reddit.BotConfig {
    log.Print("Building config from reddit client package")
    return reddit.BotConfig{
        Agent: os.Getenv(""),
        App: reddit.App{
            Username: os.Getenv(""),
            Password: os.Getenv(""),
            ID:       os.Getenv("REDDIT_CLIENT_ID"),
            Secret:   os.Getenv("REDDIT_CLIENT_SECRET"),
        }}
}
