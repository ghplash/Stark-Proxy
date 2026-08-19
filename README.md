# 📡 Stark-Proxy

A high-performance, concurrent, caching HTTP proxy server written from scratch in pure Go. This system is designed to intercept network traffic, manage operational memory safely, filter requests via custom firewall rules, and provide a live telemetry tracking API.

## 🚀 Key Architectural Features

- **Concurrent TCP Core:** Utilizes lightweight Go Goroutines (`go handleConnection`) to handle thousands of browser requests simultaneously with a minimal memory footprint (only 2KB per connection).
- **Highload In-Memory Cache (Sprint 2):** Custom thread-safe RAM storage powered by an LRU (Least Recently Used) cache algorithm and protected by atomic `sync.RWMutex` locks to eliminate Race Conditions.
- **Smart Firewall & Focus Timer:** Built-in network filtering middleware driven by background goroutine daemons to dynamically block distracting domains based on time schedules.
- **Microservices & Redis Persistence (Sprint 3):** Standalone persistent caching layer utilizing Redis inside Docker with auto-expiring keys (TTL).
- **Admin REST API & Telemetry:** Multi-threaded administration panel built with Gin Gonic, exposing real-time atomic metrics (requests counter, memory bytes saved).

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Networking:** Native `net`, `net/http`, `bufio` packages 
- **Concurreny Control:** Goroutines, Channels, `sync.RWMutex`, `sync/atomic` 
- **Database & Devops:** Redis, Docker Desktop 
- **API Framework:** Gin Gonic 