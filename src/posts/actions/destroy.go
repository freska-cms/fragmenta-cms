package postactions

import (
	"net/http"

	"github.com/freska-cms/auth/can"
	"github.com/freska-cms/mux"
	"github.com/freska-cms/server"

	"github.com/freska-cms/freska-cms/src/lib/session"
	"github.com/freska-cms/freska-cms/src/posts"
)

// HandleDestroy responds to /posts/n/destroy by deleting the post.
func HandleDestroy(w http.ResponseWriter, r *http.Request) error {

	// Fetch the  params
	params, err := mux.Params(r)
	if err != nil {
		return server.InternalError(err)
	}

	// Find the post
	post, err := posts.Find(params.GetInt(posts.KeyName))
	if err != nil {
		return server.NotFoundError(err)
	}

	// Check the authenticity token
	err = session.CheckAuthenticity(w, r)
	if err != nil {
		return err
	}

	// Authorise destroy post
	user := session.CurrentUser(w, r)
	err = can.Destroy(post, user)
	if err != nil {
		return server.NotAuthorizedError(err)
	}

	// Destroy the post
	post.Destroy()

	// Redirect to posts root
	return server.Redirect(w, r, post.IndexURL())

}
