// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
