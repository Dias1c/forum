package handler

import (
	"errors"
	"net/http"

	"forum/architecture/models"
	"forum/architecture/web/handler/cookies"
	"forum/architecture/web/handler/view"
	"forum/internal/lg"

	spostvote "forum/architecture/service/post_vote"
	ssession "forum/architecture/service/session"
)

// IndexHandler -
func (m *MainHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("IndexHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Logic
	cookies.RemoveRedirectCookie(w, r)
	switch r.Method {
	case http.MethodGet:
		cookie := cookies.GetSessionCookie(w, r)
		if cookie == nil {
			pg := getIndexPage(m, nil)
			m.view.ExecuteTemplate(w, pg, "home.html")
			return
		}

		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			pg := getIndexPage(m, nil)
			m.view.ExecuteTemplate(w, pg, "home.html")
			return
		case err != nil:
			lg.Err.Printf("IndexHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		user, err := m.service.User.GetByID(session.UserId)
		switch {
		case err == nil:
		case err != nil:
			lg.Err.Printf("IndexHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		pg := getIndexPage(m, user)
		m.view.ExecuteTemplate(w, pg, "home.html")
		return
	}
}

func getIndexPage(m *MainHandler, user *models.User) *view.Page {
	posts, err := m.service.Post.GetAll(0, 0)
	if err != nil {
		lg.Err.Printf("getIndexPage: m.service.Post.GetAll: %v\n", err)
	}
	for i := 0; i < len(posts); i++ {
		posts[i].WCategories, err = m.service.PostCategory.GetByPostID(posts[i].Id)
		switch {
		case err != nil:
			lg.Err.Printf("getIndexPage: m.service.PostCategory.GetByPostID(postId: %v): %v", posts[i].Id, err)
		}

		posts[i].WUser, err = m.service.User.GetByID(posts[i].UserId)
		switch {
		case err != nil:
			lg.Err.Printf("getIndexPage: m.service.User.GetByID(userId: %v): %v", posts[i].UserId, err)
		}

		vUp, vDown, err := m.service.PostVote.GetByPostID(posts[i].Id)
		switch {
		case err != nil:
			lg.Err.Printf("getIndexPage: m.service.PostVote.GetByPostID(id: %v): %v", posts[i].Id, err)
		}
		posts[i].WVoteUp = vUp
		posts[i].WVoteDown = vDown

		if user == nil {
			continue
		}

		vUser, err := m.service.PostVote.GetPostUserVote(user.Id, posts[i].Id)
		switch {
		case err == nil:
			posts[i].WUserVote = vUser.Vote
		case errors.Is(err, spostvote.ErrNotFound):
		case err != nil:
			lg.Err.Printf("getIndexPage: m.service.PostVote.GetPostUserVote(userId: %v, postId: %v): %v", user.Id, posts[i].Id, err)
		}
	}
	return &view.Page{User: user, Posts: posts}
}
