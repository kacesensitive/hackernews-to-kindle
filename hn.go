package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Kids  []int  `json:"kids"`
}

type Comment struct {
	ID   int    `json:"id"`
	By   string `json:"by"`
	Text string `json:"text"`
}

func createHTMLContent(story Story, comments []Comment) (string, error) {
	// Start with the story title and URL.
	html := "<h1>" + story.Title + "</h1>\n<p><a href=\"" + story.URL + "\">Link to Story</a></p>\n"

	// Add each comment.
	for _, comment := range comments {
		html += "<h2>Comment by " + comment.By + "</h2>\n<p>" + comment.Text + "</p>\n"
	}

	return html, nil
}

func getTopStories() ([]int, error) {
	resp, err := http.Get(topStoriesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ids []int
	if err := json.Unmarshal(body, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func getStoryDetails(id int) (Story, error) {
	resp, err := http.Get(fmt.Sprintf("%s%d.json", storyURL, id))
	if err != nil {
		return Story{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Story{}, err
	}

	var story Story
	if err := json.Unmarshal(body, &story); err != nil {
		return Story{}, err
	}

	return story, nil
}

func getCommentDetails(id int) (Comment, error) {
	resp, err := http.Get(fmt.Sprintf("%s%d.json", storyURL, id))
	if err != nil {
		return Comment{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Comment{}, err
	}

	var comment Comment
	if err := json.Unmarshal(body, &comment); err != nil {
		return Comment{}, err
	}

	return comment, nil
}
