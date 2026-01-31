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
	"regexp"
	"strings"
	"syscall"
	"time"
)

// --- Global Configuration ---
const OllamaURL = "http://localhost:11434/api/generate"

type ModelConfig struct {
	ID          string
	Description string
	Size        string
	Efficiency  string // Added for hardware recommendation
}

// Updated Registry with CPU-First Models
var ModelRegistry = map[string]ModelConfig{
	"1": {"qwen2.5-coder:3b", "CPU King (Balanced Smart/Fast)", "1.9GB", "Best for 4-8GB RAM"},
	"2": {"qwen2.5-coder:1.5b", "Ultra-Fast (Real-time Typing)", "900MB", "Best for <4GB RAM"},
	"3": {"qwen2.5-coder:7b", "High Intelligence (Heavier)", "4.7GB", "Requires 8GB+ RAM"},
	"4": {"phi3:mini", "Logic Specialist (Microsoft)", "2.3GB", "Excellent Reasoning"},
}

type LanguageConfig struct {
	Name string
	Ext  string
	Rules string
}

var languageRegistry = map[string]LanguageConfig{
	"1": {"C++", ".cpp", "Competitive Programming standards. Modern C++20."},
	"2": {"C", ".c", "Systems level. Memory safety and pointers focus."},
	"3": {"Python", ".py", "PEP 8 standards. Idiomatic and clean."},
	"4": {"Java", ".java", "Enterprise standards. SOLID principles."},
}

var ActiveModel string
var ActiveModelChoice string

// --- Data Types ---
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

func extractCode(text string) string {
	re := regexp.MustCompile("(?s)```(?:\\w+)?\n(.*?)\n```")
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func ensureModelExists(ollamaPath, modelsPath, targetModel string) {
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err == nil {
		defer resp.Body.Close()
		var tags TagsResponse
		json.NewDecoder(resp.Body).Decode(&tags)
		for _, m := range tags.Models {
			if m.Name == targetModel || m.Name == targetModel+":latest" {
				return
			}
		}
	}

	fmt.Printf("\n🚀 Initializing CPU-Optimized Brain: %s (%s)\n", targetModel, ModelRegistry[ActiveModelChoice].Size)
	pullCmd := exec.Command(ollamaPath, "pull", targetModel)
	env := os.Environ()
	env = append(env, "OLLAMA_MODELS="+modelsPath)
	pullCmd.Env = env
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	_ = pullCmd.Run()
}

func askAI(prompt string, context []int, sysPrompt string) ([]int, string) {
	req := OllamaRequest{
		Model:   ActiveModel,
		Prompt:  prompt,
		System:  sysPrompt,
		Stream:  true,
		Context: context,
	}

	var fullResponse strings.Builder
	jsonData, _ := json.Marshal(req)
	resp, err := http.Post(OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("\n❌ Engine unreachable. Check tools/ollama.exe")
		return nil, ""
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	fmt.Print("\nAI: ")
	for {
		var chunk OllamaResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil { break }
		fmt.Print(chunk.Response)
		fullResponse.WriteString(chunk.Response)
		if len(chunk.Context) > 0 {
			context = chunk.Context
		}
	}
	fmt.Println()
	return context, fullResponse.String()
}

// --- Main Logic ---

func main() {
	execPath, _ := os.Executable()
	fileDir := filepath.Dir(execPath)
	ollamaPath := filepath.Join(fileDir, "tools", "ollama.exe")
	modelsPath := filepath.Join(fileDir, "models")
	workspacePath := filepath.Join(fileDir, "workspace")
	
	_ = os.MkdirAll(modelsPath, 0755)
	_ = os.MkdirAll(workspacePath, 0755)

	if _, err := http.Get("http://localhost:11434/api/tags"); err != nil {
		fmt.Println("🤖 Starting CPU Inference Engine...")
		cmd := exec.Command(ollamaPath, "serve")
		env := os.Environ()
		env = append(env, "OLLAMA_MODELS="+modelsPath)
		// CPU Optimization: Limit threads to physical cores if needed
		// env = append(env, "OLLAMA_NUM_PARALLEL=1") 
		cmd.Env = env
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_ = cmd.Start()
		time.Sleep(5 * time.Second)
	}

	fmt.Println("\n==========================================")
	fmt.Println("🧠 NOMAD: CPU-OPTIMIZED INTELLIGENCE")
	fmt.Println("==========================================")
	for k, v := range ModelRegistry {
		fmt.Printf("%s. %-20s | %-6s | %s\n", k, v.ID, v.Size, v.Efficiency)
	}
	fmt.Print("\nSelect Brain [Default 1]: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ActiveModelChoice = scanner.Text()
	if ActiveModelChoice == "" { ActiveModelChoice = "1" }
	
	selected, exists := ModelRegistry[ActiveModelChoice]
	if !exists { selected = ModelRegistry["1"] }
	ActiveModel = selected.ID

	ensureModelExists(ollamaPath, modelsPath, ActiveModel)

	fmt.Println("\nSelect Stack: 1.C++ | 2.C | 3.Python | 4.Java")
	fmt.Print("Choice: ")
	scanner.Scan()
	langKey := scanner.Text()
	if langKey == "" { langKey = "1" }
	selectedConfig := languageRegistry[langKey]
	
	sysPrompt := fmt.Sprintf("You are Nomad. Specialty: %s. Rule: %s. Use <thinking> tags.", 
		selectedConfig.Name, selectedConfig.Rules)

	fmt.Printf("\n⚡ Ready! Model: %s | Mode: %s\n", ActiveModel, selectedConfig.Name)

	var chatContext []int
	var lastRes string
	for {
		fmt.Printf("\nNomad [%s] > ", selectedConfig.Name)
		if !scanner.Scan() { break }
		input := scanner.Text()

		if input == "exit" { break }
		if strings.HasPrefix(input, "save ") {
			fName := strings.TrimSpace(strings.TrimPrefix(input, "save "))
			code := extractCode(lastRes)
			if code != "" {
				if !strings.Contains(fName, ".") { fName += selectedConfig.Ext }
				_ = os.WriteFile(filepath.Join(workspacePath, fName), []byte(code), 0644)
				fmt.Printf("💾 Saved to workspace/%s\n", fName)
			}
			continue
		}

		fmt.Println("Processing...")
		chatContext, lastRes = askAI(input, chatContext, sysPrompt)
	}
}