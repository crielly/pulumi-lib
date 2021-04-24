package getters

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// GetIntWithDefault is used to get a value if present, but return a default if not
func GetIntWithDefault(cfg config.Config, key string, defValue int) (value pulumi.Int) {

	cfgValue := cfg.GetInt(key)

	if cfgValue != 0 {
		value = pulumi.Int(cfgValue)
	} else {
		value = pulumi.Int(defValue)
	}

	return
}
