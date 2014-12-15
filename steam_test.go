package steam4go

import (
    "os"
    "bufio"
    "testing"
    "github.com/k0kubun/pp"
)

const (

    UserId = 76561198049739081
    AppId = 440
    ApiKeyPath = "apikey.txt"
)

func TestGetPlayerSummaries(t *testing.T) {
    api := NewSteamApi(LoadApiKey())
    player, err := api.GetPlayerSummary(UserId)
    if err != nil {
        t.Error(err)
        return
    }
    pp.Println(player)
}

func TestGetUserStatsForGame(t *testing.T) {
    api := NewSteamApi(LoadApiKey())
    state, err := api.GetUserStatsForGame(UserId, AppId)
    if err != nil {
        t.Error(err)
        return
    }
    pp.Println(state)
}

func LoadApiKey() string {
    file, err := os.Open(ApiKeyPath)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    line, _, err := reader.ReadLine()
    if err != nil {
        panic(err)
    }
    return string(line)
}
