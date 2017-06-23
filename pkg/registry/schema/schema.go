package schema

import (
	"bytes"
	"path/filepath"

	"github.com/liszt-code/liszt/pkg/assets"
)

const assetSchemaPath = "assets/schema/"

// Build returns the graphql schema
func Build() (schema string, err error) {
	files, err := assets.AssetDir(assetSchemaPath)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	for _, fileName := range files {
		var content []byte
		content, err = assets.Asset(filepath.Join(assetSchemaPath, fileName))
		if err != nil {
			return
		}
		_, err = buf.Write(content)
		if err != nil {
			return
		}
	}
	schema = buf.String()
	return
}
