variable "TAG" {
  default = "latest"
}

group "default" {
  targets = ["api", "app"]
}

target "api" {
  context = "../typing-server"
  tags = ["ghcr.io/su-its/typing-server:${TAG}", "ghcr.io/su-its/typing-server:latest"]
  cache_from = ["golang:1.23.4", "alpine:latest"]
  platforms = ["linux/amd64"]
}

target "app" {
  context = "../typing-app"
  tags = ["ghcr.io/su-its/typing-app:${TAG}", "ghcr.io/su-its/typing-app:latest"]
  cache_from = ["node:20.11-slim"]
  platforms = ["linux/amd64"]
}
