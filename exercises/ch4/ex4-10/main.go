// Issues Prints a table of GitHub issues mathcing the search terms.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/exercises/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)

	cs := newCategories(result.Items)
	fmt.Printf("%s\n%s\n%s\n", cs.ltMonth, cs.ltYear, cs.mtYear)
}

type categories struct {
	ltMonth, ltYear, mtYear *issues
}

func newCategories(gs []*github.Issue) *categories {
	now := time.Now()
	mtyFn := func(t time.Time) bool {
		return t.Before(now.Add(time.Hour * 24 * -365))
	}
	ltyFn := func(t time.Time) bool {
		return t.Before(now.Add(time.Hour * 24 * -31))
	}
	ltmFn := func(t time.Time) bool {
		return t.After(now.Add(time.Hour * 24 * -31))
	}

	return &categories{
		ltMonth: newIssues(ltmFn, "< 1 month old", gs),
		ltYear:  newIssues(ltyFn, "< 1 year old", gs),
		mtYear:  newIssues(mtyFn, "> 1 year old", gs),
	}
}

type issues struct {
	data []*github.Issue
	age  string
}

func newIssues(fn func(time.Time) bool, age string, gs []*github.Issue) *issues {
	data := make([]*github.Issue, 0)

	for _, g := range gs {
		if fn(g.CreatedAt) {
			data = append(data, g)
		}
	}
	return &issues{data: data, age: age}
}

func (iss *issues) String() string {
	s := fmt.Sprintf("%d issues %s\n", len(iss.data), iss.age)
	for _, g := range iss.data {
		s += fmt.Sprintf("--#%-5d %9.9s %.55s\n", g.Number, g.User.Login, g.Title)
	}
	return s
}

// func categorizeByTime(issues *github.IssuesSearchResult) timeCategory {
// 	formattedIssues := timeCategory{}
// 	now := time.Now()
// 	for _, item := range issues.Items {
// 		switch {
// 		case item.CreatedAt.Before(now.Add(time.Hour * 24 * -365)):
// 			formattedIssues.moreThanYear = append(formattedIssues.moreThanYear, item)
// 		case item.CreatedAt.Before(now.Add(time.Hour * 24 * -31)):
// 			formattedIssues.lessThanYear = append(formattedIssues.lessThanYear, item)
// 		default:
// 			formattedIssues.lessThanMonth = append(formattedIssues.lessThanMonth, item)
// 		}
// 	}
// 	return formattedIssues
// }

//
