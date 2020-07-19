package options

import (
	"log"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

type group struct {
	key  string
	name string
}

var registry = map[group]interface{}{}

// Init reads all configuration flags
func Init(keys ...string) {

	valid := []string{}
	for g := range registry {
		valid = append(valid, g.key)
	}
	for _, key := range keys {
		if !has(valid, key) {
			log.Panicf("unknown options key: %s", key)
		}
	}
	if !has(keys, "logging") {
		keys = insert(keys, "logging")
	}
	parser := flags.NewNamedParser("app", flags.Default|flags.IgnoreUnknown)
	parser.NamespaceDelimiter = "-"
	for k, v := range registry {
		if has(keys, k.key) {
			var (
				group *flags.Group
				err   error
			)
			if group, err = parser.AddGroup(k.key, k.name, v); err != nil {
				log.Panicf("Initialization error, failed to add group %s: %v", k.key, err)
			}
			group.Namespace = k.key
		}
	}
	if _, err := parser.Parse(); err != nil {
		if !strings.HasPrefix(err.Error(), "Usage:") {
			logrus.Errorf("error parsing options: %v", err)
		}
	} else {
		configLogging()
	}
}

func insert(a []string, s string) []string {
	a = append(a, "")
	copy(a[1:], a[0:])
	a[0] = s
	return a
}

func has(a []string, s string) bool {
	for _, i := range a {
		if i == s {
			return true
		}
	}
	return false
}
