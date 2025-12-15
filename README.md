# AI Multi-Agent System
PHASE 0 — FOUNDATION (one canonical repo + contracts)

Core goal: stop arguing about file layout. Lock the API contract.

Deliverables

Monorepo skeleton with folders: proto/, manager/, agents/, ui/, templates/, dev/, infra/.

Canonical protobuf definitions for manager↔agent (registration, start/stop, stream logs, health).

CI pipeline scaffold (lint, proto-gen, unit test job).

CONTRIBUTING.md and developer playbook.

Why first

Prevents divergent implementations and gives clear contracts for cross-language stubs.

PHASE 1 — CORE ORCHESTRATOR (manager)

Core goal: always-on process that registers agents, enforces lifecycle, routes requests.

Deliverables

Go manager daemon with gRPC server implementing proto.

Agent registry & heartbeat mechanism.

Process supervisor API (start/stop child processes).

Simple permission model (per-request capability tokens).

Local IPC discovery (unix socket or local TCP).

Acceptance criteria

Manager accepts registration from a sample agent and maintains heartbeat.

Manager can spawn a child process (sample agent) and capture stdout/stderr.

Manager exposes a health endpoint that returns ok.

Notes

Keep manager minimal: no business logic, only orchestration and safety rules.

PHASE 2 — FIRST AGENT: AUTO-PROJECT-BOOTSTRAPPER (Python)

Core goal: working end-to-end demo where user can ask UI to bootstrap a project and get a ready workspace.

Deliverables

Python bootstrapper agent implementing proto (StartBootstrap, StreamLogs, Cancel).

Template engine (Jinja2) + 2 templates (Django, FastAPI).

CLI wrapper run_agent.py that manager can invoke.

Logging to file + stdout in a way manager can stream.

Template validation + simple post-check (run devserver or tests).

Acceptance criteria

Manager can call StartBootstrap and receive accepted = true.

Manager streams logs to UI while bootstrap runs.

Resulting project starts locally (e.g., python -m uvicorn app:app) or Docker builds succeed.

“Open in VSCode” action works from UI (manager runs the command).

Agent must run in a workspace dir created by manager and must never reach outside without capability grant.

PHASE 3 — UI INTEGRATION (desktop shell + IDE integration)

Core goal: unified user experience that talks only to manager and shows progress.

Deliverables

Tauri + React desktop app (UI shell) that calls manager via REST/gRPC-web.

Project bootstrap form (stack, name, options) and real-time log viewer.

“Open in VSCode” integration via manager (manager executes code <path>).

Basic user onboarding & permission modal.

Acceptance criteria

UI can initiate bootstrap and display streaming logs in real time.

UI shows agent status and can cancel running tasks.

UI receives and displays audit logs for actions performed.

Notes

UI never executes privileged tasks directly; it only speaks to manager.

PHASE 4 — MULTI-AGENT ECOSYSTEM (scale from 1 to ~5 agents)

Core goal: add multiple, diverse agents and show orchestrator routing.

Deliverables (pick initial agents)

Code Review Agent (Python): scans repo, produces suggestions.

Scheduler Agent (Python/Go): reads tasks and suggests priorities.

Dev-Env Agent (Python + Docker): creates devcontainer + docker compose.

PR-Triage Agent (Python): tags issues and proposes PRs.

Security Auditor (Rust or Python wrapper to trivy): scans dependencies.

Acceptance criteria

Agents register with capabilities.

Router LLM (light-weight) can route a sample user prompt to the correct agent(s) and return combined result.

Manager logs show multi-agent orchestration for a composite request.

Notes

Start agents as separate processes. Use gRPC and proto messages for task exchange.

PHASE 5 — MODEL RUNTIME & ML SERVICES (embedding + small LLM)

Core goal: local embeddings, retrieval, small LLM for routing and code-gen (if on-device).

Deliverables

Embeddings service (Python) with gRPC -> Embed(text[]), Search(query).

Model runtime wrapper (Rust/C++ or existing ggml binary) exposing gRPC to manager/agents.

Simple router LLM (local small model or cloud fallback) used by manager to dispatch.

Vector store (FAISS) or disk-backed index accessible via gRPC.

Acceptance criteria

Local embedding of docs and retrieving the most-relevant chunks works.

Manager can call model runtime with timeout/priority and receive responses.

Router LLM selects appropriate agent in 90% of sample prompts.

Notes

Always run model runtime in isolated process with resource caps.

PHASE 6 — PLUGIN SYSTEM & SANDBOX (WASM + container heavy-plugins)

Core goal: allow third-party/own plugins that extend agent capabilities safely.

Deliverables

WASM plugin loader for light plugins (Rust->WASM or Tiny Go).

Container plugin spec for heavy plugins (manifest + capability request).

Plugin registry + signing model (dev mode accept unsigned).

UI to browse and enable plugins.

Acceptance criteria

Manager can load a WASM plugin, run it with limited capabilities, and unload it without restarting.

Container plugin can be launched and killed safely with resource limits.

Plugin manifest includes requested capabilities and user consent flow.

Notes

Start with local-only plugins (no network access) until signing and security are solid.

PHASE 7 — PACKAGING & INSTALLER (native experience)

Core goal: deliver an installer that sets up the manager as a system service + desktop app.

Deliverables

Cross-platform build pipeline for manager (static Go binary).

Tauri packaged desktop apps for macOS, Windows, Linux.

Installer that registers daemon/service and creates local user data folder.

Portable docker-based distribution mode.

Acceptance criteria

Installer creates a service that auto-starts manager on login.

Desktop app can detect and start/stop the manager.

Portable docker mode fully functional.

PHASE 8 — OBSERVABILITY & RESILIENCE (logs, tracing, chaos)

Core goal: ensure production-like safety — observe, test failure modes, and bounce back.

Deliverables

Local observability dashboard (logs, agent health, model metrics).

Distributed tracing (OpenTelemetry local collector).

CI integration tests: unit, contract tests, full E2E that spins manager + agents.

Chaos tests: simulate model OOM, agent crash, network cuts.

Acceptance criteria

CI runs E2E smoke flows and fails on regression.

Chaos tests reveal at least one failure mode that is fixed or has an accepted mitigation.

PHASE 9 — PRIVACY & SECURITY (hard non-negotiables)

Core goal: ship a product people can trust locally.

Deliverables

Permission model + consent dialogs.

Audit logs + rollback for destructive actions.

Secret handling via OS keychain.

Security review checklist + remediation.

Legal basics: local-only default policy, privacy notice.

Acceptance criteria

Any action that touches network, calendar, or files outside workspace triggers consent and is logged.

Secrets are not written to plaintext.

Audit logs are readable and can be used to reconstruct actions.



Building Phase 0 has started this is all that you guys need to know when workign with this codebase 