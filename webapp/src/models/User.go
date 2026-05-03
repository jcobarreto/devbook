package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Nick       string    `json:"nick"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Followers  []User    `json:"followers"`
	Following  []User    `json:"following"`
	Posts      []Post    `json:"posts"`
}

// GetUserProfile retrieves the profile of a user by their ID
func GetUserProfile(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go getUserData(userChannel, userID, r)
	go getFollowers(followersChannel, userID, r)
	go getFollowing(followingChannel, userID, r)
	go getPosts(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("failed to load user data")
			}
			user = userLoaded

		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				return User{}, errors.New("failed to load followers")
			}
			followers = followersLoaded

		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				return User{}, errors.New("failed to load following")
			}
			following = followingLoaded

		case postsLoaded := <-postsChannel:
			if postsLoaded == nil {
				return User{}, errors.New("failed to load posts")
			}
			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// getUserData retrieves the basic data of a user by their ID
func getUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// getFollowers retrieves the followers of a user by their ID
func getFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// getFollowing retrieves the users that a user is following by their ID
func getFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

// getPosts retrieves the posts of a user by their ID
func getPosts(channel chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
