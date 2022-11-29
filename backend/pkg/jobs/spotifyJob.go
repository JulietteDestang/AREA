/** @file spotifyJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Spotify API
 * @author Timothee de Boynes
 * 
 */

package jobs

// @cond
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"os"

	"github.com/tidwall/gjson"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief Refresh spotify token in our DB if it is outdated
 * @param userID uint
 */
func RefreshSpotifyToken(userID uint) {

	userToken := *models.FindUserToken(userID)
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	refreshurl := "https://accounts.spotify.com/api/token"

	refreshData := url.Values{}
	refreshData.Set("refresh_token", userToken.SpotifyRefreshToken)
	refreshData.Set("grant_type", "refresh_token")
	refreshData.Set("client_id", utils.GetEnv("SPOTIFY_ID"))
	refreshEncodedData := refreshData.Encode()

	refreshreq, _ := http.NewRequest("POST", refreshurl, strings.NewReader(refreshEncodedData))
	refreshreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	refreshreq.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte((utils.GetEnv("SPOTIFY_ID")+":"+utils.GetEnv("SPOTIFY_SECRET")))))

	refreshResponse, _ := client.Do(refreshreq)
	
	spotifyRefreshResponse := make(map[string]interface{})
	refreshBody, _ := ioutil.ReadAll(refreshResponse.Body)
	refresherrorUnmarshal := json.Unmarshal(refreshBody, &spotifyRefreshResponse)
	if refresherrorUnmarshal != nil {
		fmt.Println(refresherrorUnmarshal)
	}
	accessToken := spotifyRefreshResponse["access_token"]

	models.SetUserToken(userID, "spotify_token", fmt.Sprintf("%s", accessToken))
}

/** @brief Returns the ID of the song based on the name given in params
 * @param userID uint , songName string
 * @return string
 */
func GetSongByName(userID uint, songName string) string {

	userToken := *models.FindUserToken(userID)

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", url.QueryEscape(songName))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	fmt.Println("TOKEN", userToken.SpotifyToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		RefreshSpotifyToken(userID)
		fmt.Fprintln(os.Stderr, "can't get song name")
		return ""
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+userToken.SpotifyToken)

	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	choosenSongID := gjson.GetBytes(body, "tracks.items.0.uri")

	return choosenSongID.String()
}

/** @brief Adds a given song to the user's queue
 * @param userID uint, params string
 */
func AddSongToQueue(userID uint, params string) {
	trackID := GetSongByName(userID, params)
	userToken := *models.FindUserToken(userID)

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/add-to-queue?uri=%s", trackID)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("can't get song name")
		return
	}
	req.Header.Add("Authorization", "Bearer "+userToken.SpotifyToken)
	client.Do(req)
}
