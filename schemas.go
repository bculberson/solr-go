package solr

import (
	"fmt"
	"net/http"
	"strings"
)

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

type HTTPer interface {
	Do(*http.Request) (*http.Response, error)
}

type CompositeKey struct {
	ShardKey string
	DocID    string
}

type HashRange struct {
	Low  int32
	High int32
}

func NewCompositeKey(id string) (CompositeKey, error) {
	keys := strings.Split(id, "!")

	if len(keys) == 1 {
		if strings.Index(id, "!") < 0 {
			return CompositeKey{DocID: keys[0]}, nil
		} else {
			return CompositeKey{ShardKey: id}, nil
		}
	}

	if len(keys) == 2 {
		return CompositeKey{ShardKey: keys[0], DocID: keys[1]}, nil
	}

	if len(keys) > 2 {
		return CompositeKey{}, fmt.Errorf("Cant deal with composite keys %s", id)
	}

	panic("failed all cases")
}

type ClusterProps struct {
	UrlScheme string `json:"urlScheme"`
}
