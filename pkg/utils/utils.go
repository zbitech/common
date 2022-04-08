package utils

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

func ReadByteConfig(in io.Reader, fileType string, cfg interface{}) error {
	v := viper.New()
	v.SetConfigType(fileType)
	if err := v.ReadConfig(in); err != nil {
		return err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func ReadConfig(path string, envKeys []string, cfg interface{}) error {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if envKeys != nil && len(envKeys) > 0 {
		for _, envKey := range envKeys {
			value := GetEnv(envKey, "")
			if value == "" {
				return fmt.Errorf("%s value not avaialble", envKey)
			}

			v.Set(envKey, value)
		}
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Base64EncodedString(value string) string {
	return string([]byte(b64.StdEncoding.EncodeToString([]byte(value))))
}

func Base64DecodeString(value string) string {
	b, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		return ""
	}

	return string(b)
}

func MarshalObject(obj interface{}) string {
	if obj != nil {
		c, err := json.Marshal(obj)
		if err != nil {
			return err.Error()
		}

		return string(c)
	}
	return ""
}

func MarshalIndentObject(obj interface{}) string {
	if obj != nil {
		c, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err.Error()
		}

		return string(c)
	}
	return ""
}
