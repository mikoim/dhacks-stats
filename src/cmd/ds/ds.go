package main

import (
	"github.com/google/go-github/github"
	"fmt"
	"log"
)

type DHacksStats struct {
	GitHub *github.Client
}

func (ds *DHacksStats) Run() error {
	repos, _, err := ds.GitHub.Repositories.ListByOrg("dhacks", nil)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Printf("%s\t%s\t%d\n", stripString(repo.FullName), stripString(repo.Language), stripInt(repo.Size))
	}

	return nil
}

func NewDHacksStats() *DHacksStats {
	ds := &DHacksStats{};

	ds.GitHub = github.NewClient(nil)

	return ds
}

// ToDo: dirty hack
func stripString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ToDo: dirty hack
func stripInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func main() {
	err := NewDHacksStats().Run()
	if err != nil {
		log.Fatalln(err)
	}
}