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
- [~] post-create-page
- [~] post-page (likes, comments)
- [ ] posts-page (likes, categories, menu)
- [ ] comment-create 
- [ ] comments (likes)
- [ ] main-page 
- [ ] categories
- [ ] filters in main page
- **FrontEnd** 
- [ ] sign-up
- [ ] sign-in
- [ ] main-page
- [ ] posts-page
- [ ] post-create-page
- [ ] post-page, comments, comment-create
- [ ] categories
- [ ] filters in main page