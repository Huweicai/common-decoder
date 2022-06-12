package decoder

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const defaultConfigFilePath = "./config.yaml"

type StaticConfigDecoder struct {
}

func (d *StaticConfigDecoder) Sniffer(text string) Possibility {
	return NotSure
}

func (d *StaticConfigDecoder) readConfig(path string) ([]*staticItem, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("empty file")
	}

	var m map[string]interface{}

	if err := yaml.Unmarshal(content, &m); err != nil {
		return nil, err
	}

	return parseItem(m, nil), nil
}

type staticItem struct {
	Tags []string
	Val  string
}

func (s staticItem) match(target string) bool {
	if s.Val == target {
		return true
	}

	regexMatch, _ := regexp.MatchString(s.Val, target)
	if regexMatch {
		return true
	}

	ip := net.ParseIP(target)
	if ip != nil {
		_, ipNet, err := net.ParseCIDR(s.Val)
		if err == nil {
			if ipNet.Contains(ip) {
				return true
			}
		}
	}

	return false
}

func parseItem(m interface{}, tags []string) (ret []*staticItem) {
	if m == nil {
		return nil
	}

	switch t := m.(type) {
	case map[string]interface{}:
		for key, val := range t {
			ret = append(ret, parseItem(val, append(tags, key))...)
		}

	case string:
		return []*staticItem{
			{
				Tags: tags,
				Val:  t,
			},
		}

	case []interface{}:
		for _, val := range t {
			ret = append(ret, &staticItem{
				Tags: tags,
				Val:  fmt.Sprint(val),
			})
		}
	}
	return ret
}

func (d *StaticConfigDecoder) Decode(text string) (result interface{}, ok bool) {
	items, err := d.readConfig(defaultConfigFilePath)
	if err != nil {
		return nil, false
	}

	var matched []*staticItem
	for _, item := range items {
		if item.match(text) {
			matched = append(matched, item)
		}
	}

	var ret []string
	for _, item := range matched {
		ret = append(ret, strings.Join(item.Tags, "."))
	}

	return &DecodeResult{
		Possibility: MustBe,
		Result:      strings.Join(ret, " | "),
	}, len(ret) != 0
}

func (d *StaticConfigDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
