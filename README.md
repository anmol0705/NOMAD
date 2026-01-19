# Nomad: The Polyglot Portable Agent

**Nomad** is a zero-footprint, self-configuring AI development environment designed for high-performance coding on the move. By decoupling the inference engine and models from the host operating system, Nomad transforms any portable storage device into a high-tier intelligence hub for C, C++, Python, and Java.

Unlike traditional AI wrappers, Nomad is built for the "Isolated Developer"—those working in restricted environments, offline, or across multiple workstations who require a consistent, expert-level pair programmer without system-wide installations.

---

## Technical Architecture

The Nomad ecosystem operates through a dual-binary system to ensure maximum portability and minimal RAM overhead.

### 1. The Setup Utility (setup.exe)

The "Builder" responsible for environment auditing. It detects the host OS, maps the drive-relative paths, creates the directory skeleton, and fetches the necessary inference binaries from official sources. It handles the "dirty work" of extraction and configuration.

### 2. The Core Agent (agent.exe)

The "Brain." A Go-powered CLI that manages the lifecycle of the Ollama server in stealth mode. It features:

- **Stateful Memory**: Context-aware conversation tracking.
- **Self-Healing**: Automatic verification and pulling of missing models.
- **Stream Processing**: Real-time token delivery with integrated logic-thinking blocks.

---

## Directory Structure

Nomad maintains a strict hierarchy to ensure zero-leakage to the host machine:

```
nomad/
├── setup.exe           # Initial environment builder
├── agent.exe           # Main interaction interface
├── models/             # Encapsulated LLM storage (GGUF/Blobs)
├── tools/              # Localized inference engines (Ollama)
└── workspace/          # Default directory for generated source code
```

---

## Polyglot Intelligence Profiles

Nomad adapts its internal system prompts based on the selected language stack to enforce industry-standard best practices:

| Language | Specialization | Coding Standards |
|----------|---|---|
| **C++** | Competitive Programming | C++20/23, STL Optimization, Zero-overhead abstraction |
| **C** | Systems Architecture | C11/C17, Pointer safety, Manual memory management |
| **Python** | Idiomatic Scripting | PEP 8, Type hints, Functional paradigms |
| **Java** | Enterprise Logic | SOLID Principles, Design Patterns, Google Style Guide |

---

## Deployment & Execution

### Phase I: Environment Provisioning

Run `setup.exe` on the portable drive. The utility will identify the OS (Windows/Linux/Darwin) and provision the `tools/` directory.

### Phase II: Agent Initialization

Run `agent.exe`. On the first execution, the agent will:

1. Initialize the Ollama background service via syscall.
2. Verify the presence of the `qwen2.5-coder:7b` model.
3. Automatically pull the model into the `/models` directory if missing.

### Phase III: Operation

Select the target language from the boot menu. Nomad will inject a specialized "Expert Personality" into the inference stream. All responses include a `<thinking>` block where the AI plans the architectural approach before emitting code.

---

## Operational Guardrails

- **Protocol Conciseness**: No conversational filler. Nomad focuses on raw logic and implementation.
- **Context Isolation**: Every session is tracked via integer slices (Context tokens) to maintain long-form memory without bloating RAM.
- **Stealth Execution**: The inference engine runs with `HideWindow: true` to prevent terminal clutter on the host machine.

---

## Technical Stack

- **Language**: Go (Golang)
- **Inference Engine**: Ollama (Localized)
- **Primary Model**: Qwen 2.5 Coder 7B (Optimized for 4-bit quantization)
- **IPC**: JSON over Localhost HTTP

---

## License

MIT License. Built for developers who believe that intelligence should be as portable as their code.

---

**Start coding anywhere. Never compromise on intelligence.**
