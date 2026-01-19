# 🧭 Nomad — The Polyglot Portable Agent

Nomad is a zero-footprint, self-configuring AI development environment engineered for developers who don't stay in one place—or one machine.

It converts any portable storage device into a fully isolated, high-performance AI coding companion for C, C++, Python, and Java, without touching the host OS.

**No installs.**
**No permissions drama.**
**No environment drift.**

Just plug in, boot up, and code with a senior-level pair programmer—anywhere.

---

## Why Nomad Exists

Most AI coding tools assume:

- Stable internet
- Admin access
- A single machine
- A bloated runtime footprint

Nomad assumes none of that.

It's built for the **Isolated Developer**:

- Working in restricted or offline environments
- Jumping between lab machines, offices, or systems
- Demanding consistency, determinism, and control

Nomad doesn't wrap intelligence around your system.
It brings the system with it.

---

## Architecture Overview

Nomad uses a dual-binary architecture designed for maximum portability, minimal RAM usage, and zero host pollution.

### 🧱 1. Setup Utility — setup.exe

**The Builder.**

This utility performs a one-time audit of the host environment and prepares Nomad's internal ecosystem.

**Responsibilities:**

- Detect host OS (Windows / Linux / macOS)
- Resolve drive-relative paths (no absolute leakage)
- Create the internal directory skeleton
- Download and extract verified inference binaries from official sources

Think of it as the logistics team. It does the dirty work so the agent never has to.

### 🧠 2. Core Agent — agent.exe

**The Brain.**

A Go-powered CLI that orchestrates the entire AI workflow while remaining invisible to the host system.

**Key capabilities:**

- **Stateful Memory** — Maintains long-form conversational context
- **Self-Healing** — Verifies models and auto-pulls missing dependencies
- **Streamed Inference** — Token-level output with structured reasoning blocks
- **Stealth Execution** — Background services run without UI clutter

This is where intelligence lives—and stays contained.

---

## Directory Layout (Zero-Leakage Guaranteed)

Nomad enforces a strict internal hierarchy. Nothing escapes. Nothing pollutes.

```
nomad/
├── setup.exe           # Environment builder
├── agent.exe           # Primary interaction interface
├── models/             # Encapsulated LLM storage (GGUF / blobs)
├── tools/              # Local inference engines (Ollama runtime)
└── workspace/          # Default output directory for generated code
```

No registry writes.
No system PATH edits.
No surprises.

---

## Polyglot Intelligence Profiles

Nomad dynamically injects language-specific expert personas to enforce best practices—not generic autocomplete.

| Language | Focus Area | Enforced Standards |
|----------|---|---|
| **C++** | Competitive & Systems | C++20/23, STL optimization, zero-overhead abstractions |
| **C** | Low-Level Architecture | C11/C17, pointer discipline, manual memory correctness |
| **Python** | Idiomatic Engineering | PEP 8, type hints, functional patterns |
| **Java** | Enterprise Design | SOLID, design patterns, Google Style Guide |

Each profile optimizes not just syntax—but thinking style.

---

## Build Instructions

```bash
// Build the Setup Utility
go build -ldflags="-s -w" -o setup.exe setup.go

// Build the Core Agent
go build -ldflags="-s -w" -o agent.exe main.go
```

Lean binaries. No debug baggage.

---

## Execution Flow

### Phase I — Environment Provisioning

Run `setup.exe` directly from the portable drive.

Nomad will:

- Identify the host OS
- Provision the `tools/` directory
- Prepare the inference runtime

One-time operation per device.

### Phase II — Agent Initialization

Run `agent.exe`.

On first launch, Nomad will:

- Spawn the Ollama server via syscall (hidden execution)
- Verify availability of `qwen2.5-coder:7b`
- Pull the model automatically if missing

No prompts. No babysitting.

### Phase III — Operation

Select your target language from the boot menu.

Nomad will:

- Inject the corresponding expert system prompt
- Maintain session memory using compact context tokens
- Emit responses with a structured `<thinking>` block before code output

You see the plan. Then the implementation.

---

## Operational Guardrails

Nomad enforces discipline by design:

- **No Filler** — Output is logic-first, not conversational
- **Context Isolation** — Memory is token-tracked, not RAM-bloated
- **Stealth Mode** — Background processes never hijack the host terminal
- **Deterministic Behavior** — Same prompt, same intelligence, anywhere

This is a tool—not a chatbot.

---

## Technology Stack

- **Language**: Go (Golang)
- **Inference Runtime**: Ollama (localized)
- **Primary Model**: Qwen 2.5 Coder 7B (4-bit quantized for performance and portability)
- **IPC**: JSON over localhost HTTP

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
