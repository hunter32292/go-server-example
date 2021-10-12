package options

import (
	"strings"

	"github.com/spf13/pflag"
)

// ServerFlags - the flags used to setup a simple server
type ServerFlags struct {
	Name     string
	Host     string
	Port     int
	CertPath string
	KeyPath  string
}

// DefaultFlags - Setup Default flags for init of server
// For TLS
// certPath  = "./cert.pem"
// keyPath   = "./key.pem"
func DefaultFlags() *ServerFlags {

	return &ServerFlags{
		Name:     "example-server",
		Host:     "localhost",
		Port:     8080,
		CertPath: "",
		KeyPath:  "",
	}

}

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}
