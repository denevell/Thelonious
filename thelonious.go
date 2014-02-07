package main

import "os"
import "log"
import "Thelonious/fetch"
import "Thelonious/clone"
import "Thelonious/build"
import "Thelonious/run"
import "net/http"
import "Thelonious/utils"

var projects_url string

func main() {
	// Get projects url from vargs
	if len(os.Args) != 2 {
		log.Fatal("Usage: Thelonious http://url_to_projects_json_file")
	}
	projects_url = os.Args[1]
	// Start service
        http.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
			refreshProjects(w)
        })
        log.Fatal(http.ListenAndServe(":8000", nil))
}

func refreshProjects(w http.ResponseWriter) {
	// Get project info
	utils.PrintlnAndFlush(w, "== START")
	var err error
	var projects []fetch.Projectlister
	if projects, err = fetch.FetchProjectsFromInternet(projects_url); err != nil {
		log.Fatal("Couldn't fetch projects", err)
	} 
	for _, p := range projects {
		// Clone projects
		utils.PrintAndFlush(w, "Cloning project: ")
		var dir string
		dir, err = clone.CloneProject(p.GetUrl())
		utils.PrintlnAndFlush(w, dir)
		// Build them
		utils.PrintlnAndFlush(w, "Building project")
		_, err = build.BuildProject(dir)
		if err != nil {
			log.Fatal(err)
		}
		// Run project
		utils.PrintlnAndFlush(w, "Running project")
		err = run.RunProject(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
	utils.PrintlnAndFlush(w, "== FINISH")
}
