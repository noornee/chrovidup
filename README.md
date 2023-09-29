### Task Description (hngx stage 5 task (2))
- Pair yourself with one or more FE devs in your team.
- You will develop the BE for their Chrome task.
- This backend will receive the video they send, save it to disk and render the page where the video can be played.
- Do not add authentication.

### chrovidup
chrovidup ->  <b>chro</b>me-extension <b>vid</b>eo <b>up</b>load :P

this essentially just uploads a video to a "server" and returns the url to the path of the stored video

### Run server locally

```bash
go run *.go
```

by default, it runs on port `8080`

```bash
# sample request

curl -X POST 'localhost:8080/api/v1/videos/upload' \
-F 'video=@"lagtrain.webm"' \
-H "Content-Type: multipart/form-data"

## sample success response
{
  "data": {
    "url": "localhost:8080/api/v1/videos/01b6cd85-695a-43fc-821d-9ffd59d63d7b.webm"
  },
  "message": "success",
  "status": 200
}

```
