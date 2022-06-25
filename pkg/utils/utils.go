package utils

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
	"strconv"
	"strings"

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

func GetIntEnv(key string, fallback int) int {
	val := GetEnv(key, string(fallback))
	y, e := strconv.Atoi(val)
	if e != nil {
		return fallback
	}
	return y
}

func Base64EncodeString(value string) string {
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

func GetResourceField(obj *unstructured.Unstructured, path string) interface{} {
	var content = obj.UnstructuredContent()
	parts := strings.Split(path, ".")
	for index, part := range parts {
		if index == len(parts)-1 {
			return content[part]
		} else {
			entry := content[part]
			if entry == nil {
				return nil
			}
			content = entry.(map[string]interface{})
		}
	}
	return nil
}

func ReadResourceField(obj *unstructured.Unstructured, path string, data interface{}) error {
	value := GetResourceField(obj, path)
	if value != nil {

		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return json.Unmarshal(valueBytes, &data)
	}
	return nil
}

func SetResourceField(obj *unstructured.Unstructured, path string, value interface{}) error {
	var content = obj.UnstructuredContent()
	parts := strings.Split(path, ".")
	for index, part := range parts {
		if index == len(parts)-1 {
			content[part] = value
		} else {
			entry := content[part]
			if entry == nil {
				entry = make(map[string]interface{})
				content[part] = entry
			}
			content = entry.(map[string]interface{})
		}
	}
	return nil
}

func AddResourceField(obj *unstructured.Unstructured, path string, value interface{}) error {
	var content = obj.UnstructuredContent()
	parts := strings.Split(path, ".")
	for index, part := range parts {
		if index == len(parts)-1 {
			var array []interface{}
			entry := content[part]
			if entry != nil {
				array = entry.([]interface{})
			}
			array = append(array, value)
			content[part] = array
		} else {
			entry := content[part]
			if entry == nil {
				entry = make(map[string]interface{})
				content[part] = entry
			}
			content = entry.(map[string]interface{})
		}
	}
	return nil
}

func RemoveResourceField(obj *unstructured.Unstructured, path string) error {
	var content = obj.UnstructuredContent()
	parts := strings.Split(path, ".")
	for index, part := range parts {
		if index == len(parts)-1 {
			if content[part] != nil {
				delete(content, part)
			}
		} else {
			entry := content[part]
			if entry == nil {
				entry = make(map[string]interface{})
				content[part] = entry
			}
			content = entry.(map[string]interface{})
		}
	}
	return nil
}
