package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// --- Global Constants & Configuration ---

const (
	OllamaURL = "http://localhost:11434/api/generate"
	ModelName = "qwen2.5-coder:7b"
)

// LanguageConfig holds specific rules for each language
type LanguageConfig struct {
	Name  string
	Rules string
}

var languageRegistry = map[string]LanguageConfig{
	"1": {"C++", "Legendary Competitive Programmer. Use Modern C++20/23, 'using namespace std;', STL containers, and ultra-optimized logic."},
	"2": {"C", "Low-level Systems Architect. Focus on pointer safety, manual memory management, and C11/C17 standards. No globals."},
	"3": {"Python", "Pythonic Architect. Follow PEP 8 strictly. Use type hints, list comprehensions, and write idiomatic, readable code."},
	"4": {"Java", "Enterprise Software Architect. Follow Google Java Style Guide. Prioritize SOLID principles and Design Patterns."},
}

// --- Data Types for JSON Communication ---

type OllamaRequest struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	System  string `json:"system"`
	Stream  bool   `json:"stream"`
	Context []int  `json:"context"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Context  []int  `json:"context"`
}

type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

// --- Helper Functions ---

func getDynamicSystemPrompt(config LanguageConfig) string {
	return fmt.Sprintf(`You are the 'Nomad' Polyglot Expert. Your specialty is %s.
    
    CORE OPERATIONAL RULES:
    - THINKING: Before writing any code, briefly describe your logical approach inside <thinking> tags.
    - STYLE: %s
    - CONCISENESS: No filler phrases like 'Certainly' or 'I can help with that'. Go straight to the solution.
    - VARIABLE NAMES: Use descriptive, human-like names (e.g., 'currentUser' instead of 'u').
    - GUARDRAIL: If the user asks non-coding questions, say 'NOMAD is strictly a coding environment.'`,
		config.Name, config.Rules)
}

func ensureModelExists(ollamaPath, modelsPath string) {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err == nil {
		defer resp.Body.Close()
		var tags TagsResponse
		json.NewDecoder(resp.Body).Decode(&tags)
		for _, m := range tags.Models {
			if m.Name == ModelName || m.Name == ModelName+":latest" {
				return
			}
		}
	}

	fmt.Printf("\nModel '%s' not found. Downloading (this may take time)...\n", ModelName)
	pullCmd := exec.Command(ollamaPath, "pull", ModelName)
	env := os.Environ()
	env = append(env, "OLLAMA_MODELS="+modelsPath)
	pullCmd.Env = env
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	_ = pullCmd.Run()
}

func askAI(prompt string, context []int, sysPrompt string) []int {
	req := OllamaRequest{
		Model:   ModelName,
		Prompt:  prompt,
		System:  sysPrompt,
		Stream:  true,
		Context: context,
	}

	jsonData, _ := json.Marshal(req)
	resp, err := http.Post(OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error connecting to Ollama:", err)
		return nil
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	fmt.Print("\nAI: ")
	for {
		var chunk OllamaResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil {
			break
		}
		fmt.Print(chunk.Response)
		if len(chunk.Context) > 0 {
			context = chunk.Context
		}
	}
	fmt.Println()
	return context
}

func checkOllama() bool {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err == nil {
		resp.Body.Close()
		return true
	}
	return false
}

// --- Main Logic ---

func main() {
	execPath, _ := os.Executable()
	fileDir := filepath.Dir(execPath)
	ollamaPath := filepath.Join(fileDir, "tools", "ollama.exe")
	modelsPath := filepath.Join(fileDir, "models")
	_ = os.MkdirAll(modelsPath, 0755)

	var cmd *exec.Cmd
	if !checkOllama() {
		fmt.Println("Ollama not found. Starting it now...")
		cmd = exec.Command(ollamaPath, "serve")
		env := os.Environ()
		env = append(env, "OLLAMA_MODELS="+modelsPath)
		cmd.Env = env
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_ = cmd.Start()
		time.Sleep(5 * time.Second) // Initial wait
	}

	ensureModelExists(ollamaPath, modelsPath)

	// --- Language Selection Menu ---
	fmt.Println("\n==========================================")
	fmt.Println("🌍 NOMAD POLYGLOT INTERFACE")
	fmt.Println("==========================================")
	fmt.Println("Select your stack:")
	fmt.Println("1. C++ (Competitive)  2. C (Systems)  3. Python (Scripts)  4. Java (Enterprise)")
	fmt.Print("\nSelection (1-4) [Default 1]: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	if choice == "" {
		choice = "1"
	}

	selectedConfig, exists := languageRegistry[choice]
	if !exists {
		selectedConfig = languageRegistry["1"]
	}

	finalSystemPrompt := getDynamicSystemPrompt(selectedConfig)
	fmt.Printf("\n✅ Nomad is now configured for %s\n", strings.ToUpper(selectedConfig.Name))

	// --- Chat Loop ---
	var chatContext []int
	for {
		fmt.Printf("\nNomad [%s] > ", selectedConfig.Name)
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()

		if userInput == "exit" {
			break
		}
		if userInput == "" {
			continue
		}

		fmt.Println("Thinking...")
		chatContext = askAI(userInput, chatContext, finalSystemPrompt)
	}

	if cmd != nil && cmd.Process != nil {
		_ = cmd.Process.Kill()
	}
	fmt.Println("Goodbye!")
}
