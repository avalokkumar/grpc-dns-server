# System Architecture

## Overview
The gRPC DNS server resolves domain names using upstream DNS servers and supports DNS-over-HTTPS (DoH). It includes caching, load balancing, and Prometheus metrics for monitoring.

## Components
1. **DNS Resolver**: Resolves domain names using upstream DNS servers.
2. **DoH Resolver**: Resolves domain names using DNS-over-HTTPS.
3. **Cache**: In-memory caching for faster responses.
4. **Load Balancer**: Distributes requests across multiple DNS servers.
5. **gRPC Server**: Exposes DNS resolution as a gRPC service.
6. **Prometheus Metrics**: Monitors query performance.