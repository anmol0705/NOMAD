# Nomad
## Portable C++ Coding Agent

## Overview

The Portable C++ Coding Agent is a zero-dependency, command-line tool designed for offline C++ development. Built with Go and Ollama, it provides an AI-powered coding assistant that runs entirely from an SD card or portable storage device, without requiring system-wide installation or modifying the host machine's configuration.

## Key Features

- **Portable Deployment**: Runs entirely from SD card or portable storage without system-wide dependencies
- **Offline Capability**: Complete offline operation with no internet connectivity required after model download
- **Zero System Footprint**: No installation to system drives (e.g., C:\ drive); all configuration remains local to the deployment directory
- **Intelligent Code Generation**: Optimized for C++ with support for modern standards (C++20/23)
- **Context-Aware Assistance**: Maintains conversation context for continuous development support
- **Clean Code Output**: Produces idiomatic C++ code with meaningful variable names and proper namespace usage

## Architecture

The Portable C++ Coding Agent comprises the following components:

```
/MyCodingAgent
├── agent.exe              # Main executable (Go-based CLI)
├── main.go                # Source code
├── README.md              # Documentation
├── models/                # AI model storage (initially empty)
└── tools/
    └── ollama.exe         # Ollama inference engine
```

## Installation and Setup

### Prerequisites

- Windows operating system
- Portable storage device (SD card, USB drive, external SSD)
- Sufficient storage for AI models (approximately 5-10 GB depending on model selection)

### Initial Setup

1. **Deploy the Agent**
   - Copy the entire project directory to your portable storage device

2. **Configure Environment Variables**
   - Open a terminal in the `tools/` directory
   - Set the models directory path:
     ```
     set OLLAMA_MODELS=E:\MyCodingAgent\models
     ```
     (Replace `E:` with your actual drive letter)

3. **Download AI Model**
   - Execute the following command:
     ```
     ollama pull qwen2.5-coder:7b
     ```
   - This downloads the AI model to your local models directory

### Running the Agent

- **GUI Method**: Double-click `agent.exe`
- **CLI Method**: Run `agent.exe` from command line or terminal

## Design Philosophy

This tool prioritizes portability and independence. All configuration, models, and dependencies are contained within the deployment directory, enabling seamless operation across different machines without system modifications.

### Code Style Standards

The generated code follows competitive programming practices with emphasis on:
- Modern C++ (C++20/23) features
- Idiomatic namespace usage
- Meaningful variable naming conventions
- Optimized algorithmic patterns

## Scope and Limitations

This agent is specialized for C++ development tasks. Queries outside of C++ programming will be politely declined with a reminder of the tool's specific purpose.

## License

MIT License

## Support and Contributions

For issues, feature requests, or contributions, please refer to the project repository documentation.
