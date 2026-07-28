# Catalyst Architecture

## Overview

Catalyst is an AI-powered job application management platform designed to transform unstructured recruiter emails into structured opportunities and enable intelligent querying using Retrieval-Augmented Generation (RAG).

The application is built using a modular layered architecture to separate concerns between API handling, business logic, AI orchestration, and persistence.

---

# High-Level Architecture

```
                              User
                                │
               ┌────────────────┴────────────────┐
               │                                 │
          REST API                          CLI Interface
               │                                 │
               └──────────────┬──────────────────┘
                              │
                       AI Orchestrator
                              │
     ┌──────────────┬──────────┴──────────────┬──────────────┐
     │              │                         │              │
 Gmail Service   Sync Service           RAG Service    Opportunity Service
     │              │                         │              │
     └──────────────┴──────────────┬──────────┴──────────────┘
                                   │
                           Repository Layer
                                   │
                             PostgreSQL
                                   │
                               pgvector
                                   │
                              OpenAI API
```

---

# Architectural Principles

Catalyst follows several software engineering principles:

- Layered Architecture
- Dependency Injection
- Repository Pattern
- Service-Oriented Design
- Separation of Concerns
- Single Responsibility Principle
- Composition over Inheritance

Each package owns one responsibility and communicates through clearly defined interfaces.

---

# Application Layers

## Presentation Layer

Responsible for exposing Catalyst through HTTP endpoints and the CLI.

Components:

```
cmd/
internal/handlers/
```

Responsibilities:

- Receive user requests
- Validate input
- Call business services
- Return HTTP responses

---

## Service Layer

Contains the application's business logic.

```
internal/services/
```

Includes:

- Gmail Integration
- Email Synchronization
- Embedding Generation
- AI Services
- Retrieval-Augmented Generation

The service layer contains no HTTP-specific logic and can be reused across multiple interfaces.

---

## AI Layer

Responsible for all AI-related operations.

Current responsibilities include:

- Metadata Extraction
- Embedding Generation
- Semantic Retrieval
- Prompt Construction
- LLM Response Generation

This layer isolates AI logic from application logic, making it easier to evolve independently.

---

## Persistence Layer

Responsible for data storage.

```
Repository
        ↓
PostgreSQL
        ↓
pgvector
```

Responsibilities:

- CRUD operations
- Vector storage
- Similarity search
- Transaction management

The rest of the application never interacts directly with SQL.

---

# Request Flow

A typical request follows this lifecycle:

```
Client
    │
    ▼
HTTP Handler
    │
    ▼
Business Service
    │
    ▼
Repository
    │
    ▼
Database
```

AI requests introduce additional retrieval and generation steps before the final response.

---

# AI Request Flow

```
Question

↓

Embedding Generation

↓

Vector Search

↓

Relevant Opportunities

↓

Prompt Construction

↓

OpenAI GPT

↓

Grounded Response
```

The model never answers directly from the user's question.

Instead, responses are grounded using retrieved application data.

---

# Gmail Synchronization Flow

```
OAuth Login

↓

Access Token

↓

Fetch Emails

↓

Metadata Extraction

↓

Opportunity Creation

↓

Embedding Generation

↓

Store in PostgreSQL

↓

Store Vector in pgvector
```

This ensures every recruiter interaction becomes searchable using semantic similarity.

---

# Package Responsibilities

| Package | Responsibility |
|----------|----------------|
| handlers | HTTP endpoints |
| repository | Database access |
| services/gmail | Gmail integration |
| services/sync | Email synchronization |
| services/embeddings | Embedding generation |
| services/rag | Retrieval-Augmented Generation |
| services/ai | OpenAI communication |
| extractor | Metadata extraction |
| agents | AI workflow orchestration |

---

# Error Handling Strategy

Catalyst follows a fail-fast philosophy.

Errors are propagated upward until they can be translated into meaningful HTTP responses.

Responsibilities:

- Repository handles database errors.
- Services handle business logic.
- Handlers convert errors into HTTP responses.

This keeps concerns isolated.

---

# Extensibility

The architecture allows new capabilities to be added without affecting existing modules.

Examples:

- Calendar integration
- LinkedIn integration
- Resume analysis
- Interview preparation
- Multiple LLM providers

New services can be introduced while preserving the existing architecture.

---

# Benefits of the Architecture

- Modular
- Testable
- Maintainable
- AI Provider Agnostic
- Easy to Extend
- Clear Separation of Concerns
- Production-Oriented Design

The architecture emphasizes simplicity over unnecessary complexity while remaining flexible enough to support future AI capabilities.