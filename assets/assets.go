package assets

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Dir returns the directory containing the project's assets in the source checkout.
func Dir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok || !filepath.IsAbs(file) {
		return "", fmt.Errorf("Asset-Verzeichnis nicht ermittelbar: absoluter Quellpfad fehlt")
	}
	return filepath.Dir(file), nil
}
