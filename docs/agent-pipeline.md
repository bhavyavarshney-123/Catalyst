# AI Agent Pipeline

## Overview

Catalyst uses an AI workflow to transform raw recruiter emails into structured knowledge and generate grounded responses to user queries.

Rather than relying solely on a language model, the system combines data retrieval, semantic search, prompt engineering, and LLM reasoning into a coordinated pipeline.

This design ensures responses are based on the user's actual job application data instead of the model's internal knowledge.

---

# End-to-End Pipeline

```
                    Gmail

                      │

                      ▼

          Email Synchronization

                      │

                      ▼

          AI Metadata Extraction

                      │

                      ▼

          Structured Opportunity

                      │

                      ▼

          Embedding Generation

                      │

                      ▼

          PostgreSQL + pgvector

                      │

          ┌───────────┴────────────┐

          │                        │

          ▼                        ▼

   User Question             Reply Request

          │                        │

          └───────────┬────────────┘

                      ▼

              Embedding Generation

                      ▼

              Semantic Retrieval

                      ▼

             Context Construction

                      ▼

              Prompt Generation

                      ▼

                 OpenAI GPT

                      ▼

              Grounded Response
```

---

# Workflow Overview

The AI workflow consists of two independent pipelines:

1. Knowledge Ingestion Pipeline
2. Intelligent Query Pipeline

Separating ingestion from retrieval allows Catalyst to answer questions efficiently without repeatedly processing the same recruiter emails.

---

# Pipeline 1 — Knowledge Ingestion

This pipeline converts unstructured recruiter emails into structured semantic knowledge.

## Step 1 — Email Synchronization

```
Gmail API
      │
      ▼
 Fetch Emails
```

Catalyst authenticates with Gmail using OAuth2 and synchronizes recruiter emails.

Duplicate emails are ignored to avoid redundant processing.

---

## Step 2 — Metadata Extraction

```
Raw Email

↓

LLM

↓

Structured Opportunity
```

The AI extracts information such as:

- Company
- Job Role
- Application Status
- Interview Date
- Meeting Link
- Assessment Link
- Notes

Instead of relying on brittle regular expressions, the language model interprets recruiter emails written in natural language.

---

## Step 3 — Opportunity Creation

```
Structured Metadata

↓

Repository

↓

PostgreSQL
```

The extracted information becomes a persistent Opportunity record.

This separates application knowledge from raw email content.

---

## Step 4 — Embedding Generation

```
Opportunity

↓

OpenAI Embedding Model

↓

Vector
```

Every opportunity is converted into a dense vector representation.

These embeddings capture semantic meaning rather than exact wording.

---

## Step 5 — Vector Storage

```
Embedding

↓

pgvector
```

The embedding is stored alongside structured opportunity data inside PostgreSQL.

This allows Catalyst to perform semantic retrieval while maintaining transactional consistency.

---

# Pipeline 2 — Intelligent Query Processing

Once the knowledge base has been created, Catalyst can answer natural language questions.

---

## Step 1 — User Query

Example:

```
When is my Amazon interview?
```

---

## Step 2 — Query Embedding

```
Question

↓

Embedding Model

↓

Vector
```

The user's question is converted into the same embedding space used by stored opportunities.

This enables semantic comparison.

---

## Step 3 — Semantic Retrieval

```
Query Vector

↓

pgvector Similarity Search

↓

Relevant Opportunities
```

Rather than searching for exact keywords, Catalyst retrieves opportunities with the highest semantic similarity.

Example:

```
Amazon Backend Engineer

Amazon SDE

AWS Software Engineer
```

may all be considered relevant depending on the query.

---

## Step 4 — Context Construction

The retrieved opportunities are assembled into a structured prompt.

Example:

```
User Question

+

Relevant Opportunities

+

Recruiter Context

+

System Instructions
```

Only relevant information is included to reduce hallucinations and improve response quality.

---

## Step 5 — Response Generation

```
Prompt

↓

OpenAI GPT

↓

Answer
```

The language model generates a response using retrieved context rather than relying on its internal knowledge.

This significantly improves factual accuracy.

---

# Reply Generation Pipeline

Reply generation follows a similar workflow.

```
Reply Request

↓

Retrieve Opportunity

↓

Retrieve Recruiter Context

↓

Construct Prompt

↓

Generate Email

↓

Return Draft
```

Because replies include recruiter history and opportunity details, they remain consistent with previous interactions.

---

# Why This Architecture?

Traditional chatbots answer questions directly.

Catalyst introduces an intermediate retrieval layer.

```
Question

↓

Retrieve

↓

Ground

↓

Generate
```

This architecture provides:

- Better factual accuracy
- Reduced hallucinations
- Lower prompt size
- Better scalability
- Explainable AI workflow

---

# Design Principles

The AI pipeline follows several engineering principles.

## Separation of Concerns

Each stage has one responsibility.

- Gmail Service retrieves emails.
- Extractor structures data.
- Embedding Service generates vectors.
- Repository persists data.
- RAG Service retrieves context.
- AI Service communicates with the LLM.

---

## Retrieval Before Generation

Catalyst never generates responses without first retrieving relevant context.

This ensures responses remain grounded in user-specific information.

---

## Structured Knowledge

Rather than querying raw emails directly, Catalyst transforms recruiter communication into structured opportunities.

This enables:

- filtering
- reporting
- semantic search
- future analytics

---

## Modular Workflow

Each pipeline stage can evolve independently.

Examples:

- Replace OpenAI with another LLM provider.
- Introduce hybrid keyword + vector search.
- Add reranking.
- Add calendar integration.
- Introduce background workers.

No changes are required to unrelated components.

---

# Future Enhancements

The architecture is designed to support future AI capabilities without major refactoring.

Potential improvements include:

- Hybrid Search (Keyword + Vector)
- Cross-Encoder Reranking
- Background Synchronization
- Conversation Memory
- Multi-Agent Workflow
- Prompt Versioning
- AI Evaluation Metrics
- LLM Provider Abstraction
- Streaming Responses

---

# Key Takeaways

Catalyst is more than a chatbot.

It is an AI-powered backend system that transforms recruiter communication into structured knowledge and uses Retrieval-Augmented Generation to deliver accurate, context-aware assistance.

The separation between knowledge ingestion and intelligent retrieval allows the system to remain scalable, explainable, and easy to extend while maintaining grounded AI responses.