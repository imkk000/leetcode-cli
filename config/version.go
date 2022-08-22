package config

var (
	Version = "1.0.5"
	Debug   = "true"
)

func IsDebugMode() bool {
	return Debug == "true"
}

func GetMode() string {
	if IsDebugMode() {
		return "debug"
	}
	return "release"
}
