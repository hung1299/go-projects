package user

import (
	"encoding/json"
	"fmt"
	"hung1299/go-projects/github-user-activity/internal/http"
	"os"
	"strconv"
)

type UserEvent struct {
	Type string `json:"type"`
	Repo struct{ 
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Ref string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			SHA string `json:"sha"`
		} `json:"commits"`
	} `json:"payload"`
}

type UserEvents []UserEvent

func GetEvents(name string) {
	n := fmt.Sprintf(`https://api.github.com/users/%s/events`, name)
	bs := http.Get(n)
	var ues UserEvents
	
	if err := json.Unmarshal(bs, &ues); err != nil {
		fmt.Printf("May be user %s was not found or request got rate limit", name)
		os.Exit(1)
	}
	
	for _, ue := range ues {
		printByEvent(ue)
	}
}

func printByEvent(ue UserEvent) {
	switch ue.Type {
	case "PushEvent":
		totalCommits := len(ue.Payload.Commits)
		commitText :=  strconv.FormatInt(int64(totalCommits), 10) + " commit"

		if totalCommits > 1 {
			commitText += "s"
		}

		fmt.Printf("- Pushed %s to %s\n", commitText, ue.Repo.Name)
	case "CreateEvent":
		text := fmt.Sprintf(`- Created a new %s named `, ue.Payload.RefType)
		if ue.Payload.RefType == "branch" {
			text += fmt.Sprintf(`%s in %s`, ue.Payload.Ref, ue.Repo.Name)
		} else if ue.Payload.RefType == "repository" {
			text += fmt.Sprintf(`%s`, ue.Repo.Name)
		} else {
			text += " i don't know"
		}

		fmt.Println(text)
	case "WatchEvent":
		fmt.Printf("- Watched %s\n", ue.Repo.Name)
	case "ForkEvent":
		fmt.Printf("- Forked %s\n", ue.Repo.Name)
	case "PullRequestEvent":
		fmt.Printf("- Opened a Pull Request at %s\n", ue.Repo.Name)
	case "PullRequestReviewEvent":
		fmt.Printf("- Opened a Pull Request Review at %s\n", ue.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("- Opened a Issues at %s\n", ue.Repo.Name)
	case "IssueCommentEvent":
		fmt.Printf("- Commented in an issue at %s\n", ue.Repo.Name)
	default:
		fmt.Println("- I did not think about this type before", ue.Type)
	}
}