package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendshipService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"user_id": "12345"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 12345, "name": "Doug Williams"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipCreateParams{UserID: 12345}
	user, _, err := client.Friendships.Create(params)
	assert.Nil(t, err)
	expected := &User{ID: 12345, Name: "Doug Williams"}
	assert.Equal(t, expected, user)
}

func TestFriendshipService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"source_screen_name": "foo", "target_screen_name": "bar"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{ "relationship": { "source": { "can_dm": false, "muting": true, "id_str": "8649302", "id": 8649302, "screen_name": "foo"}, "target": { "id_str": "12148", "id": 12148, "screen_name": "bar", "following": true, "followed_by": false } } }`)
	})

	client := NewClient(httpClient)
	params := &FriendshipShowParams{SourceScreenName: "foo", TargetScreenName: "bar"}
	relationship, _, err := client.Friendships.Show(params)
	assert.Nil(t, err)
	expected := &Relationship{
		Source: RelationshipSource{ID: 8649302, ScreenName: "foo", IDStr: "8649302", CanDM: false, Muting: true, WantRetweets: false},
		Target: RelationshipTarget{ID: 12148, ScreenName: "bar", IDStr: "12148", Following: true, FollowedBy: false},
	}
	assert.Equal(t, expected, relationship)
}

func TestFriendshipService_Destroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"user_id": "12345"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 12345, "name": "Doug Williams"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipDestroyParams{UserID: 12345}
	user, _, err := client.Friendships.Destroy(params)
	assert.Nil(t, err)
	expected := &User{ID: 12345, Name: "Doug Williams"}
	assert.Equal(t, expected, user)
}
