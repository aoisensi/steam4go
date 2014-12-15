package steam4go

import (
    "net/url"
    "strings"
    "encoding/json"
)

type Player struct {
    //public data
    SteamId SteamId `json:",string"`
    PersonaName string
    ProfileUrl string
    Avater string
    AvaterMedium string
    AvaterFull string
    PersonaState int
    CommunityVisibilityState int
    ProfileState int
    LastLogoff int
    CommentPermission int

    //private data
    RealName string
    PrimaryClanId string
    TimeCreated int
    GameId GameId `json:",string"`
    GameServerIp string
    GameExtraInfo string
    CityId int // DEPRECATED
    LocCountryCode string
    LocStateCode string
    LocCityId int
}

func (p *SteamApi) GetPlayerSummary(steamid SteamId) (*Player, error) {
    players, err := p.GetPlayerSummaries([]SteamId{steamid})
    if err != nil {
        return nil, err
    }
    return &players[0], nil
}

func (p *SteamApi) GetPlayerSummaries(steamids []SteamId) ([]Player, error) {
    params := url.Values{}
    ids := make([]string, len(steamids))
    for i, id := range steamids {
        ids[i] = id.String()
    }
    params.Add("steamids", strings.Join([]string(ids), ","))
    url := p.genUrl("ISteamUser", "GetPlayerSummaries", params)
    body, err := getJsonFromUrl(url)
    if err != nil {
        return nil, err
    }
    var r struct {Response struct {Players []Player}}
    err = json.Unmarshal(body, &r)
    if err != nil {
        return nil, err
    }
    return r.Response.Players, nil
}
