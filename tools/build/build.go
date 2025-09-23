package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type buildTarget struct {
	GOOS       string
	GOARCH     string
	OutputDir  string
	OutputName string
}

// Путь к точке входа приложения
const sourcePath = "./cmd/app/main.go"

func main() {
	// Конфигурация сборки
	targets := []buildTarget{
		{
			GOOS:       "windows",
			GOARCH:     "amd64",
			OutputDir:  "./build",
			OutputName: "ftp_client_windows.exe",
		},
		{
			GOOS:       "linux",
			GOARCH:     "amd64",
			OutputDir:  "./build",
			OutputName: "ftp_client_linux",
		},
		{
			GOOS:       "darwin",
			GOARCH:     "amd64",
			OutputDir:  "./build",
			OutputName: "ftp_client_macos",
		},
	}

	log.Println("Starting full build process...")

	for _, target := range targets {
		log.Printf("Building for %s/%s...", target.GOOS, target.GOARCH)

		// Создаем директорию для билда
		if err := os.MkdirAll(target.OutputDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory %s: %v", target.OutputDir, err)
		}

		outputPath := filepath.Join(target.OutputDir, target.OutputName)
		cmd := exec.Command("go", "build", "-o", outputPath, sourcePath)

		// Установка переменных окружения для кросс-компиляции
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("GOOS=%s", target.GOOS),
			fmt.Sprintf("GOARCH=%s", target.GOARCH),
		)

		// Выполнение команды сборки
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("ERROR building for %s/%s.", target.GOOS, target.GOARCH)
			log.Fatalf("Command failed with error: %v\nOutput:\n%s", err, string(output))
		}

		log.Printf("Successfully built: %s", outputPath)
	}

	log.Println("All builds completed successfully!")
}
