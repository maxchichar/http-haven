# Exercise 7: Simple Redirector (Status Codes)

**Goal:** Create a `/legacy` route. Whenever a user hits this endpoint, permanently redirect them to a new route `/v2` with a friendly "Welcome to version 2" message.

**Key Tasks:**
Redirect traffic using the `http.Redirect` helper function.
Use the proper status code for a permanent move (`http.StatusMovedPermanently`).
