package fetch

import "io/ioutil"
import "encoding/json"
import "errors"
import "net/http"
import "fmt"

type project struct {
	Url string
	Binary string
}

func FetchProjectsFromInternet(url string) ([]Projectlister, error) {
	m := make([]project, 1)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	file, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(file))

	if err == nil {
		if err = json.Unmarshal(file, &m); err == nil {
			ints := make([]Projectlister, len(m))
			for i, v := range m {
				ints[i] = v
			}
			return ints, nil
		} else { 
			return nil, err
		}
	} else {
		return nil, err
	}
	return nil, errors.New("Couldn't fetch projects")
}

func (p project) GetUrl() string {
	return p.Url
}

func (p project) GetBinary() string {
	return p.Binary
}
