# 🌐 LocalAI: Your Open Source AI Gateway

<p align="center">
  <img width="320" src="./core/http/static/logo.png" alt="LocalAI Logo">
</p>

<p align="center">
  <em>Decentralized. Versatile. Developer-friendly.</em><br>
  A community-driven, fully local alternative to OpenAI APIs—bringing powerful multimodal AI to your fingertips.
</p>

---

## 🚀 Overview

**LocalAI** is a powerful drop-in replacement for OpenAI-compatible APIs, enabling developers to run large language models (LLMs), generate images, process audio, and build multimodal applications—all on consumer-grade hardware, without cloud dependency or GPU requirements. Developed and maintained by [Ettore Di Giacinto](https://github.com/mudler), LocalAI empowers developers with an extensible, modular architecture and a rapidly growing ecosystem.

* 🧩 API-compatible with OpenAI, ElevenLabs, Anthropic and more
* 🧠 Supports multiple LLMs including llama.cpp, GPT-J, Mixtral, Mistral, Phi-2, etc.
* 🖼️ Built-in multimodal capabilities (audio, vision, text, and image)
* 🔌 Easily deployable across CPUs, NVIDIA, AMD, Intel GPUs and Jetson boards
* 🚫 No vendor lock-in, no cloud required, no telemetry

---

## 📦 Key Features

* 🌍 **Run Offline**: Execute LLMs, TTS, STT, and image generation locally—without any remote calls.
* 🧠 **Multi-backend Support**: Compatible with `llama.cpp`, `transformers`, `vllm`, and many others.
* 🖼️ **Multimodal Generation**: Text, audio, vision, and image generation APIs all in one platform.
* 💬 **OpenAI-Compatible API**: Seamlessly plug into existing apps using OpenAI-like endpoints.
* 🔧 **Agentic Capabilities**: Integrate with [LocalAGI](https://github.com/mudler/LocalAGI) and [LocalRecall](https://github.com/mudler/LocalRecall) for persistent memory and intelligent workflows.
* 🧩 **Modular and Extendable**: Install new backends or features dynamically via OCI containers.
* 🔒 **Privacy-First**: Keep data on-premises. No tracking. No data leaves your machine.
* 🧪 **Fast CI/CD Release Cycle**: Continuous improvement, regular releases, and vibrant community.

---

## 🧰 Quickstart

### 🐧 Run on Linux/macOS:

```bash
curl https://localai.io/install.sh | sh
```

### 🐳 Docker (CPU-Only)

```bash
docker run -ti --name local-ai -p 8080:8080 localai/localai:latest
```

### 🧠 Load Your First Model

```bash
local-ai run huggingface://TheBloke/phi-2-GGUF/phi-2.Q8_0.gguf
```

More installation options: [https://localai.io/basics/getting\_started](https://localai.io/basics/getting_started)

---

## 🖼️ Screenshots

| Chat UI                                                   | Image Generation                                           | Audio TTS                                                | Models Gallery                                               |
| --------------------------------------------------------- | ---------------------------------------------------------- | -------------------------------------------------------- | ------------------------------------------------------------ |
| ![](./docs/assets/images/screenshots/screenshot_chat.png) | ![](./docs/assets/images/screenshots/screenshot_image.png) | ![](./docs/assets/images/screenshots/screenshot_tts.png) | ![](./docs/assets/images/screenshots/screenshot_gallery.png) |

---

## 🧬 The Local Stack Family

| Project                                              | Description                                                      |
| ---------------------------------------------------- | ---------------------------------------------------------------- |
| [LocalAGI](https://github.com/mudler/LocalAGI)       | A powerful local AI orchestration platform with agentic support. |
| [LocalRecall](https://github.com/mudler/LocalRecall) | RESTful memory module for storing and retrieving context.        |

---

## 🧠 Model Compatibility

* LLaMA / LLaMA2 / LLaMA3 (GGUF)
* GPT-J / GPT-NeoX
* Falcon
* Mistral / Mixtral
* Phi-2
* Whisper (STT)
* Bark / Piper / ElevenLabs (TTS)
* Stable Diffusion / DALL-E / stablediffusion.cpp

---

## 🌐 WebUI and Integrations

* [LocalAI Frontend](https://github.com/go-skynet/LocalAI-frontend)
* [VSCode Plugin](https://github.com/badgooooor/localai-vscode-plugin)
* [LangChain Provider](https://python.langchain.com/docs/integrations/providers/localai)
* [Terminal Client (ShellOracle)](https://github.com/djcopley/ShellOracle)
* [Slack Bot](https://github.com/mudler/LocalAGI/tree/main/examples/slack)
* [Home Assistant](https://github.com/sammcj/homeassistant-localai)
* [QA-Pilot](https://github.com/reid41/QA-Pilot)

---

## 📚 Learn & Build

* 📖 [LLM Finetuning Guide](https://localai.io/docs/advanced/fine-tuning/)
* ⚙️ [Building From Source](https://localai.io/basics/build/)
* 🚢 [Kubernetes Deployment](https://localai.io/basics/getting_started/#run-localai-in-kubernetes)
* 🔁 [Distributed Inference](https://localai.io/features/distribute/)
* 🎯 [P2P Swarm + Federated Mode](https://github.com/mudler/LocalAI/pull/2723)

---

## 📣 Stay Connected

* 💬 [Discord](https://discord.gg/uJAeKSAGDy)
* 💡 [FAQ](https://localai.io/faq/)
* 🧑‍💻 [GitHub Discussions](https://github.com/go-skynet/LocalAI/discussions)
* 🐦 [@LocalAI\_API on X](https://twitter.com/LocalAI_API)
* 🎥 [YouTube Demos](https://www.youtube.com/watch?v=PKrDNuJ_dfE)
* 📬 [Telegram Bot](https://t.me/localaiofficial_bot)

---

## 💖 Contribute

LocalAI thrives because of its amazing community. You can help by:

* Reporting bugs & issues
* Submitting PRs
* Improving documentation
* Sharing your use cases & integrations

> 📌 Contribution Guide: [How to contribute](https://github.com/opencv/opencv/wiki/How_to_contribute)

---

## 🧪 CI/CD Status

[![Tests](https://github.com/go-skynet/LocalAI/actions/workflows/test.yml/badge.svg)](https://github.com/go-skynet/LocalAI/actions/workflows/test.yml)
[![Release](https://github.com/go-skynet/LocalAI/actions/workflows/release.yaml/badge.svg)](https://github.com/go-skynet/LocalAI/actions/workflows/release.yaml)
[![Image Build](https://github.com/go-skynet/LocalAI/actions/workflows/image.yml/badge.svg)](https://github.com/go-skynet/LocalAI/actions/workflows/image.yml)

---

## 🧾 License

MIT License
Created by Ettore Di Giacinto ([mudler@localai.io](mailto:mudler@localai.io))

---

## 🙏 Acknowledgements

Thanks to the open-source community and projects that made this possible:

* [llama.cpp](https://github.com/ggerganov/llama.cpp)
* [whisper.cpp](https://github.com/ggerganov/whisper.cpp)
* [bark.cpp](https://github.com/suno-ai/bark)
* [stable-diffusion.cpp](https://github.com/EdVince/Stable-Diffusion-NCNN)
* [alpaca.cpp](https://github.com/antimatter15/alpaca.cpp)

---

## 🌟 Star History

[![LocalAI Star history Chart](https://api.star-history.com/svg?repos=go-skynet/LocalAI\&type=Date)](https://star-history.com/#go-skynet/LocalAI&Date)

## ❤️ Sponsors

[Become a sponsor](https://github.com/sponsors/mudler) and support the evolution of LocalAI. Help us continue maintaining CI, hosting model galleries, and developing new features.

<p align="center">
  <a href="https://www.spectrocloud.com/" target="blank">
    <img height="200" src="https://github.com/user-attachments/assets/72eab1dd-8b93-4fc0-9ade-84db49f24962">
  </a>
  <a href="https://www.premai.io/" target="blank">
    <img height="200" src="https://github.com/mudler/LocalAI/assets/2420543/42e4ca83-661e-4f79-8e46-ae43689683d6">
  </a>
</p>

---

> Local-first. Open-source. Multimodal. The future of AI is in your hands with LocalAI.
