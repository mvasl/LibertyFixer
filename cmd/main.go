package main

import (
	"fmt"
	"os"
	"os/exec"

	w32 "golang.org/x/sys/windows"

	"LibertyFixer/gpu"
	"LibertyFixer/shader"
)

func main() {
	gpuType, err := gpu.DetectGPUType()
	switch {
	case err != nil || gpuType == gpu.TypeUnknown:
		w32.MessageBox(0, w32Str(err.Error()), w32Str("GPU Detection"), w32.MB_OK)
		os.Exit(1)
	case gpuType == gpu.TypeIntel:
		w32.MessageBox(0, w32Str("Your GPU is not supported yet"), w32Str(""), w32.MB_OK)
		os.Exit(1)
	}
	if err = shader.ForceShaderRegen(gpuType); err != nil {
		w32.MessageBox(0, w32Str(err.Error()), w32Str("Shader handler"), w32.MB_OK)
		os.Exit(1)
	}

	launchGame()
}

func launchGame() {
	proc := exec.Command(".\\Cyberpunk2077.exe")
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	if err := proc.Run(); err != nil {
		w32.MessageBox(0, w32Str(fmt.Sprintf("Could not launch game: %v", err)), w32Str("Launcher"), w32.MB_OK)
	}
}

func w32Str(s string) *uint16 {
	return w32.StringToUTF16Ptr(s)
}
