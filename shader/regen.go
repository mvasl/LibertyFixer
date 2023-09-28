package shader

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"LibertyFixer/gpu"
)

func ForceShaderRegen(gpuType gpu.Type) error {
	switch gpuType {
	case gpu.TypeAMD:
		if err := forceRegenAMDDx12Shaders(); err != nil {
			return fmt.Errorf("force AMD shader recompilation: %w", err)
		}
		return nil
	}

	return errors.New("unsupported GPU")
}

func forceRegenAMDDx12Shaders() error {
	appdata := os.Getenv("APPDATA")
	lastSlash := strings.LastIndex(appdata, "\\")
	appdata = appdata[:lastSlash]
	cachePath := filepath.Join(appdata, "Local", "AMD")

	if err := filepath.Walk(filepath.Join(cachePath, "DxcCache"), rmIgnoreErrors); err != nil {
		return err
	}
	if err := filepath.Walk(filepath.Join(cachePath, "DxCache"), rmIgnoreErrors); err != nil {
		return err
	}
	return nil
}

func rmIgnoreErrors(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if filepath.Ext(path) == ".parc" {
		_ = os.Remove(path)
	}

	return nil
}
