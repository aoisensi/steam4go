package steam4go

import (
    "io/ioutil"
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
    "strconv"
)

const (
    BaseUrl = "http://api.steampowered.com"
    ApiVersion = "v0002"
)

type SteamId uint64
type GameId uint32

type SteamApi struct {
    key string
}

func NewSteamApi(key string) *SteamApi {
    return &SteamApi{key: key}
}

func (p *SteamApi) genUrl(ifname, mtname string, params url.Values) string {
    url := strings.Join([]string{BaseUrl, ifname, mtname, ApiVersion}, "/")
    params.Add("key", p.key)
    return url + "?" + params.Encode()
}

func (p SteamId) String() string {
    return strconv.FormatUint(uint64(p), 10)
}

func (p GameId) String() string {
    return strconv.FormatUint(uint64(p), 10)
}

func getJsonFromUrl(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
