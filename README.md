# forum
Forum with clean architecture

## How to run local
Run:
```bash
go run ./cmd/
```

Run with building file
```bash
go build -o forum.exe ./cmd
./forum.exe
```

## Run on Docker
1. Build image
```bash
docker build . -t forum-image
```
2. Run container
```bash
docker run -p 80:8080 --rm --name forum-container forum-image
```
3. Check [`http://localhost`](http://localhost:80)

## For Developers
<details>
<summary>The Project's Definition of Done</summary>

> It helps to you write forum with this order

**Preparing**
- [x] Configs (config files or params)
- [x] Database schema
- [x] Create and Connect to DB
- [x] Write dockerfile

**Logic**
- [x] sign-up
- [x] sign-in
- [x] middleware (session tracker)
- [x] post-create
- [x] post-view
- [x] main-page
- [x] post-delete
- [x] post-edit-page
- [x] post-page (likes, comments)
- [x] posts-own
- [x] posts-voted
- [x] comment-create 
- [x] comment-delete
- [x] comments (likes)
- [x] posts-categories-page (filtering, posts)

**Opt**
- [ ] Makrdown
  - [ ] About project (Description)
  - [x] How To Run (Examples)
</details>
