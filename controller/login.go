package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/riankrishandi/dans/auth"
	"github.com/riankrishandi/dans/repo"
	"golang.org/x/crypto/bcrypt"
)

var (
	errUserNotFound        = errors.New("user not found")
	errUserInvalidPassword = errors.New("invalid password")
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Controller.HandleLogin.
func (c *Controller) HandleLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("[controller.HandleLogin] failed to read request body: %s\n", err.Error())
			c.renderer.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		var userReq User
		err = json.Unmarshal(bytes, &userReq)
		if err != nil {
			log.Printf("[controller.HandleLogin] failed to unmarshal json: %s\n", err.Error())
			c.renderer.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		// Check if user exists.
		userRepoRes, err := c.repo.GetUser(repo.GetUserRepoParam{
			Ctx:      r.Context(),
			Username: userReq.Username,
		})
		if err != nil {
			c.renderer.RenderJSON(w, http.StatusFailedDependency, err)
			return
		} else if userRepoRes == nil {
			c.renderer.RenderJSON(w, http.StatusNotFound, errUserNotFound)
			return
		}

		// Check if password valid.
		valid := checkHashPassword(userReq.Password, userRepoRes.User.Password)
		if !valid {
			c.renderer.RenderJSON(w, http.StatusBadRequest, errUserInvalidPassword)
			return
		}

		token, err := auth.EncodeToken(auth.CustomClaims{
			ID:       strconv.Itoa(userRepoRes.User.ID),
			Username: userRepoRes.User.Username,
		})
		if err != nil {
			c.renderer.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		c.renderer.RenderJSON(w, http.StatusOK, map[string]interface{}{
			"token": token,
		})
	})
}

func checkHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
