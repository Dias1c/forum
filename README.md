# forum
Forum with clean architecture

## How to run local
Run without build
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

## Definition of Done
- [x] Configs system
- [x] Database schema
- [x] Connect database
- [x] Write dockerfile
- **BackEnd**
- [x] sign-up
- [x] sign-in
- [x] middleware
- [x] post-create
- [x] post-delete
- [x] post-edit-page
- [~] post-page (likes, comments)
- [ ] comment-create 
- [ ] comment-delete
- [ ] comments (likes)
- [ ] categories-page
- [~] main-page (posts, posts-filters, users posts, posts-sort)
- **FrontEnd** 
- [ ] sign-up
- [ ] sign-in
- [ ] main-page
- [ ] posts-page
- [ ] post-create-page
- [ ] post-page, comments, comment-create
- [ ] categories
- [ ] filters in main page