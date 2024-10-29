# Amfoss-Project
### Description
My Task was supposed to to scrape the top 10 Google search results for a given keyword, Utilizing Colly Library
### Features
The scraper would send a request to Google with the given keyword and capture the
search results page. It will then parse the HTML content and extract the top 10 search result's
### Issues Faced
I couldnt get the elements based on class and id's this was solved by adding
```
c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36"
```
