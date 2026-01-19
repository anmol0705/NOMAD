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
	"syscall" // For hiding the background window on Windows
	"time"
)

// --- Global Constants & Configuration ---

const (
	OllamaURL    = "http://localhost:11434/api/generate"
	ModelName    = "qwen2.5-coder:7b"
	SystemPrompt = "You are a Legendary Competitive Programmer and C++ Grandmaster. " +
		"Your mission is to provide ultra-optimized, high-performance C++ code that is both elegant and efficient. " +
		"CONSTRAINTS: " +
		"1. Always include 'using namespace std;' at the top to keep the workspace clean and concise. " +
		"2. Write code like a professional CP veteran: use efficient algorithms (O(log n), O(n), etc.), " +
		"proper use of STL containers, and clean logic. " +
		"3. Keep variable names human-like and descriptive yet short (e.g., 'currentMax' instead of 'm'). " +
		"4. NO UNNECESSARY COMMENTS. The code should be so clean it explains itself. " +
		"5. STRICT GUARDRAIL: If the user asks about anything other than C++ or logic, politely tell them 'My brain is only compiled for C++' and redirect them. " +
		"6. Prioritize Modern C++ (C++20/23) features for performance."
)

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

// Struct to check existing models
type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

// --- Helper Functions ---

// ensureModelExists checks if the model is available locally; if not, it pulls it.
func ensureModelExists(ollamaPath, modelsPath string) {
	// 1. Check existing models via API
	resp, err := http.Get("http://localhost:11434/api/tags")
	if err == nil {
		defer resp.Body.Close()
		var tags TagsResponse
		json.NewDecoder(resp.Body).Decode(&tags)

		for _, m := range tags.Models {
			if m.Name == ModelName || m.Name == ModelName+":latest" {
				return // Model found, exit function
			}
		}
	}

	// 2. Pull model if not found
	fmt.Printf("\nModel '%s' not found. Downloading (this may take time)...\n", ModelName)

	pullCmd := exec.Command(ollamaPath, "pull", ModelName)

	// Ensure the pull command uses our custom models folder
	env := os.Environ()
	env = append(env, "OLLAMA_MODELS="+modelsPath)
	pullCmd.Env = env

	// Redirect output to console so user sees progress
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr

	err = pullCmd.Run()
	if err != nil {
		fmt.Println("Error downloading model:", err)
		return
	}
	fmt.Println("✅ Model download complete!")
}

func askAI(prompt string, context []int) []int {
	req := OllamaRequest{
		Model:   ModelName,
		Prompt:  prompt,
		System:  SystemPrompt,
		Stream:  true,
		Context: context,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return nil
	}

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
			fmt.Println("\nError decoding stream:", err)
			return nil
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
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error finding executable path:", err)
		return
	}

	fileDir := filepath.Dir(execPath)
	ollamaPath := filepath.Join(fileDir, "tools", "ollama.exe")
	modelsPath := filepath.Join(fileDir, "models")

	_ = os.MkdirAll(modelsPath, 0755)

	var cmd *exec.Cmd

	if checkOllama() {
		fmt.Println("Ollama is already running. Skipping launch...")
	} else {
		fmt.Println("Ollama not found. Starting it now...")

		cmd = exec.Command(ollamaPath, "serve")
		env := os.Environ()
		env = append(env, "OLLAMA_MODELS="+modelsPath)
		cmd.Env = env
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting Ollama:", err)
			return
		}

		ready := false
		for i := 1; i <= 15; i++ {
			fmt.Printf("Waiting for Ollama to wake up... (Attempt %d)\n", i)
			if checkOllama() {
				ready = true
				break
			}
			time.Sleep(1 * time.Second)
		}

		if !ready {
			fmt.Println("❌ Ollama failed to start in time.")
			return
		}
		fmt.Println("✅ Ollama is ready!")
	}

	// CHECK & DOWNLOAD MODEL BEFORE STARTING CHAT
	ensureModelExists(ollamaPath, modelsPath)

	fmt.Println("\n--- C++ AI Assistant is Ready! ---")
	fmt.Println("Type your question or 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	var chatContext []int

	for {
		fmt.Print("\nC++ Assistant > ")
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()

		if userInput == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		if userInput == "" {
			continue
		}

		fmt.Println("Thinking...")
		chatContext = askAI(userInput, chatContext)
	}

	if cmd != nil && cmd.Process != nil {
		fmt.Println("\nStopping Ollama service...")
		_ = cmd.Process.Kill()
	}

	fmt.Println("Goodbye!")
}
