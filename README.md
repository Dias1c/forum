# forum
Forum with clean architecture

## How to run?
```bash
go run ./cmd/
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

## Defenition of Done
- [x] Configs system
- [x] Database schema
- [x] Connect database
- [x] Write dockerfile
- **BackEnd**
- [ ] sign-up
- [ ] sign-in
- [ ] main-page
- [ ] posts-page
- [ ] post-create-page
- [ ] post-page
- [ ] comments
- [ ] comment-create
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