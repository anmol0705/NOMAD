# 🧭 Nomad — The Polyglot Portable Agent

Nomad is a zero-footprint, self-configuring AI development environment engineered for developers who don't stay in one place—or one machine.

It converts any portable storage device into a fully isolated, high-performance AI coding companion for C, C++, Python, and Java, without touching the host OS.

**No installs.**
**No permissions drama.**
**No environment drift.**

Just plug in, boot up, and code with a senior-level pair programmer—anywhere.

---

## Why Nomad Exists

Most AI coding tools assume a stable internet, admin access, and a bloated runtime footprint. Nomad assumes none of that. It is built for the **Isolated Developer** working in restricted or offline environments who demands consistency, determinism, and control. Nomad doesn't wrap intelligence around your system—it brings the system with it.

---

## Architecture Overview

Nomad uses a dual-binary architecture designed for maximum portability, minimal RAM usage, and zero host pollution.

### 🧱 1. Setup Utility — setup.exe

**The Builder.**

This utility performs a one-time audit of the host environment and prepares Nomad's internal ecosystem.

**Responsibilities:** Detects host OS (Windows/Linux/macOS), resolves drive-relative paths, creates the internal directory skeleton, and downloads/unzips verified inference binaries.

### 🧠 2. Core Agent — agent.exe

**The Brain.** A Go-powered CLI that orchestrates the AI workflow while remaining invisible to the host system.

**Intelligence Manager:** Features a dynamic Model Selector that prioritizes CPU-optimized Small Language Models (SLMs) for lag-free performance on standard hardware.

**Stateful Memory:** Maintains conversational context via session-based tokens.

**File Generation:** Includes a specialized save command to extract pure code and write it to the workspace.

---

## Directory Layout (Zero-Leakage Guaranteed)

Nomad enforces a strict internal hierarchy. Nothing escapes. Nothing pollutes.

```
nomad/
├── setup.exe           # Environment builder & engine provisioner
├── agent.exe           # Primary interaction & intelligence interface
├── models/             # Encapsulated LLM storage (GGUF / blobs)
├── tools/              # Local inference engines (Ollama runtime)
└── workspace/          # Default output directory for generated code
```

No registry writes. No system PATH edits. No machine-specific artifacts.

---

## Intelligence Registry (CPU-Optimized)

In 2026, the focus is on efficiency. Nomad prioritizes "Small Language Models" (SLMs) that provide high intelligence with minimal lag on standard CPUs.

| Choice | Model ID | Size | Target Hardware |
|--------|----------|------|----------|
| 1 | qwen2.5-coder:3b | 1.9 GB | Recommended (Balanced 4-8GB RAM) |
| 2 | qwen2.5-coder:1.5b | 900 MB | Ultra-Fast (<4GB RAM / Legacy CPUs) |
| 3 | qwen2.5-coder:7b | 4.7 GB | High Logic (8GB+ RAM / Modern CPUs) |
| 4 | phi3:mini | 2.3 GB | Microsoft Logic (Reasoning focused) |

---

## Build Instructions

```bash
// Compile the Provisioner
go build -ldflags="-s -w" -o setup.exe setup.go

// Compile the Agent
go build -ldflags="-s -w" -o agent.exe main.go
```

Note: Uses -s -w flags to strip debug symbols for the smallest possible portable binary size.

---

## Execution Flow

### Phase I — Environment Provisioning

Run setup.exe from the portable drive. It identifies the host OS, provisions the tools/ directory, and downloads the inference runtime.

### Phase II — Agent Initialization

Run agent.exe. Select your "Brain" (Model) and "Stack" (Language). If the model is missing, Nomad pulls it automatically to the /models folder using the localized engine.

### Phase III — Development Loop

Interact with the agent. When you are satisfied with a solution, use the save command:

```
Nomad [C++] > save my_algorithm
```

The agent extracts the pure code block and saves it as workspace/my_algorithm.cpp.

---

## Technology Stack

**Language:** Go (Golang)

**Inference Runtime:** Ollama (Localized)

**Primary Architecture:** Qwen 2.5 Coder (optimized for CPU/RAM constraints)

**Quantization:** 4-bit (K-Quants) for maximum logic-to-size ratio

---

## License

MIT License.

Free as in freedom.
Portable as your code.

---

## Final Word

Nomad isn't trying to be flashy.
It's trying to be reliable in places where flash fails.

If you believe intelligence should move with you—not tie you down—
you already understand why Nomad exists.

**Plug in. Boot up. Stay sharp.**
