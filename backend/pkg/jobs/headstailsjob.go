package jobs

import(
	"net/http"
	"errors"
	"io/ioutil"
	"fmt"

	"github.com/tidwall/gjson"
)

func GetHeadTailData() (string, error) {

	url := "https://coin-flip1.p.rapidapi.com/headstails"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2")
	req.Header.Add("X-RapidAPI-Host", "coin-flip1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("heads-tails api down")
		return "", myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	result := gjson.GetBytes(body, "outcome").String()

	fmt.Println(string(body))

	return result, nil 
}

func HeadsOrTails(params string) bool{
	result, _ := GetHeadTailData()
	if (result != params) {
		return true
	} else {
		return false
	}
}