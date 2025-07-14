# LocalAI: Empowering Your AI Journey – Locally and Open Source 🚀

<p align="center">
<img width="300" src="https://raw.githubusercontent.com/go-skynet/LocalAI/master/core/http/static/logo.png" alt="LocalAI Logo"> <br>
<br>
</p>

<p align="center">
<a href="https://github.com/go-skynet/LocalAI/fork" target="_blank">
<img src="https://img.shields.io/github/forks/go-skynet/LocalAI?style=for-the-badge" alt="LocalAI forks"/>
</a>
<a href="https://github.com/go-skynet/LocalAI/stargazers" target="_blank">
<img src="https://img.shields.io/github/stars/go-skynet/LocalAI?style=for-the-badge" alt="LocalAI stars"/>
</a>
<a href="https://github.com/go-skynet/LocalAI/pulls" target="_blank">
<img src="https://img.shields.io/github/issues-pr/go-skynet/LocalAI?style=for-the-badge" alt="LocalAI pull-requests"/>
</a>
<a href="https://github.com/go-skynet/LocalAI/releases" target="_blank">
<img src="https://img.shields.io/github/release/go-skynet/LocalAI?&label=Latest&style=for-the-badge">
</a>
</p>

<p align="center">
<a href="https://hub.docker.com/r/localai/localai" target="_blank">
<img src="https://img.shields.io/badge/dockerhub-images-important.svg?logo=Docker" alt="LocalAI Docker hub"/>
</a>
<a href="https://quay.io/repository/go-skynet/local-ai?tab=tags&tag=latest" target="_blank">
<img src="https://img.shields.io/badge/quay.io-images-important.svg" alt="LocalAI Quay.io"/>
</a>
</p>

<p align="center">
<a href="https://twitter.com/LocalAI_API" target="_blank">
<img src="https://img.shields.io/badge/X-%23000000.svg?style=for-the-badge&logo=X&logoColor=white&label=LocalAI_API" alt="Follow LocalAI_API"/>
</a>
<a href="https://discord.gg/uJAeKSAGDy" target="_blank">
<img src="https://dcbadge.vercel.app/api/server/uJAeKSAGDy?style=flat-square&theme=default-inverted" alt="Join LocalAI Discord Community"/>
</a>
</p>

<p align="center">
<a href="https://trendshift.io/repositories/5539" target="_blank"><img src="https://trendshift.io/api/badge/repositories/5539" alt="mudler%2FLocalAI | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>
</p>

> \:bulb: **Get Assistance!** Find answers in our [❓FAQ](https://localai.io/faq/), engage in [💭Discussions](https://github.com/go-skynet/LocalAI/discussions), connect on [:speech\_balloon: Discord](https://discord.gg/uJAeKSAGDy), or delve into our comprehensive [:book: Documentation website](https://localai.io/).
>
> **Quick Links:** [💻 Quickstart Guide](https://localai.io/basics/getting_started/) | [🖼️ Model Gallery](https://models.localai.io/) | [🚀 Roadmap](https://github.com/mudler/LocalAI/issues?q=is%3Aissue+is%3Aopen+label%3Aroadmap) | [🥽 Live Demo](https://demo.localai.io) | [🌍 Model Explorer](https://explorer.localai.io) | [🛫 Practical Examples](https://github.com/mudler/LocalAI-examples) | Try it on [Telegram Bot](https://t.me/localaiofficial_bot)

---

[](https://github.com/go-skynet/LocalAI/actions/workflows/test.yml)
[](https://github.com/go-skynet/LocalAI/actions/workflows/release.yaml)
[](https://github.com/go-skynet/LocalAI/actions/workflows/image.yml)
[](https://github.com/go-skynet/LocalAI/actions/workflows/bump_deps.yaml)
[](https://artifacthub.io/packages/search?repo=localai)

**LocalAI** stands as the **free, open-source alternative to OpenAI**, providing a powerful and versatile REST API for local AI inferencing. Designed to be a drop-in replacement compatible with OpenAI, Elevenlabs, and Anthropic API specifications, LocalAI empowers you to run a diverse range of **Large Language Models (LLMs)**, **generate images**, and **create audio** right on your own machine or on-premise infrastructure. What makes LocalAI truly remarkable is its ability to operate effectively with **consumer-grade hardware**, eliminating the need for a dedicated GPU for many tasks. This innovative project is proudly created and maintained by [Ettore Di Giacinto](https://github.com/mudler).

---

## 📚🆕 Introducing the Local Stack Family: A Unified AI Ecosystem

LocalAI is no longer a standalone solution; it's now a core component of a comprehensive suite of interconnected AI tools, meticulously designed to work together seamlessly and amplify your AI capabilities:

| Project                                                  | Logo                                                                                                                                                                                                      | Description                                                                                                                                                                                 |
| -------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **[LocalAGI](https://github.com/mudler/LocalAGI)**       | <a href="https://github.com/mudler/LocalAGI"><img src="https://raw.githubusercontent.com/mudler/LocalAGI/refs/heads/main/webui/react-ui/public/logo_2.png" width="200" alt="LocalAGI Logo"></a>           | A robust Local AI agent management platform that functions as a drop-in replacement for OpenAI's Responses API, enhanced with advanced agentic capabilities for sophisticated AI workflows. |
| **[LocalRecall](https://github.com/mudler/LocalRecall)** | <a href="https://github.com/mudler/LocalRecall"><img src="https://raw.githubusercontent.com/mudler/LocalRecall/refs/heads/main/static/localrecall_horizontal.png" width="200" alt="LocalRecall Logo"></a> | A powerful REST-ful API and knowledge base management system that provides persistent memory and storage capabilities, offering a crucial long-term memory solution for your AI agents.     |

---

## 📸 Visual Showcase: Explore LocalAI in Action

Witness the intuitive interfaces and powerful functionalities of LocalAI through these engaging screenshots:

| Talk Interface          | Generate Audio          |
| ----------------------- | ----------------------- |
| *(Screenshots go here)* | *(Screenshots go here)* |

| Models Overview         | Generate Images         |
| ----------------------- | ----------------------- |
| *(Screenshots go here)* | *(Screenshots go here)* |

| Chat Interface          | Home                    |
| ----------------------- | ----------------------- |
| *(Screenshots go here)* | *(Screenshots go here)* |

| Login                   | Swarm                   |
| ----------------------- | ----------------------- |
| *(Screenshots go here)* | *(Screenshots go here)* |

---

## 💻 Get Started Quickly: Your Journey with LocalAI

Launching LocalAI is straightforward, whether you prefer an installer script or Docker containers.

### Effortless Installation

For a basic setup, simply run our convenient installer script:

```bash
# Basic installation for quick deployment
curl https://localai.io/install.sh | sh
```

Looking for more tailored installation procedures? Explore our [Installer Options](https://localai.io/docs/advanced/installer/) for advanced configurations.

### Docker Deployments: Versatility at Your Fingertips

LocalAI provides a variety of Docker images optimized for different hardware configurations:

#### **CPU-Only Image:**

Ideal for environments without dedicated GPUs.

```bash
docker run -ti --name local-ai -p 8080:8080 localai/localai:latest
```

#### **NVIDIA GPU Images:**

Leverage the power of NVIDIA GPUs for accelerated inferencing.

* **CUDA 12.0:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --gpus all localai/localai:latest-gpu-nvidia-cuda-12
  ```
* **CUDA 11.7:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --gpus all localai/localai:latest-gpu-nvidia-cuda-11
  ```
* **NVIDIA Jetson (L4T) ARM64:**
  Optimized for NVIDIA Jetson platforms.

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --gpus all localai/localai:latest-nvidia-l4t-arm64
  ```

#### **AMD GPU Images (ROCm):**

For systems equipped with AMD GPUs utilizing ROCm.

```bash
docker run -ti --name local-ai -p 8080:8080 --device=/dev/kfd --device=/dev/dri --group-add=video localai/localai:latest-gpu-hipblas
```

#### **Intel GPU Images (oneAPI):**

Harness the capabilities of Intel GPUs with oneAPI support.

* **Intel GPU with FP16 support:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --device=/dev/dri/card1 --device=/dev/dri/renderD128 localai/localai:latest-gpu-intel-f16
  ```
* **Intel GPU with FP32 support:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --device=/dev/dri/card1 --device=/dev/dri/renderD128 localai/localai:latest-gpu-intel-f32
  ```

#### **Vulkan GPU Images:**

Utilize Vulkan for cross-platform GPU acceleration.

```bash
docker run -ti --name local-ai -p 8080:8080 localai/localai:latest-gpu-vulkan
```

#### **All-in-One (AIO) Images (Pre-downloaded Models):**

Get started even faster with AIO images that include pre-downloaded models.

* **CPU version:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 localai/localai:latest-aio-cpu
  ```
* **NVIDIA CUDA 12 version:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --gpus all localai/localai:latest-aio-gpu-nvidia-cuda-12
  ```
* **NVIDIA CUDA 11 version:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --gpus all localai/localai:latest-aio-gpu-nvidia-cuda-11
  ```
* **Intel GPU version:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 localai/localai:latest-aio-gpu-intel-f16
  ```
* **AMD GPU version:**

  ```bash
  docker run -ti --name local-ai -p 8080:8080 --device=/dev/kfd --device=/dev/dri --group-add=video localai/localai:latest-aio-gpu-hipblas
  ```

For more in-depth information about our AIO images and the pre-downloaded models they contain, please consult our [Container Documentation](https://localai.io/basics/container/).

---

## 🧠 Seamless Model Loading

LocalAI provides flexible options for loading AI models:

```bash
# Load a model directly from the LocalAI model gallery
local-ai models load <model-name>
```

You can also specify model files locally or remotely when starting the server:

```bash
local-ai --model-path /path/to/your/model.bin
```

For detailed instructions, visit [Model Management](https://localai.io/basics/models/).

---

## 🔧 API Compatibility & Usage

LocalAI is designed to be API-compatible with:

* OpenAI's API
* ElevenLabs API
* Anthropic API

You can switch your API base URL to point to LocalAI and use your existing client code seamlessly.

---

## 📖 Learn More

Dive into our full [Documentation](https://localai.io/docs/) for advanced usage, deployment guides, API references, and community resources.

---

Thank you for using LocalAI — the free, open-source AI API designed for you, on your machine, your data, your privacy.
