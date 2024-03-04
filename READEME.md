# Home24 Web Analyzer

Home24 Web Analyzer is a web application designed to analyze web pages, providing users with detailed insights into the structure and content of the given URL.

## Features

- Users can input the URL of the webpage to be analyzed.
- Upon submission, the application retrieves and analyzes the webpage content.
- The application presents the following information to the user:
  - HTML version of the document.
  - Page title.
  - Number of headings of each level in the document.
  - Number of internal and external links in the document.
  - Number of inaccessible links, if any.
  - Presence of a login form on the page.
- In case the URL provided by the user is unreachable, an appropriate error message is displayed, containing the HTTP status code and a descriptive error message.

## Usage

To run the Home24 Web Analyzer application using Docker Compose, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Make sure you have Docker and Docker Compose installed on your machine.
4. Run the following command to start the application:

```bash
docker-compose up
```

# Building From Source
Pre-requisites:

* Git
* Go 1.20 < 
* Gin Framework

Steps to run the application without Docker:
```bash

~$ git clone https://github.com/krishan1988/home24-web-analyser

~$ cd home24-web-analyser/cmd/web

~$ go run web.go
```


Once the application starts running, open the browser and go to `http://localhost:8080` to access the web view.

