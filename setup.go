package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// URL Registry for Ollama Releases
var downloadMap = map[string]string{
	"windows": "https://github.com/ollama/ollama/releases/latest/download/ollama-windows-amd64.zip",
	"darwin":  "https://github.com/ollama/ollama/releases/latest/download/Ollama-darwin.zip",
	"linux":   "https://ollama.com/install.sh",
}

func main() {
	fmt.Println("==========================================")
	fmt.Println("🌍 NOMAD ARCHITECT: UNIVERSAL SETUP")
	fmt.Println("==========================================")

	// 1. Path & OS Detection
	baseDir, _ := os.Getwd()
	detectedOS := runtime.GOOS
	fmt.Printf("📍 Location: %s\n", baseDir)
	fmt.Printf("🔍 Detected OS: %s\n", strings.ToUpper(detectedOS))

	// 2. Interactive Menu
	fmt.Println("\nSelect Target OS for this SD Card:")
	fmt.Printf("1. %s (Auto-detected)\n", strings.ToUpper(detectedOS))
	fmt.Println("2. Windows (AMD64)")
	fmt.Println("3. macOS (Apple Silicon/Intel)")
	fmt.Println("4. Linux (Install Script)")
	fmt.Print("\nEnter choice (1-4) [Default 1]: ")

	var choice string
	fmt.Scanln(&choice)

	targetOS := detectedOS
	switch choice {
	case "2":
		targetOS = "windows"
	case "3":
		targetOS = "darwin"
	case "4":
		targetOS = "linux"
	}

	downloadURL, exists := downloadMap[targetOS]
	if !exists {
		fmt.Println("❌ Error: Selected OS is not supported yet.")
		return
	}

	// 3. Create Nomad Directory Structure
	fmt.Println("\n📁 Building Nomad workspace...")
	folders := []string{"tools", "models", "workspace"}
	for _, f := range folders {
		path := filepath.Join(baseDir, f)
		_ = os.MkdirAll(path, 0755)
		fmt.Printf("   ✅ Ready: /%s\n", f)
	}

	// 4. Handle Platform Specific Installation
	if targetOS == "linux" {
		handleLinuxSetup(baseDir, downloadURL)
	} else {
		handleZipSetup(baseDir, downloadURL, targetOS)
	}

	fmt.Println("\n==========================================")
	fmt.Printf("🎉 NOMAD IS READY FOR %s!\n", strings.ToUpper(targetOS))
	fmt.Println("   - Engine: /tools")
	fmt.Println("   - Models: /models")
	fmt.Println("   - Build:  go build -o agent.exe main.go")
	fmt.Println("==========================================")
	fmt.Println("Press Enter to finish...")
	fmt.Scanln()
}

// --- Specialized Setup Handlers ---

func handleZipSetup(baseDir, url, osName string) {
	zipPath := filepath.Join(baseDir, "tools", "ollama_engine.zip")
	toolsDir := filepath.Join(baseDir, "tools")

	fmt.Printf("\n📡 Downloading Ollama for %s...\n", strings.ToUpper(osName))
	err := downloadFile(zipPath, url)
	if err != nil {
		fmt.Printf("❌ Download failed: %v\n", err)
		return
	}

	fmt.Println("\n📦 Extracting engine binaries...")
	err = unzip(zipPath, toolsDir)
	if err != nil {
		fmt.Printf("❌ Extraction failed: %v\n", err)
		return
	}

	fmt.Println("🧹 Removing temporary archive...")
	os.Remove(zipPath)
}

func handleLinuxSetup(baseDir, url string) {
	fmt.Println("\n🐧 Linux detected: Downloading install.sh...")
	scriptPath := filepath.Join(baseDir, "tools", "install.sh")
	err := downloadFile(scriptPath, url)
	if err != nil {
		fmt.Printf("❌ Failed to get script: %v\n", err)
		return
	}
	fmt.Println("💡 Tip: On Linux, run 'chmod +x tools/install.sh' to initialize.")
}

// --- Core Plumbing (Download & Unzip) ---

type WriteCounter struct {
	Total      uint64
	Downloaded uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Downloaded += uint64(n)
	percentage := float64(wc.Downloaded) / float64(wc.Total) * 100
	fmt.Printf("\r   📥 Downloading: %.2f%% ", percentage)
	return n, nil
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error: %s", resp.Status)
	}

	counter := &WriteCounter{Total: uint64(resp.ContentLength)}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	return err
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
