# ServAuth Roadmap

This roadmap outlines the planned development phases for the ServAuth identity provider service.

---

## Phase 1: Core Authentication (In Progress)
- [x] **User registration & login** — Email/password signup and login endpoints. [June 29, 2026]
- [x] **OAuth2/OIDC provider** — Authorization code flow and client credential token issuance. [June 29, 2026]
- [x] **Password reset & account lockout** — Recovery flows and lockout gates. [June 29, 2026]
- [x] **Serv-lang integration** — `auth.*` builtin helpers. [June 29, 2026]

## Phase 2: Advanced Directory Capabilities
- [x] **Multi-tenant directories** — Isolated pools. [June 29, 2026]
- [x] **Social login federation** — Google/GitHub integration. [June 29, 2026]
- [x] **MFA support** — TOTP / WebAuthn. [June 29, 2026]
- [x] **Session management** — Active session revocation. [June 29, 2026]
- [x] **API key issuance** — Scoped service accounts API keys. [June 29, 2026]
- [x] **User management API** — Admin user CRUD and session lists. [June 29, 2026]
- [x] **Secrets Envelope encryption** — KMS encryption simulator. [June 29, 2026]

## Phase 3: Production Hardening & Resilience (Completed)
- [x] **State-Isolated Unit & Validation Tests** — Implement table-driven tests for verification of key validation and lockouts. [June 30, 2026]
- [x] **Interface Abstractions & Decoupled Storage** — Extract storage access behind `UserStore` interface for mockability. [June 30, 2026]
- [x] **Structured Logging & OTel Tracing** — Add TraceMiddleware for tracing context propagation and JSON log format. [June 30, 2026]
- [x] **SIGTERM Graceful Shutdown** — Register listener to shut down HTTP listener cleanly with a 5-second timeout. [June 30, 2026]

## Phase 4: Architectural Depth (Pending)
- [ ] **Secrets Envelope Key Rotation** — Secret KMS rotation schedule & API key hashing (SEC.8)
