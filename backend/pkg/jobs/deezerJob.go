/** @file deezerJob.go
 * @brief Deezer reaction : adding song to playlist with both their respective titles (by the use of searchs)
 * @author Erwan
 */

// @cond

package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/tidwall/gjson"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief gets a track, album, artist or label id through a search on Deezer's API
 *
 * @param[string] query
 * @param[string] queryType
 * @param[string] accessToken
 *
 * @return queryId
 */
func GetQueryId(query string, queryType string, accessToken string) int64 {
	Url := "https://api.deezer.com/search?q=" + queryType + ":\"" + url.QueryEscape(query) + "\"&access_token=" + accessToken
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "client: could not read response body: ", err)
	}
	parsedResponse := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &parsedResponse)
	if errorUnmarshal != nil {
		fmt.Fprintln(os.Stderr, errorUnmarshal)
	}
	id := gjson.GetBytes(body, "data.0.id")
	idConverted := id.Int()
	return (idConverted)
}

/** @brief gets a playlist ID based on its EXACT name
 *
 * @param[string] accessToken
 * @param[string] playlistName
 *
 * @return playlistId
 */
func GetPlaylistIdByName(accessToken string, playlistName string) int64 {
	Url := "https://api.deezer.com/user/me/playlists?limit=999&access_token=" + accessToken
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "client: could not read response body:", err)
		return -1
	}
	title := gjson.GetBytes(body, "data.0.title")
	for i := 0; i < 999; i++ {
		currentIndex := strconv.Itoa(i)
		title = gjson.GetBytes(body, "data."+currentIndex+".title")
		if title.Str == playlistName {
			playlistId := gjson.GetBytes(body, "data."+currentIndex+".id")
			idConverted := playlistId.Int()
			return (idConverted)
		}
	}
	return (-1)
}

/** @brief Adds a song to a playlist
 *
 * The song and the playlist title are passed in the `params` parameter
 *
 * @param[uint] userId
 * @param[string] params
 */
func AddSongToPlaylist(userID uint, params string) {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 2 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return
	}
	userToken := *models.FindUserToken(userID)
	trackName := paramsArr[0]
	playlistName := paramsArr[1]

	playlistId := GetPlaylistIdByName(userToken.DeezerToken, playlistName)
	playlistIdConverted := strconv.FormatInt(playlistId, 10)
	trackId := GetQueryId(trackName, "track", userToken.DeezerToken)
	trackIdConverted := strconv.FormatInt(trackId, 10)

	Url := "https://api.deezer.com/playlist/" + playlistIdConverted + "/tracks?songs=" + trackIdConverted + "&access_token=" + userToken.DeezerToken
	req, _ := http.NewRequest("POST", Url, nil)
	hc := &http.Client{}
	hc.Do(req)
}
