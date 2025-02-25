# OpenTelemetry Log Sampling Demo in Go

This project demonstrates how to implement OpenTelemetry logging in Go with a simple counter application that sends logs to an OpenTelemetry Collector.

## Prerequisites

- Go 1.23.5 or later
- Docker and Docker Compose
- OpenTelemetry Collector

## Project Structure

The project consists of:
- A Go application that generates logs with a counter
- An OpenTelemetry Collector configuration
- Docker Compose setup for running the collector

## Installation

1. Clone the repository:
   ```bash
   git clone github.com/RichardBray/otel-log-sampling-go
   cd otel-log-sampling-go
   ```

2. Install dependencies:
   ```bash
   go mod downloa
This configuration:
- Receives OTLP logs via gRPC (4317) and HTTP (4318)
- Processes logs in batches
- Exports logs to debug output for demonstration

## Running the Application

1. Start the OpenTelemetry Collector:
   ```bash
   docker-compose up
   ```

2. Run the Go application:
   ```bash
   go run main.go
   ```

The application will:
- Initialize an OTLP HTTP log exporter
- Create a logger provider with batch processing
- Generate logs every second with an incrementing counter

## Code Example

The main application logic can be found in `main.go`:
d
   ```

## Configuration

The OpenTelemetry Collector is configured in `otel-collector-config.yaml`:


This configuration:
- Receives OTLP logs via gRPC (4317) and HTTP (4318)
- Processes logs in batches
- Exports logs to debug output for demonstration

## Running the Application

1. Start the OpenTelemetry Collector:
   ```bash
   docker-compose up
   ```

2. Run the Go application:
   ```bash
   go run main.go
   ```

The application will:
- Initialize an OTLP HTTP log exporter
- Create a logger provider with batch processing
- Generate logs every second with an incrementing counter
