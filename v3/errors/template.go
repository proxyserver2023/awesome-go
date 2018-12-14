package errors

/* import (
	"fmt"
	"io/ioutil"
	"strings"
	"gopkg.in/yaml.v2"
) */

type (
	Params        map[string]interface{}
	errorTemplate struct {
		Message string `yaml:"`
	}
)
