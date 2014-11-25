package heroku

import "os"

func GetEnv(key string, defVal string) (env string) {
	env = os.Getenv(key)

	if env == "" {
		env = defVal
	}
	return
}
