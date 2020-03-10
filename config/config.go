package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

//LoadFromFile : load parameter from .env
func LoadFromFile() error {
	return godotenv.Load(".env")

}

// GetVal : function to read an environment or return a default value
func GetVal(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}

// GetValAsInt : function to read an environment variable into integer or return a default value
func GetValAsInt(name string) int {
	valueStr := GetVal(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return 0
}

// GetValAsBool : function to read an environment variable into a bool or return default value
func GetValAsBool(name string) bool {
	valStr := GetVal(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return false
}

// GetValAsSlice : function to read an environment variable into a string slice or return default value
// Print out each role
//for _, role := range conf.UserRoles {
//	fmt.Println(role)
//    }
func GetValAsSlice(name string, sep string) []string {
	valStr := GetVal(name)
	if valStr == "" {
		return []string{}
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
