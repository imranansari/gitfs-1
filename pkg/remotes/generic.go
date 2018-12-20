package remotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mjpitz/gitfs/pkg/config"
	"github.com/nytlabs/gojee"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewGenericRemote(config *config.Generic) Remote {
	return &genericRemote{
		config: config,
	}
}

var _ Remote = &genericRemote{}

type genericRemote struct {
	config *config.Generic
}

func parseJQResult(jqResult string) []string {
	lines := strings.Split(jqResult, "\n")

	parsed := make([]string, len(lines))
	for i, line := range lines {
		cleaned := strings.TrimPrefix(line, "\"")
		cleaned = strings.TrimSuffix(cleaned, "\"")
		parsed[i] = cleaned
	}
	return parsed
}

func (r *genericRemote) ListRepositories() ([]string, error) {
	tokens, err := jee.Lexer(r.config.Selector)
	if err != nil {
		return nil, err
	}

	parser, err := jee.Parser(tokens)
	if err != nil {
		return nil, err
	}

	repositories := make([]string, 0)
	for page := 1; true; page++ {
		fullUrl := fmt.Sprintf(
			"%s%s?%s=%d&%s=%d",
			r.config.BaseUrl,
			r.config.Path,
			r.config.PageParameter,
			page,
			r.config.PerPageParameter,
			r.config.PageSize,
		)

		resp, err := http.Get(fullUrl)
		if err != nil {
			return nil, errors.Wrap(err,
				fmt.Sprintf("failed to get url: %s", fullUrl))
		}

		if resp.StatusCode == http.StatusNotFound {
			logrus.Infof("encountered a 404. assuming end of data")
			break
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read body")
		}

		var umsg jee.BMsg
		if err := json.Unmarshal(body, &umsg); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal JSON")
		}

		result, err := jee.Eval(parser, umsg)
		if err != nil {
			return nil, errors.Wrapf(err,
				fmt.Sprintf("failed to extract response from page using selector: %s", r.config.Selector))
		}

		resultArray := result.([]interface{})
		for _, entry := range resultArray {
			entryString := entry.(string)
			repositories = append(repositories, entryString)
		}

		if int32(len(resultArray)) < r.config.PageSize {
			logrus.Infof("encountered an incomplete page. assuming end of data")
			break
		}
	}

	return repositories, nil
}