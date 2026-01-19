# Nomad
### The Portable C++ Coding Agent

**Nomad** is a lightweight, self-contained AI-powered coding assistant engineered for C++ developers who value independence and portability. Packed with advanced language models and zero system footprint, it transforms your SD card into a fully-functional offline development environment—no installation, no system modifications, no compromises.

Built with **Go** and **Ollama**, Nomad is the perfect companion for developers working offline, in restricted environments, or simply seeking ultimate portability without sacrificing intelligence.

---

## ✨ Core Features

| Feature | Description |
|---------|-------------|
| **🚀 True Portability** | Execute entirely from SD card or portable storage—no system installation required |
| **🔌 Complete Offline** | Full functionality without internet after initial model setup |
| **🎯 Zero System Impact** | All configuration and models remain isolated—nothing touches your system drives |
| **🧠 Intelligent Code Generation** | AI-powered assistance optimized for modern C++ (C++20/23) standards |
| **💬 Context-Aware** | Maintains conversation memory for seamless, continuous development sessions |
| **✅ Production-Ready Code** | Generates idiomatic C++ with meaningful names and optimized patterns |

## 🏗️ Project Structure

```
nomad/
├── agent.exe              ✨ Main executable (Go-powered)
├── main.go                📝 Source implementation
├── README.md              📖 Documentation
├── models/                🧠 AI model directory (self-managed)
└── tools/
    └── ollama.exe         ⚙️  Inference engine
```

Every component is self-contained and portable—nothing needs to be installed system-wide.

## 🚀 Getting Started

### Requirements

- **OS**: Windows
- **Storage**: Portable device (SD card, USB drive, external SSD)
- **Space**: 5–10 GB for AI models
- **Internet**: Required only for initial model download

### Setup in 3 Steps

#### 1️⃣ Deploy Nomad
```bash
# Copy the entire project to your portable storage device
# That's it! No installation scripts, no package managers.
```

#### 2️⃣ Configure Model Storage
```bash
# Navigate to tools/ directory and run:
set OLLAMA_MODELS=E:\nomad\models
# Replace E: with your actual drive letter
```

#### 3️⃣ Download Your AI Model
```bash
# From tools/ directory:
ollama pull qwen2.5-coder:7b
# The model downloads to your local models/ directory
```

### Running Nomad

```bash
# Option 1: Double-click agent.exe
# Option 2: CLI mode
./agent.exe
```

Done. Your portable C++ assistant is ready.

## 💡 Design Philosophy

**Nomad** embodies three core principles:

### 🎯 **Independence**
No system-wide dependencies. No registry modifications. No hidden files scattered across your machine. Everything you need lives in one portable folder.

### ⚡ **Efficiency**
Competitive programming-inspired code generation. Optimized algorithms. Meaningful variable names. Modern C++ idioms. Fast, intelligent, production-ready output.

### 🌍 **Universality**
Plug your device into any Windows machine and code immediately. Same environment, same models, same results. True development portability.

## 📋 What Nomad Can Do

✅ **C++ Code Generation** – Write, refactor, and optimize C++ with AI assistance  
✅ **Algorithm Design** – Competitive programming-optimized solutions  
✅ **Code Review** – Real-time feedback on style, performance, and correctness  
✅ **Documentation** – Generate clear, professional code comments  
✅ **Problem-Solving** – Contextual assistance for debugging and architecture  

### Scope

Nomad is a specialized C++ development assistant. While it can engage with general programming concepts, it maintains focus on C++-specific tasks and optimizations.

---

## 🛠️ Technical Stack

| Component | Technology |
|-----------|-----------|
| **Runtime** | Go (compiled binary) |
| **AI Engine** | Ollama + Qwen 2.5 Coder 7B |
| **Language** | C++ (target) |
| **Deployment** | Portable (no system dependencies) |

---

## 📄 License

MIT License – Built for developers, by developers.

---

## 🤝 Contributing

Found a bug? Have an idea? Contributions and feedback are welcome. Check the project repository for guidelines.

---

**Start coding anywhere. Never compromise on intelligence.**
