package tomlenv

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	pkgerrors "github.com/pkg/errors"
)

// DecodeEnv is parse multiple files to v.
// common(s) at first, env at last.
// if common not given, default try parsing common.yml.
func DecodeEnv(env string, path string, v interface{}, common ...string) (toml.MetaData, error) {
	if env == "" || path == "" {
		return toml.MetaData{}, errors.New("env and path are both required.")
	}

	if len(common) == 0 {
		common = []string{"common"}
	}

	for _, b := range common {
		if _, err := toml.DecodeFile(fmt.Sprintf("%s/%s.toml", path, b), v); err != nil {
			return toml.MetaData{}, pkgerrors.Wrapf(err, "%s/%s.toml DecodeFile failed", path, b)
		}
	}

	return toml.DecodeFile(fmt.Sprintf("%s/%s.toml", path, env), v)
}
