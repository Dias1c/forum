package handler

import (
	"fmt"
	"forum/architecture/models"
	"forum/internal/lg"
	"net/http"
	"strconv"
)

// PostVoteHandler -
func (m *MainHandler) PostVoteHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostVoteHandler", r)

	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	iUserId := r.Context().Value("UserId")
	if iUserId == nil {
		lg.Err.Println("PostVoteHandler: r.Context().Value(\"UserId\") is nil")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userId := iUserId.(int64)

	switch r.Method {
	case http.MethodGet:
		strPostId := r.URL.Query().Get("post_id")
		postId, err := strconv.ParseInt(strPostId, 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		strVote := r.URL.Query().Get("vote")
		vote, err := strconv.ParseInt(strVote, 3, 8)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		postVote := &models.PostVote{PostId: postId, UserId: userId, Vote: int8(vote)}
		err = m.service.PostVote.Record(postVote)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("PostVoteHandler: m.service.PostVote.Record: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// TODO: Протестируй отправляя не существующий постАйди или ЮсерАйди
		lg.Info.Printf("PostID: %d Vote: %d Err: %v\n", postId, int8(vote), err)

		http.Redirect(w, r, fmt.Sprintf("/post/get?id=%v", postId), http.StatusSeeOther)
		return
		// m.service.PostVote.GetByPostID(postId)
		// m.service.PostVote
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
