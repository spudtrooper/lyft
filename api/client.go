package api

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/bluele/gcache"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/io"
)

var (
	token       = flags.String("lyft_token", "auth token")
	userCreds   = flag.String("lyft_user_creds", ".user_creds.json", "file with user credentials")
	noUserCreds = flag.Bool("lyft_no_user_creds", false, "Don't use user creds event if it exists")
)

// TODO: need to make genopts regex more flexible, need to have "--function Base" exactly here, which isn't great
//go:generate genopts --function Base token:string

// Client is a client for lyft.com
type Client struct {
	token    string
	urlCache gcache.Cache
}

// NewClientFromFlags creates a new client from command line flags
func NewClientFromFlags() (*Client, error) {
	if *token != "" {
		client := NewClient(*token)
		return client, nil
	}
	if *noUserCreds {
		client := NewClient("")
		return client, nil
	}
	if *userCreds != "" {
		client, err := NewClientFromFile(*userCreds)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return nil, errors.Errorf("Must set --user & --token or --creds_file")
}

// NewClient creates a new client directly from the API Key
func NewClient(token string) *Client {
	return &Client{
		token:    token,
		urlCache: gcache.New(20).LRU().Build(),
	}
}

// NewClientFromFile creates a new client from a JSON file like `user_creds-example.json`
func NewClientFromFile(credsFile string) (*Client, error) {
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &Client{
		token:    creds.LyftAccessToken,
		urlCache: gcache.New(20).LRU().Build(),
	}, nil
}

type Creds struct {
	LyftAccessToken string `json:"lyft_access_token"`
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	if !io.FileExists(credsFile) {
		return
	}
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
}

var backoff = []time.Duration{1 * time.Second, 2 * time.Second, 5 * time.Second}

func canRetry(err error) bool {
	if strings.Contains(err.Error(), "status code: 403") {
		return true
	}
	return false
}
