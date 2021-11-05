package main

import (
	"github.com/booscaaa/zenvia-go-sdk/zenvia"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get configuration variavables
	authToken := viper.GetString("zenvia.auth_token")

	// Configure Sandbox access into Zenvia api

	zenviaAccess := zenvia.ZenviaConfig().
		ApiToken(authToken).
		Sandbox()

	zenviaSdk := zenvia.Instance(zenviaAccess)

	// From and To with this format: PREFIX DDD NUMBER WITHOUT MASK
	zenviaSdk.SendSingleSMS("5554999999999", "5554999999999", "Message")
}
