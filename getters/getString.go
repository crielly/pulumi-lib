package getters

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// GetStringWithDefault is used to get a value if present, but return a default if not
func GetStringWithDefault(cfg config.Config, key string, defValue string) (value pulumi.String) {

	cfgValue := cfg.Get(key)

	if cfgValue != "" {
		value = pulumi.String(cfgValue)
	} else {
		value = pulumi.String(defValue)
	}

	return
}
