target "docker-metadata-action" {
  tags = ["local"] # just a placeholder for local debug
}

group "default" {
  targets = ["api", "app"]
}

target "api" {
  inherits = ["docker-metadata-action"]
  annotations = [
    "index,manifest:org.opencontainers.image.title=typing-server",
    "index,manifest:org.opencontainers.image.description=API of typing game"
  ]
  context = "../typing-server"
  labels = {
    "org.opencontainers.image.title" = "typing-server"
    "org.opencontainers.image.description" = "API of typing game"
  }
  tags = [for tag in target.docker-metadata-action.tags : "ghcr.io/su-its/typing-server:${tag}"]
  cache-from = ["docker.io/library/golang:1.23.4", "docker.io/library/alpine:latest"]
  platforms = ["linux/amd64"]
}

target "app" {
  inherits = ["docker-metadata-action"]
  annotations = [
    "index,manifest:org.opencontainers.image.title=typing-app",
    "index,manifest:org.opencontainers.image.description=Web frontend of typing game"
  ]
  context = "../typing-app"
  labels = {
    "org.opencontainers.image.title" = "typing-app"
    "org.opencontainers.image.description" = "Web frontend of typing game"
  }
  tags = [for tag in target.docker-metadata-action.tags : "ghcr.io/su-its/typing-app:${tag}"]
  cache-from = ["docker.io/library/node:20.11-slim"]
  platforms = ["linux/amd64"]
}
