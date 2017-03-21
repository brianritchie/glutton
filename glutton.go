package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

// This is an error function that triggers if an Error occurs
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//This is an introductory message
	fmt.Println("Welcome to Glutton")

	//Declaring source file of Subreddits to check
	dat, err := ioutil.ReadFile("subreddit.txt")

	//This line checks if the file declared above exists. If not - panic!
	check(err)

	//The next two lines are a token output of all the Subreddits declared in the file
	fmt.Print("\nSubreddits to parse are...\n")
	fmt.Print("\nhttps://reddit.com/r/"+ string(dat))

	//Creation of Client as suggested to handle requests to Reddit
	client := &http.Client{}

	//Http Request is fed the URL of the Subreddit - Currently only supporting a single line
	request, err := http.NewRequest("GET", "http://reddit.com/r/"+string(dat), nil)

	//Lines to check if there are any errors that occured on the URL supplied above - Still single line
	if err != nil {
		log.Fatalln(err)
	}

	//Manually setting User-Agent because Reddit doesn't like bots
	request.Header.Set("User-Agent", "[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36]")

	//Make Request with the Header set
	resp, err := client.Do(request)

	//Capture body in Variable bytes to be read
	bytes, _ := ioutil.ReadAll(resp.Body)

	//Output to screen the captured output
	fmt.Println("HTML:\n\n", string(bytes))

	//Close the file since we are done looking at the output
	resp.Body.Close()



}

/*

Current Issues

- Single line reading for Subreddits
- Reddit returns the sidebar and not the list of posts on page

Current To-Dos in Mind

- Write tests for TDD
- Http Proxies for query if required
- Randomized User Agents
- Read more than a single line of input from file and parse each as URL
- Decide if we want to rely on graw as a wrapper or write own

 */
