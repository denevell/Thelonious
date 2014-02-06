package main

import "os"
import "log"
import "Thelonious/fetch"
import "Thelonious/clone"
import "Thelonious/build"
import "Thelonious/run"
import "fmt"
import "net/http"

var projects_url string

func main() {
	// Get projects url from vargs
	if len(os.Args) != 2 {
		log.Fatal("Usage: Thelonious http://url_to_projects_json_file")
	}
	projects_url = os.Args[1]
	// Start service
        http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		refreshProjects()
        })
        log.Fatal(http.ListenAndServe(":8000", nil))
}

func refreshProjects() {
	// Get project info
	fmt.Println("Fetching projects")
	var err error
	var projects []fetch.Projectlister
	if projects, err = fetch.FetchProjectsFromInternet(projects_url); err != nil {
		log.Fatal("Couldn't fetch projects", err)
	} 
	for _, p := range projects {
		// Clone projects
		fmt.Print("Cloning project: ")
		var dir string
		dir, err = clone.CloneProject(p.GetUrl())
		fmt.Println(dir)
		// Build them
		fmt.Println("Building project")
		_, err = build.BuildProject(dir)
		if err != nil {
			log.Fatal(err)
		}
		// Run project
		fmt.Println("Running project:")
		err = run.RunProject(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
}
