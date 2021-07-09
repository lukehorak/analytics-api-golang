# Work Sample for Product Role, Golang Variant

This project was a submission for a take-home assignment, the prompt for which is included below.

## Assignment Prompt

#### 1. Data structure

Implement counters to support event tracking (views and clicks) by content selection and time. Example counter: Key `"sports:2020-01-08 22:01"`, Value `{views: 100, clicks: 4}`.

#### 2. Mock store

Implement a mock store for storing counters. It can be in-memory, filesystem-based, or satellite-based (satellite not provided). The content of the store is to be queried and returned by the `stats` handler.

#### 3. Goroutine

Create a goroutine to upload counters to the mock store every 5 seconds.

#### 4. Global rate-limiting

Implement a global (not per-client) rate limit for the `stats` handler.

