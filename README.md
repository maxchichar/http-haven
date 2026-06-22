# HTTP Haven

A series of progressive Go HTTP exercises focused on building core web server skills.

---

## About This Project

This repository contains my solutions to a set of hands-on Go exercises designed to strengthen understanding of the `net/http` package.

**These exercises were given to me by my mentor** as a structured learning path. Each exercise introduces new concepts such as routing, handling different HTTP methods, query parameters, headers, request bodies, status codes, and more.

The exercises progress from a simple ping server to more advanced endpoints involving calculations, text processing, headers, authorization, and redirects.

> This repository is great for anyone who wants to **practice building HTTP servers in Go** from the ground up using only the standard library.

---

## Exercises Overview

| Exercise | Focus Area                        | Key Features |
|----------|-----------------------------------|--------------|
| ex01     | Basic Server & Routing            | `/ping` endpoint returning "pong" |
| ex02     | Query Parameters                  | `/hello?name=...` with default value |
| ex03     | HTTP Methods & Request Body       | `/count` – word counter via POST |
| ex04     | Math Operations & Query Params    | `/calculate?op=add&a=10&b=5` |
| ex05     | Reading Headers                   | `/agent` – echoes User-Agent |
| ex06     | Authorization & Protected Routes  | `/dashboard` with API key check |
| ex07     | Redirects & Status Codes          | `/legacy` → permanent redirect to `/v2` |

---

## Tech Stack

- **Language**: Go (using only the standard library)
- **Tools**: Shell script for automated testing

---

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/maxchichar/http-haven.git
   cd http-haven
   ```

2. Go into any exercise folder (e.g.):bash
  ```bash
  cd ex01
  go run main.go
  ```

3. Run all tests at once (from root):
   ```bash
   ./test_endpoints.sh
   ```


## Learning ObjectivesSetting up HTTP servers with `net/http`

* Handling GET, POST, and different request types
* Working with query parameters, headers, and request bodies
* Proper use of HTTP status codes and redirects
* Basic input validation and error handling
* Good code organization and comments

This project is an excellent resource for **practicing Go web development** you can study the code, modify the endpoints, or use the structure as a base for your own small web servers.


## Acknowledgments

Special thanks to **my mentor** for creating and sharing these exercises. They provided a clear and practical way to level up my backend skills with Go.


## License

MIT License © 2026 Chibueze Charles Maxwell


