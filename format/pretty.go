package format

import (
	"encoding/json"
	"net/url"
)

func Prettify(s string) (string, error) {
	u, err := url.Parse(s)

	if err != nil {
		return "", err
	}

	params, err := url.ParseQuery(u.RawQuery)

	pretty := map[string]interface{}{
		"scheme":   u.Scheme,
		"host":     u.Host,
		"path":     u.Path,
		"params":   params,
		"fragment": u.Fragment,
	}

	output, err := json.MarshalIndent(pretty, "", "    ")

	return string(output), err
}
