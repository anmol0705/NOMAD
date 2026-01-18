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
		"3. Keep variable names human-like and descriptive yet short (e.g., 'currentMax' instead of 'm' or 'llm_generated_variable_name'). " +
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

// --- Helper Functions ---

func askAI(prompt string, context []int) []int {
	req := OllamaRequest{
		Model:   ModelName,
		Prompt:  prompt,
		System:  SystemPrompt,
		Stream:  true, // 1. Stream enable kiya
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

	// 2. JSON Decoder initialize kiya stream read karne ke liye
	decoder := json.NewDecoder(resp.Body)

	fmt.Print("\nAI: ") // Response header

	for {
		var chunk OllamaResponse
		// 3. Ek-ek chunk (token) decode karo
		if err := decoder.Decode(&chunk); err == io.EOF {
			break // Stream khatam ho gayi
		} else if err != nil {
			fmt.Println("\nError decoding stream:", err)
			return nil
		}

		// 4. Token-by-token print (no newline)
		fmt.Print(chunk.Response)

		// 5. Jab final chunk aaye, context update karo
		if len(chunk.Context) > 0 {
			context = chunk.Context
		}
	}
	fmt.Println() // Response ke baad ek naya line
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
	// 1. Setup Paths
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error finding executable path:", err)
		return
	}

	fileDir := filepath.Dir(execPath)
	ollamaPath := filepath.Join(fileDir, "tools", "ollama.exe")
	modelsPath := filepath.Join(fileDir, "models")

	// 2. Ensure models folder exists
	_ = os.MkdirAll(modelsPath, 0755)

	var cmd *exec.Cmd

	// 3. Launch or Skip Ollama
	if checkOllama() {
		fmt.Println("Ollama is already running. Skipping launch...")
	} else {
		fmt.Println("Ollama not found. Starting it now...")

		cmd = exec.Command(ollamaPath, "serve")

		// Set environment and hide window on Windows
		env := os.Environ()
		env = append(env, "OLLAMA_MODELS="+modelsPath)
		cmd.Env = env
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting Ollama:", err)
			return
		}

		// Health Check Loop
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

	// 4. The Agent Loop
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

	// 5. Cleanup
	if cmd != nil && cmd.Process != nil {
		fmt.Println("\nStopping Ollama service...")
		_ = cmd.Process.Kill()
	}

	fmt.Println("Goodbye!")
}
