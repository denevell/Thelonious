
package fetch

import ioutil "io/ioutil"
import json "encoding/json"
import errors "errors"

func FetchProjects() ([]Projectlister, error) {
	m := make([]project, 1)
	file, err := ioutil.ReadFile("projects.json")
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
