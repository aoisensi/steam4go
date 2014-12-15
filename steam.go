package steam4go

import (
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"
)

const (
    BaseUrl = "http://api.steampowered.com"
    ApiVersion = "v0002"
)

type SteamApi struct {
    key string
}

type Player struct {
    //public data
    SteamID string
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
    GameId int
    GameServerIp string
    GameExtraInfo string
    //CityId int
    LocCountryCode string
    LocStateCode string
    LocCityId int
}

func NewSteamApi(key string) *SteamApi {
    return &SteamApi{key: key}
}

func (p *SteamApi) genUrl(ifname, mtname string, params url.Values) string {
    url := strings.Join([]string{BaseUrl, ifname, mtname, ApiVersion}, "/")
    params.Add("key", p.key)
    return url + "?" + params.Encode()
}

func (p *SteamApi) GetPlayerSummary(steamid string) (player Player, err error) {
    players, err := p.GetPlayerSummaries([]string{steamid})
    if err != nil {
        return Player{}, err
    }
    return players[0], nil
}

func (p *SteamApi) GetPlayerSummaries(steamids []string) ([]Player, error) {
    params := url.Values{}
    params.Add("steamids", strings.Join(steamids, ","))
    url := p.genUrl("ISteamUser", "GetPlayerSummaries", params)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    var r struct { Response struct {Players []Player} }
    err = json.Unmarshal(body, &r)
    if err != nil {
        return nil, err
    }
    return r.Response.Players, nil
}
