package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/service/post_vote"
	"github.com/Dias1c/forum/architecture/web/handler/cookies"
	"github.com/Dias1c/forum/architecture/web/handler/view"
	"github.com/Dias1c/forum/internal/lg"

	spost "github.com/Dias1c/forum/architecture/service/post"
	ssession "github.com/Dias1c/forum/architecture/service/session"
)

// PostCreateHandler -
func (m *MainHandler) PostViewHandler(w http.ResponseWriter, r *http.Request) {
	debugLogHandler("PostViewHandler", r)

	// Allowed Methods
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var user *models.User
	cookie := cookies.GetSessionCookie(w, r)
	switch cookie {
	case nil:
		user = nil
	default:
		session, err := m.service.Session.GetByUuid(cookie.Value)
		switch {
		case err == nil:
			user, _ = m.service.User.GetByID(session.UserId)
		case errors.Is(err, ssession.ErrExpired) || errors.Is(err, ssession.ErrNotFound):
			cookies.RemoveSessionCookie(w, r)
		default:
			lg.Err.Printf("PostViewHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}
	}

	switch r.Method {
	case http.MethodGet:
		strPostId := r.URL.Query().Get("id")
		postId, err := strconv.ParseInt(strPostId, 10, 64)
		if err != nil || postId < 1 {
			http.Error(w, "Invalid query id", http.StatusBadRequest)
			return
		}

		post, err := m.service.Post.GetByID(postId)
		switch {
		case err == nil:
		case errors.Is(err, spost.ErrNotFound):
			// TODO: error page
			http.Error(w, "Post Not Found", http.StatusNotFound)
			return
		}

		categories, err := m.service.Category.GetByPostID(post.Id)
		switch {
		case err == nil:
			post.WCategories = categories
		default:
			lg.Err.Printf("PostViewHandler: m.service.Session.GetByUuid: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		up, down, err := m.service.PostVote.GetByPostID(postId)
		switch {
		case err == nil:
			post.WVoteUp = up
			post.WVoteDown = down
		case err != nil:
			lg.Err.Printf("PostViewHandler: m.service.PostVote.GetByPostID: %v\n", err)
			http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
			return
		}

		if user != nil {
			usrVote, err := m.service.PostVote.GetPostUserVote(user.Id, post.Id)
			switch {
			case err == nil:
				post.WUserVote = usrVote.Vote
			case errors.Is(err, post_vote.ErrNotFound):
			case err != nil:
				lg.Err.Printf("PostViewHandler: m.service.PostVote.GetPostUserVote: %v\n", err)
				http.Error(w, "something wrong, maybe try again later", http.StatusInternalServerError)
				return
			}
		}

		post.WUser, err = m.service.User.GetByID(post.UserId)
		switch {
		case err == nil:
		default:
			lg.Err.Printf("PostViewHandler: m.service.User.GetByID(userId: %v): %v\n", post.UserId, err)
		}

		post.WComments, err = m.service.PostComment.GetAllByPostID(post.Id, 0, models.SqlLimitInfinity)
		switch {
		case err == nil:
			for _, comment := range post.WComments {
				comment.WUser, err = m.service.User.GetByID(comment.UserId)
				if err != nil {
					lg.Err.Printf("PostViewHandler: m.service.User.GetByID: %w", err)
				}
			}
		case err != nil:
			lg.Err.Printf("PostViewHandler: m.service.PostComment.GetAllByPostID: %v\n", err)
		}

		for _, comment := range post.WComments {
			comment.WVoteUp, comment.WVoteDown, err = m.service.PostCommentVote.GetByCommentID(comment.Id)
			if err != nil {
				lg.Err.Printf("PostViewHandler: m.service.PostCommentVote.GetByCommentID(commentId: %v): %v\n", comment.Id, err)
			}
			vt, err := m.service.PostCommentVote.GetCommentUserVote(user.Id, comment.Id)
			switch {
			case err == nil:
				comment.WUserVote = vt.Vote
			case err != nil:
				lg.Err.Printf("PostViewHandler: m.service.PostCommentVote.GetCommentUserVote: %v\n", err)
			}
		}

		pg := &view.Page{User: user, Post: post}
		m.view.ExecuteTemplate(w, pg, "post-view.html")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
