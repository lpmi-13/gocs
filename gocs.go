package gocs

import (
  "bytes"
  "context"
  "encoding/json"
  "fmt"
  "io"
  "io/util"
  "net/http"
  "net/url"
  "reflect"
  "strconv"
  "strings"
  "sync"
  "time"

  "github.com/google/go-querystring/query"
)

// we're hardcoding the region (dub) for now, but we need to have a way to eventually set this dynamically
const (
  libraryVersion = "0.1.0"
  defaultBaseURL = "https://dub.cloudsigma.com/api/2.0/"
  userAgent = "gocs/" + libraryVersion
  mediaType = "application/json"

  // we probably won't end up using these, not sure if the API returns this info
  headerReateLimit = "RateLimit-Limit"
  headerRateRemaining = "RateLimit-Remaining"
  headerRateReset = "RateLimit-Reset"
)

// this is the main client that manages communication with the Cloud Sigma V2.0 API
type Client struct {
  client *http.Client

  BaseURL *url.URL

  UserAgent string

  // individual services we're using to group types of API communication
  Balance  BalanceService
  Server  ServerService
  Region  RegionService

  onRequestCompleted RequestCompletionCallback

  // optional headers to send along
  headers map[string]string
}
