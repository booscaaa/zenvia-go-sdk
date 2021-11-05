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
	// From and To with this format: PREFIX DDD NUMBER WITHOUT MASK
	zenviaAccess := zenvia.ZenviaConfig().
		ApiToken(authToken).
		From("5554999999999").
		To("5554999999999").
		Sandbox()

	zenviaSdk := zenvia.Instance(zenviaAccess)

	zenviaSdk.SendSMS("Message")
}
