# REVER Returns API

A Go-based REST API simulating an eCommerce returns processing system, inspired by REVER’s SaaS platform. This project implements a series of incremental problems to build a production-ready API, showcasing Go backend skills, RESTful design, and eCommerce domain knowledge.

## Overview

This repository contains a single Go project in the `returns-api` folder, developed over a weekend to address 13 core problems (3–15) and optional stretch goals (18, 19, 20, 26, 34). Each problem adds a feature to the API, starting with a basic POST endpoint and culminating in a full-featured Returns API with SQLite storage, middleware, authentication, and webhook support. The problems align with REVER’s needs for processing returns, notifying clients, and ensuring scalability.

The commit history reflects these problems, with messages like `Add problem 3: Basic POST /returns endpoint to accept and echo JSON` corresponding to specific features. This README summarizes each problem to provide context for the commits and demonstrate the API’s evolution.

## Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/rever-returns-api.git
   cd rever-returns-api
   ```

2. **Navigate to the project**:
   ```bash
   cd returns-api
   ```

3. **Run the API**:
   ```bash
   go run main.go
   ```
   The API runs on `http://localhost:8080`.

4. **Test endpoints** (example for Problem 3):
   ```bash
   curl -X POST http://localhost:8080/returns -H "Content-Type: application/json" -d '{"order_id":"123","reason":"defective","product_id":"456"}'
   ```

## Problems Solved

Below is a list of the problems implemented, each adding a feature to the Returns API. The problems are incremental, building on previous ones to create a cohesive system. Each problem corresponds to a commit in the Git history, making it easy to track progress.

### Core Problems

1. **Problem 3: Basic Returns Endpoint**  
   - **Description**: Implements a `POST /returns` endpoint that accepts a JSON payload (e.g., `{"order_id":"123","reason":"defective","product_id":"456"}`) and echoes it back.  
   - **Purpose**: Establishes the core POST request handling for return submissions.  
   - **REVER Relevance**: Foundation for REVER’s return processing workflow.  
   - **Commit**: `Add problem 3: Basic POST /returns endpoint to accept and echo JSON`

2. **Problem 4: Simple In-Memory Storage**  
   - **Description**: Stores returns from `POST /returns` in an in-memory map with auto-generated IDs. Adds a `GET /returns/:id` endpoint to retrieve returns.  
   - **Purpose**: Introduces persistence and retrieval of return data.  
   - **REVER Relevance**: Enables tracking and querying returns, a key feature for REVER.  
   - **Commit**: `Add problem 4: In-memory storage and GET /returns/:id endpoint`

3. **Problem 5: JSON Validation**  
   - **Description**: Adds validation to `POST /returns` (e.g., `order_id` and `product_id` non-empty, `reason` max 100 chars), returning 400 for invalid input.  
   - **Purpose**: Ensures data integrity for return requests.  
   - **REVER Relevance**: Critical for maintaining valid data in REVER’s platform.  
   - **Commit**: `Add problem 5: JSON validation for POST /returns endpoint`

4. **Problem 6: Logging Middleware**  
   - **Description**: Adds middleware to log request method, path, and timestamp for all endpoints.  
   - **Purpose**: Enhances observability for debugging and monitoring.  
   - **REVER Relevance**: Supports production-grade monitoring for REVER’s APIs.  
   - **Commit**: `Add problem 6: Logging middleware for all API endpoints`

5. **Problem 7: Simple Rate Limiter**  
   - **Description**: Implements a rate limiter (e.g., 10 requests/min) for `POST /returns` using an in-memory counter.  
   - **Purpose**: Protects the API from abuse with concurrency-safe logic.  
   - **REVER Relevance**: Ensures reliability under high traffic, vital for REVER’s clients.  
   - **Commit**: `Add problem 7: Rate limiter middleware for POST /returns`

6. **Problem 8: Error Response Standardization**  
   - **Description**: Standardizes error responses (e.g., `{"error": "invalid order_id"}`) across all endpoints.  
   - **Purpose**: Improves client communication with consistent errors.  
   - **REVER Relevance**: Provides professional feedback for REVER’s SaaS users.  
   - **Commit**: `Add problem 8: Standardized JSON error responses for all endpoints`

7. **Problem 9: Goroutine-Based Logging**  
   - **Description**: Modifies logging middleware to write logs asynchronously using goroutines and channels.  
   - **Purpose**: Optimizes performance with non-blocking logging.  
   - **REVER Relevance**: Enhances scalability for REVER’s high-traffic APIs.  
   - **Commit**: `Add problem 9: Asynchronous logging with goroutines and channels`

8. **Problem 10: Basic Webhook Simulator**  
   - **Description**: Adds a `POST /webhook` endpoint to simulate notifying a client URL with return status updates.  
   - **Purpose**: Introduces external integrations via HTTP clients.  
   - **REVER Relevance**: Mimics REVER’s client notifications for return statuses.  
   - **Commit**: `Add problem 10: Webhook simulator with POST /webhook endpoint`

9. **Problem 11: Query Filtering**  
   - **Description**: Extends `GET /returns` to filter returns by status (e.g., `?status=pending`).  
   - **Purpose**: Supports flexible data retrieval for clients.  
   - **REVER Relevance**: Enables targeted queries, a standard eCommerce feature.  
   - **Commit**: `Add problem 11: Query filtering for GET /returns endpoint`

10. **Problem 12: Basic SQLite Integration**  
    - **Description**: Replaces in-memory storage with SQLite for returns (table: `id`, `order_id`, `reason`, `product_id`, `status`). Updates `POST /returns` and `GET /returns/:id`.  
    - **Purpose**: Adds production-like persistent storage.  
    - **REVER Relevance**: Ensures durable data storage for REVER’s platform.  
    - **Commit**: `Add problem 12: SQLite storage for returns API`

11. **Problem 13: Context for Timeouts**  
    - **Description**: Adds a 5-second timeout to the webhook endpoint’s HTTP client using context.  
    - **Purpose**: Manages slow external responses reliably.  
    - **REVER Relevance**: Ensures robust integrations with REVER’s partners.  
    - **Commit**: `Add problem 13: Context timeouts for webhook endpoint`

12. **Problem 14: Basic Authentication**  
    - **Description**: Adds API key authentication middleware (via `X-API-Key` header) for all endpoints.  
    - **Purpose**: Secures the API with basic authentication.  
    - **REVER Relevance**: Protects REVER’s client-facing APIs.  
    - **Commit**: `Add problem 14: API key authentication middleware`

13. **Problem 15: Returns Processing API (REVER Project 1)**  
    - **Description**: Combines all features into a full API: `POST /returns`, `GET /returns/:id`, `GET /returns?status=`, with validation, SQLite, logging, rate limiting, authentication, and webhooks. Includes polished code and documentation.  
    - **Purpose**: Delivers a production-ready API showcasing all skills.  
    - **REVER Relevance**: Represents REVER’s core returns platform, demonstrating scalability and professionalism.  
    - **Commit**: `Add problem 15: Full Returns Processing API with all features`

### Stretch Goal Problems (Optional)

These problems enhance the API for extra portfolio polish, implemented only if time allows.

1. **Problem 18: Batch Return Processing**  
   - **Description**: Adds `POST /returns/batch` to process multiple returns in one request.  
   - **Purpose**: Supports bulk operations for efficiency.  
   - **REVER Relevance**: Handles high-volume returns, common in eCommerce.  
   - **Commit**: `Add problem 18: Batch return processing with POST /returns/batch`

2. **Problem 19: Retry Mechanism**  
   - **Description**: Adds retries to the webhook endpoint for failed HTTP calls with exponential backoff.  
   - **Purpose**: Builds resilient integrations.  
   - **REVER Relevance**: Ensures reliable notifications for REVER’s clients.  
   - **Commit**: `Add problem 19: Retry mechanism for webhook endpoint`

3. **Problem 20: Pagination for Returns**  
   - **Description**: Adds pagination to `GET /returns` (e.g., `?page=1&limit=10`).  
   - **Purpose**: Enables efficient data retrieval for large datasets.  
   - **REVER Relevance**: Supports scalable queries for REVER’s clients.  
   - **Commit**: `Add problem 20: Pagination for GET /returns endpoint`

4. **Problem 26: API Documentation**  
   - **Description**: Generates OpenAPI/Swagger documentation using comments or `swag`.  
   - **Purpose**: Provides client-friendly API specs.  
   - **REVER Relevance**: Enhances professionalism for REVER’s API users.  
   - **Commit**: `Add problem 26: OpenAPI/Swagger documentation for API`

5. **Problem 34: Dockerized API**  
   - **Description**: Containerizes the API with Docker and `docker-compose.yml`.  
   - **Purpose**: Prepares the API for deployment.  
   - **REVER Relevance**: Demonstrates DevOps skills for production readiness.  
   - **Commit**: `Add problem 34: Dockerized API with Dockerfile and docker-compose`

## Progress Tracking

The Git commit messages (e.g., `Add problem 3: Basic POST /returns endpoint to accept and echo JSON`) correspond to the problems above, allowing viewers to follow the API’s development. Major milestones (e.g., Problem 15) are tagged (e.g., `v1.0-returns-api`) and released on GitHub for portfolio visibility. See the `returns-api/README.md` for detailed implementation notes per problem.

## Future Improvements

If time permits, additional features like advanced validation, metrics, or Kubernetes deployment could further enhance the API, aligning with REVER’s scalability needs.

## Contact

For questions or feedback, reach out via [GitHub Issues](https://github.com/your-username/rever-returns-api/issues) or [your-email@example.com].