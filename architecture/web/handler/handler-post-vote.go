package handler

import (
	"forum/internal/lg"
	"net/http"
	"strconv"
)

// PostVoteHandler -
func (m *MainHandler) PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostVoteHandler", r)

	switch r.Method {
	case http.MethodGet:
		strPostId := r.URL.Query().Get("post_id")
		postId, err := strconv.ParseInt(strPostId, 10, 64)
		strVote := r.URL.Query().Get("vote")
		vote, err := strconv.ParseInt(strVote, 3, 8)
		// TODO: Протестируй отправляя не существующий постАйди или ЮсерАйди
		lg.Info.Printf("PostID: %d Vote: %d Err: %d\n", postId, vote, err)
		// m.service.PostVote.GetByPostID(postId)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
