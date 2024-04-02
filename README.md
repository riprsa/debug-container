# 🐳🔧 A tiny debug Docker image

<img width="686" alt="image" src="https://github.com/riprsa/debug-container/assets/44776391/a8f1d801-e191-4588-b502-6cf6468f5e7e">

## Overview

Less than 10MB image for debug purposes. You can use this image to create tens of containers while exploring Traefik or similar systems.

### Setup

Docker Compose

```yml
services:
  debug_container:
    image: ghcr.io/riprsa/debug-container:latest
    ports:
      - "80:8081"
    environment:
      # PORT of inner HTTP server. Default value is ":80"
      - PORT=:8081
    # other Compose properties that you would like to use
```

By default, the container assumes its ID via `HOSTNAME`. However, you can override it. [Read more](https://docs.docker.com/network/#ip-address-and-hostname)

### Usage

You can use an HTTP client to make a request to the container. It sends an HTML page for GET requests and JSON for other methods. JSON signtaure is:

```json
{
  "container_id": "your_container_id"
}
```
