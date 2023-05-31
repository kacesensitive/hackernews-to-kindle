package main

import (
	"fmt"
	"log"
)

func main() {
	// Fetch top stories
	topStoryIDs, err := getTopStories()
	if err != nil {
		log.Println("Error fetching top stories:", err)
		return
	}

	// Limit to top 10 stories
	if len(topStoryIDs) > 10 {
		topStoryIDs = topStoryIDs[:10]
	}

	var stories []Story
	for _, id := range topStoryIDs {
		story, err := getStoryDetails(id)
		if err != nil {
			log.Println("Error fetching story details:", err)
			continue
		}
		stories = append(stories, story)

		// Fetch top 10 comments for each story
		if len(story.Kids) > 10 {
			story.Kids = story.Kids[:10]
		}

		var comments []Comment
		for _, commentID := range story.Kids {
			comment, err := getCommentDetails(commentID)
			if err != nil {
				log.Println("Error fetching comment details:", err)
				continue
			}
			comments = append(comments, comment)
		}

		// Create HTML content
		htmlContent, err := createHTMLContent(story, comments)
		if err != nil {
			log.Println("Error creating HTML content:", err)
			continue
		}

		// Convert HTML content to MOBI

		htmlFileName := "story.html"
		err = writeToFile(htmlFileName, htmlContent)
		if err != nil {
			log.Println("Error writing HTML content to file:", err)
			continue
		}

		err = convertHTMLToMobi(htmlFileName)
		if err != nil {
			log.Println("Error converting HTML to MOBI:", err)
			continue
		}

		// Send MOBI to Kindle
		err = sendToKindle(mobiFileName)
		if err != nil {
			log.Println("Error sending MOBI to Kindle:", err)
			continue
		}

		fmt.Println("Story successfully sent to Kindle:", story.Title)
	}

	fmt.Println("Done!")
}
