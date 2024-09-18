package main

import (
    "log"
    "path/filepath"

    "github.com/zelenin/go-tdlib/client"
)

func main() {
    // client authorizer
    authorizer := client.ClientAuthorizer()
    go client.CliInteractor(authorizer)

    // or bot authorizer
    // botToken := "000000000:gsVCGG5YbikxYHC7bP5vRvmBqJ7Xz6vG6td"
    // authorizer := client.BotAuthorizer(botToken)

    const (
        apiId   = 00000
        apiHash = "8pu9yg32qkuukj83ozaqo5zzjwhkxhnk"
    )

    authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
        UseTestDc:              false,
        DatabaseDirectory:      filepath.Join(".tdlib", "database"),
        FilesDirectory:         filepath.Join(".tdlib", "files"),
        UseFileDatabase:        true,
        UseChatInfoDatabase:    true,
        UseMessageDatabase:     true,
        UseSecretChats:         false,
        ApiId:                  apiId,
        ApiHash:                apiHash,
        SystemLanguageCode:     "en",
        DeviceModel:            "Server",
        SystemVersion:          "1.0.0",
        ApplicationVersion:     "1.0.0",
//        EnableStorageOptimizer: true,
//        IgnoreFileNames:        false,
    }

    _, err := client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
	NewVerbosityLevel: 1,
    })
    if err != nil {
	log.Fatalf("SetLogVerbosityLevel error: %s", err)
    }
    
    tdlibClient, err := client.NewClient(authorizer)
    if err != nil {
        log.Fatalf("NewClient error: %s", err)
    }

    optionValue, err := client.GetOption(&client.GetOptionRequest{
        Name: "version",
    })
    if err != nil {
        log.Fatalf("GetOption error: %s", err)
    }

    log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

    me, err := tdlibClient.GetMe()
    if err != nil {
        log.Fatalf("GetMe error: %s", err)
    }

    log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Usernames)
}