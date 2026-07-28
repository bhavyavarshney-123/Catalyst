# Design Decisions

One of the goals of Catalyst was not only to build an AI-powered application but also to explore software architecture decisions commonly used in production backend systems.

This document explains the reasoning behind the major technical decisions made throughout the project.

---

# Why Go?

## Decision

Catalyst is implemented entirely in Go.

## Why?

The application is primarily a backend service responsible for:

- API development
- Gmail integration
- Database operations
- AI orchestration
- Concurrent processing

Go provides several advantages for these workloads:

- Excellent performance
- Simple concurrency using goroutines
- Fast compilation
- Strong standard library
- Easy deployment as a single binary
- Low operational complexity

Rather than splitting the system across multiple languages, keeping everything in Go reduces complexity while maintaining good performance and readability.

---

# Why PostgreSQL?

## Decision

PostgreSQL was selected as the primary datastore.

## Why?

Catalyst stores highly structured information:

- Companies
- Roles
- Interview dates
- Application status
- Meeting links
- Recruiter notes

A relational database provides:

- ACID transactions
- Strong consistency
- Mature indexing
- Rich querying capabilities
- Excellent ecosystem

PostgreSQL also supports extensions such as pgvector, allowing structured and semantic data to coexist within the same database.

This avoids introducing additional infrastructure solely for vector storage.

---

# Why pgvector?

## Decision

Semantic embeddings are stored using pgvector.

## Why?

Several dedicated vector databases exist, including Pinecone, Weaviate, Milvus, and Qdrant.

For Catalyst, introducing another database would increase operational complexity without providing significant benefits.

pgvector allows:

- vector similarity search
- relational joins
- transactional consistency
- simplified deployment

Both structured opportunity data and semantic embeddings live inside a single PostgreSQL instance.

---

# Why Retrieval-Augmented Generation (RAG)?

## Decision

Catalyst uses Retrieval-Augmented Generation instead of relying on the language model's internal knowledge.

## Why?

Recruiter emails and interview schedules are private, dynamic, and unavailable during model training.

Instead of expecting the LLM to know this information, Catalyst retrieves relevant application data before generating a response.

The workflow is:

```

User Question
↓
Embedding Generation
↓
Semantic Retrieval
↓
Context Construction
↓
LLM Response

```

This produces responses that are grounded in the user's own data rather than relying on the model's memory.

---

# Why Embeddings?

## Decision

Every opportunity is converted into a vector embedding.

## Why?

Traditional keyword search only matches exact words.

For example:

```

Amazon SDE

```

may not match a query such as:

```

Software Engineer at Amazon

```

Embeddings capture semantic meaning instead of literal text, allowing Catalyst to retrieve relevant opportunities even when different wording is used.

---

# Why Semantic Search?

## Decision

Catalyst performs semantic similarity search before every AI response.

## Why?

Semantic search enables natural language interaction.

Instead of requiring users to remember exact recruiter names or job titles, they can ask questions conversationally.

Examples:

- "When is my Amazon interview?"
- "Which recruiter asked me to complete an assessment?"
- "Generate a follow-up email."

The system retrieves relevant opportunities based on meaning rather than exact keywords.

---

# Why Repository Pattern?

## Decision

Database operations are isolated behind repositories.

## Why?

Repositories separate business logic from persistence logic.

Benefits include:

- cleaner services
- easier testing
- reduced SQL duplication
- easier database migrations
- improved maintainability

Business services never need to know how data is stored.

---

# Why Service-Oriented Design?

## Decision

Each major capability is implemented as its own service.

Examples:

- Gmail Service
- Sync Service
- Embedding Service
- RAG Service
- AI Service

## Why?

Each service owns a single responsibility.

This improves:

- readability
- modularity
- testability
- extensibility

New services can be introduced without affecting unrelated parts of the application.

---

# Why Dependency Injection?

## Decision

Dependencies are constructed once and injected into services.

## Why?

Rather than allowing services to create their own dependencies, objects receive everything they need through constructors.

Benefits:

- loose coupling
- easier testing
- explicit dependencies
- improved maintainability

This follows Go's preference for composition over global state.

---

# Why Docker?

## Decision

Catalyst is containerized using Docker.

## Why?

Docker provides a consistent development environment by packaging the application and its dependencies together.

Benefits include:

- reproducible development environments
- simplified onboarding
- consistent deployments
- isolation from host machine configuration

Developers can run the complete stack with minimal setup.

---

# Why AI-Assisted Metadata Extraction?

## Decision

Catalyst uses an LLM to convert recruiter emails into structured opportunity data.

## Why?

Recruiter emails vary significantly in format and wording.

Rule-based parsing would require maintaining numerous brittle regular expressions.

An LLM is able to interpret natural language and consistently extract structured fields such as:

- company
- role
- interview date
- meeting links
- application status

This makes the system more robust to different writing styles.

---

# Why an AI Orchestration Layer?

## Decision

Catalyst separates AI workflows from application logic through an orchestration layer.

## Why?

Generating a response involves multiple coordinated steps:

- retrieve context
- gather related opportunities
- construct prompts
- call the language model
- return grounded responses

Encapsulating these steps behind a dedicated orchestration layer keeps HTTP handlers and business services focused on their own responsibilities.

It also makes future AI workflows easier to extend without tightly coupling them to the rest of the application.

---

# Simplicity Over Complexity

One of the primary design goals of Catalyst was to solve a real-world problem without introducing unnecessary infrastructure.

Rather than adopting additional services simply because they are popular, each technology was selected only when it provided a clear architectural benefit.

This philosophy resulted in a system that is easier to understand, easier to maintain, and simpler to extend.

---

# Future Architectural Improvements

Potential future enhancements include:

- Background job processing
- Multi-user authentication
- Event-driven synchronization
- Multiple LLM providers
- AI workflow visualization
- Observability and tracing
- Calendar integration
- Resume analysis
- Interview preparation assistant

The current architecture is designed to accommodate these additions while preserving the existing modular structure.