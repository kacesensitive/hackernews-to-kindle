# HackerNews-to-Kindle

This is a casual project that fetches the top stories from Hacker News and sends them to your Kindle. It was built to fit my needs, so feel free to modify it to suit yours!

## Setup

1. Clone the repository.

```bash
git clone https://github.com/kacesensitive/hackernews-to-kindle.git
```

2. Navigate to the cloned repository.

```
cd hackernews-to-kindle
```

3. Update config.go with your details. You'll find the file in the main package.

```
package main

const (
	hackerNewsAPI = "https://hacker-news.firebaseio.com/v0/"
	topStoriesURL = hackerNewsAPI + "topstories.json"
	storyURL      = hackerNewsAPI + "item/"

	mobiConvertCommand = "/path/to/ebook-convert"
	mobiFileName       = "top-stories.mobi"

	kindleEmail   = "your-kindle-email@kindle.com"
	sendFromEmail = "your-email@example.com"
)
```

4. Make sure to get and set the SendGrid API key.

```
export SENDGRID_API_KEY='your_sendgrid_api_key'
```

5. Once you have configured the program, you can compile and run it with:

```
go build
./hackernews-to-kindle
```

Remember, this is a fun project, and it's tailored to my needs. Feel free to tinker around and adjust things as you see fit. Enjoy your reading!

## Note
This project uses SendGrid's Go Library for email dispatch and Calibre's ebook-convert for conversion from HTML to MOBI. You need to have these installed and properly configured for the program to work correctly.

Please replace 'yourusername', 'your-kindle-email@kindle.com', 'your-email@example.com', and 'your_sendgrid_api_key' with your actual GitHub username, Kindle email, your email, and your SendGrid API key respectively.
