package steam4go

import (
    "net/url"
    "encoding/json"
)

type PlayerStats struct {
    SteamId SteamId `json:",string"`
    GameName string
    Stats []struct {
        Name string
        Value int
    }
    Achievements []struct {
        Name string
        Achieved int
    }
}

func (p *SteamApi) GetUserStatsForGame(steamid SteamId, appid GameId) (*PlayerStats, error) {
    params := url.Values{}
    params.Add("steamid", steamid.String())
    params.Add("appid", appid.String())
    url := p.genUrl("ISteamUserStats", "GetUserStatsForGame", params)
    body, err := getJsonFromUrl(url)
    if err != nil {
        return nil, err
    }
    var r struct{PlayerStats PlayerStats}
    err = json.Unmarshal(body, &r)
    if err != nil {
        return nil, err
    }
    return &r.PlayerStats, nil
}
