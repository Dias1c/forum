package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dias1c/forum/architecture/models"
	"github.com/Dias1c/forum/architecture/service/post_comment_vote"
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

		var userId int64
		if user != nil {
			userId = user.Id
		}
		m.service.FillPost(post, userId)

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
			lg.Err.Printf("PostViewHandler: service.PostComment.GetAllByPostID: %v\n", err)
		}

		for _, comment := range post.WComments {
			comment.WVoteUp, comment.WVoteDown, err = m.service.PostCommentVote.GetByCommentID(comment.Id)
			if err != nil {
				lg.Err.Printf("PostViewHandler: service.PostCommentVote.GetByCommentID(commentId: %v): %v\n", comment.Id, err)
			}
			vt, err := m.service.PostCommentVote.GetCommentUserVote(userId, comment.Id)
			switch {
			case err == nil:
				comment.WUserVote = vt.Vote
			case errors.Is(err, post_comment_vote.ErrNotFound):
			case err != nil:
				lg.Err.Printf("PostViewHandler: service.PostCommentVote.GetCommentUserVote: %v\n", err)
			}
		}

		pg := &view.Page{User: user, Post: post}
		m.view.ExecuteTemplate(w, pg, "post-view.html")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
