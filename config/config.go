package config

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv : function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	defaultVal = "Default value kosong"
	if defaultVal == "" {
		defaultVal = "Default value kosong"
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetEnvAsInt : function to read an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// GetEnvAsBool : function to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsSlice : function to read an environment variable into a string slice or return default value
// Print out each role
//for _, role := range conf.UserRoles {
//	fmt.Println(role)
//    }
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	return val
}

/*
GITHUB_USERNAME=craicoverflow
GITHUB_API_KEY=TCtQrZizM1xeo1v92lsVfLOHDsF7TfT5lMvwSno
MAX_USERS=10
USER_ROLES_SLICE=admin,super_admin,guest
DEBUG_MODE=false
*/
