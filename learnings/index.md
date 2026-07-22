# Learnings

> Per-incident postmortems and non-obvious patterns discovered while working in this repo.

| Learning | Area | Summary |
|---|---|---|
| [init-functions-are-not-dead-code.md](init-functions-are-not-dead-code.md) | architecture, static-analysis | Go `init()` functions in this repo auto-register coin codecs; static AST analysis flags them as dead code — they are not |
