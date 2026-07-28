# Catalyst

> An AI-powered job application management platform built in Go that transforms recruiter emails into structured opportunities, semantic knowledge, and actionable insights.

Catalyst automates the tedious process of tracking job applications by integrating directly with Gmail, extracting recruiter information using LLMs, storing semantic embeddings with pgvector, and enabling Retrieval-Augmented Generation (RAG) for intelligent career assistance.

Instead of manually searching through hundreds of recruiter emails, Catalyst allows users to ask natural language questions such as:

> *"When is my Amazon interview?"*

> *"Generate a reply to the recruiter asking to reschedule the interview."*

> *"Which companies haven't responded in the last two weeks?"*

---

# ✨ Features

### Gmail Integration

- OAuth2 authentication
- Fetch recruiter emails
- Incremental synchronization
- Automatic duplicate prevention

---

### AI Metadata Extraction

Catalyst uses an LLM to extract structured information from unstructured recruiter emails, including:

- Company Name
- Job Role
- Application Status
- Interview Date
- Meeting Links
- Assessment Links
- Recruiter Notes

---

### Opportunity Management

Complete CRUD operations for managing job opportunities.

Each opportunity is enriched using information extracted directly from recruiter emails.

---

### Semantic Search

Every opportunity is converted into a vector embedding and stored inside PostgreSQL using **pgvector**.

Instead of keyword search, Catalyst performs semantic retrieval to find contextually relevant opportunities.

---

### Retrieval-Augmented Generation (RAG)

Catalyst combines semantic retrieval with GPT to answer natural language questions.

Example:

```
User:

"When is my Microsoft interview?"
```

The system:

- retrieves relevant opportunities
- retrieves recruiter context
- builds a prompt
- generates a grounded response

---

### AI Email Reply Generation

Generate recruiter replies using the retrieved context.

Instead of generic responses, replies are grounded using:

- recruiter history
- interview information
- opportunity status

---

# 🏗️ System Architecture

```
                    Browser / CLI
                          │
                          │
                    HTTP API (Chi)
                          │
        ┌─────────────────┼─────────────────┐
        │                 │                 │
        │                 │                 │
 Gmail Service      Opportunity Service    AI Engine
        │                 │                 │
        └─────────────────┼─────────────────┘
                          │
                    PostgreSQL
                          │
                      pgvector
                          │
                  Semantic Retrieval
                          │
                    Prompt Builder
                          │
                     OpenAI GPT
```

---

# 🧠 AI Pipeline

```
Recruiter Email
       │
       ▼
 Gmail Synchronization
       │
       ▼
 Metadata Extraction
       │
       ▼
 Opportunity Creation
       │
       ▼
 Embedding Generation
       │
       ▼
 pgvector Storage
       │
       ▼
 Semantic Retrieval
       │
       ▼
 Prompt Construction
       │
       ▼
 GPT Response
```

---

# 📂 Project Structure

```
cmd/
    cli/                # Interactive command line interface

internal/

    agents/             # AI workflow orchestration

    database/           # Database connection

    extractor/          # Metadata extraction

    handlers/           # HTTP handlers

    models/             # Domain models

    repository/         # PostgreSQL repository layer

    services/

        ai/             # OpenAI integration

        embeddings/     # Embedding generation

        gmail/          # Gmail API

        rag/            # Retrieval-Augmented Generation

        sync/           # Email synchronization
```

---

# 🚀 How It Works

### 1. Connect Gmail

Users authenticate using Gmail OAuth.

↓

### 2. Synchronize Emails

Recruiter emails are downloaded.

↓

### 3. Extract Metadata

LLMs convert unstructured emails into structured opportunities.

↓

### 4. Generate Embeddings

Each opportunity is converted into a semantic vector.

↓

### 5. Store in PostgreSQL

Vectors are stored using pgvector.

↓

### 6. Ask Questions

Users ask questions in natural language.

↓

### 7. Semantic Retrieval

Relevant opportunities are retrieved.

↓

### 8. GPT Generates Response

Grounded responses are generated using retrieved context.

---

# 🛠 Tech Stack

## Backend

- Go
- Chi Router

## Database

- PostgreSQL
- pgvector

## AI

- OpenAI GPT
- OpenAI Embeddings

## Email

- Gmail API
- OAuth2

## Infrastructure

- Docker

---

# 💡 Example Queries

```
When is my Amazon interview?
```

```
Generate a reply asking to reschedule my interview.
```

```
Which recruiters have not responded recently?
```

```
Show opportunities currently in interview stage.
```

---

# 🎯 Why Catalyst?

Recruiter conversations are often scattered across hundreds of emails.

Traditional job trackers require manual updates and keyword searching.

Catalyst uses semantic search and LLMs to automatically organize recruiter communications, maintain structured opportunity records, and provide intelligent assistance throughout the job application process.

---

# 🔮 Future Improvements

- Multi-user support
- Background email synchronization
- Resume optimization
- Calendar integration
- Vector-based recruiter memory
- AI interview preparation
- Personalized career insights

---

# 📹 Demo

Coming Soon

---

# 👨‍💻 Author

**Bhavya Varshney**

Backend Engineer | Golang | AI Systems | RAG | PostgreSQL | Docker