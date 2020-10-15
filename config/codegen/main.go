// +build codegen

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"text/template"
)

const (
	sharedConfigType = "&SharedConfig{}"
	envConfigType    = "&EnvConfig{}"
)

var implAsserts = map[string][]string{
	"SharedConfigProfileProvider":              {envConfigType, `WithSharedConfigProfile("")`},
	"SharedConfigFilesProvider":                {envConfigType, `WithSharedConfigFiles(nil)`},
	"CustomCABundleProvider":                   {envConfigType, `WithCustomCABundle(nil)`},
	"RegionProvider":                           {envConfigType, sharedConfigType, `WithRegion("")`, `WithEC2IMDSRegion{}`},
	"CredentialsProviderProvider":              {`WithCredentialsProvider(nil)`},
	"DefaultRegionProvider":                    {`WithDefaultRegion("")`},
	"EC2RoleCredentialOptionsProvider":         {`WithEC2RoleCredentialOptions(nil)`},
	"EndpointCredentialOptionsProvider":        {`WithEndpointCredentialOptions(nil)`},
	"EndpointResolverProvider":                 {`WithEndpointResolver(nil)`},
	"APIOptionsProvider":                       {`WithAPIOptions(nil)`},
	"HTTPClientProvider":                       {`WithHTTPClient(nil)`},
	"AssumeRoleCredentialOptionsProvider":      {`WithAssumeRoleCredentialOptions(nil)`},
	"WebIdentityRoleCredentialOptionsProvider": {`WithWebIdentityRoleCredentialOptions(nil)`},
}

var tplProviderTests = template.Must(template.New("tplProviderTests").Funcs(map[string]interface{}{
	"SortKeys": func(m map[string][]string) []string {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		return keys
	},
}).Parse(`
// Code generated by github.com/aws/aws-sdk-go-v2/config. DO NOT EDIT.

package config

{{ $sortedKeys := SortKeys . }}
{{- range $_, $provider := $sortedKeys }}
	{{- $implementors := index $ $provider -}}
	// {{ $provider }} implementor assertions
	var (
		{{- range $_, $impl := $implementors }}
		_ {{ $provider }} = {{ $impl }}
		{{- end }}
	)
{{ end -}}
`))

var config = struct {
	OutputPath string
}{}

func init() {
	flag.StringVar(&config.OutputPath, "output", "", "output file path")
	flag.Parse()
}

func main() {
	if len(config.OutputPath) == 0 {
		panic(fmt.Errorf("output path must be specified"))
	}

	file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %v", err))
	}
	defer file.Close()

	err = tplProviderTests.Execute(file, implAsserts)
	if err != nil {
		panic(fmt.Errorf("failed to execute template: %v", err))
	}
}
