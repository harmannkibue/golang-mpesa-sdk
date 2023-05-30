package daraja

import (
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httpCaching"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
	"time"
)

const (
	SANDBOX    = 1
	PRODUCTION = 2
)

var (
	BackOffStrategy = []time.Duration{
		1 * time.Second,
		3 * time.Second,
		4 * time.Second,
	}
)

// DarajaService is  Mpesa Service -.
type DarajaService struct {
	Cache       httpCaching.ICache
	HttpRequest httprequest.IhttpRequest
	ApiKey      string
	ApiSecret   string
	Env         int
}

// New return a new Mpesa DarajaService -.
func New(appKey, appSecret string, env int) (*DarajaService, error) {

	if SANDBOX != env && env != PRODUCTION {
		return nil, fmt.Errorf("invalid env tag.Pass SANDBOX or PRODUCTION")
	}
	cache := httpCaching.New("darajaToken", 60*time.Minute)
	httpReq := httprequest.New()

	return &DarajaService{cache, httpReq, appKey, appSecret, env}, nil
}

// base url determined by the environment tags passed -.
func (s DarajaService) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://api.safaricom.co.ke/"
	}
	return "https://sandbox.safaricom.co.ke/"
}
