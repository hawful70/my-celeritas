# Celeritas

**Celeritas** is a Laravel-inspired web framework written in Go, designed to bring high performance, type safety, and developer productivity together. While maintaining Go's simplicity and speed, Celeritas recreates many of Laravel’s most useful features, empowering developers to build robust web applications with ease.

## Description

This project is about implementing Laravel-style components using idiomatic Go. Unlike PHP, Go is statically typed and compiled, resulting in applications that are significantly faster and more resilient. Celeritas focuses on giving developers familiar, powerful tools while leveraging Go's concurrency, performance, and simplicity.

## Features

- **Eloquent-like ORM**  
  A database-agnostic Object-Relational Mapper with fluent query building and relational data modeling.

- **Database Migration System**  
  Structured, version-controlled schema changes with rollback and replay capabilities.

- **User Authentication**  
  - Session-based authentication for traditional web apps  
  - Token-based authentication for APIs  
  - Built-in password reset system  
  - One-command installer

- **Templating System**  
  Supports both Go's standard templates and Jet templates for component-based UI rendering.

- **Caching Layer**  
  Unified caching API with support for Redis and Badger backends.

- **Session Management**  
  Session storage in cookies, databases (MySQL/PostgreSQL), or Redis.

- **Flexible Response Formats**  
  Easily return HTML, JSON, XML, or downloadable files from handlers.

- **Validation**  
  Built-in tools for validating form inputs and JSON data.

- **Mailing System**  
  Email sending via SMTP and API-based providers including MailGun, SparkPost, and SendGrid.

- **Command Line Interface (CLI)**  
  Tools to scaffold:
  - Models
  - Handlers
  - Email templates  
  Includes a single-command project generator:  
  `celeritas new <myproject>`

---

This project is ideal for developers who love Laravel’s ergonomics and want to bring that development experience to Go.

