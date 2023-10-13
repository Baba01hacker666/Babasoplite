package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"os/exec"
	"strings"
	"log"
	"bufio"
)

func main() {
	fmt.Println("Welcome to doradorid")
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <ip>")
		return
	}

	ip := os.Args[1]
	resp, err := http.Get("http://" + ip)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		fmt.Println("[+] HTTP 200 - No account authorization needed")
		runNiktoCommand(ip)
		checkDefaultPasswords(ip)
		bruteForceDirectories(ip)
	case 401:
		fmt.Println("[+] HTTP 401 - Account authorization needed")
		checkDefaultPasswords(ip)
	case 404:
		fmt.Println("Sorry, web page not found")
	case 403:
		fmt.Println("Web page not allowed to be opened")
	case 502:
		fmt.Println("Bad gateway")
	default:
		fmt.Println("Unknown status code:", resp.StatusCode)
	}
}

func runNiktoCommand(ip string) {
	cmd := exec.Command("nikto", "-url", "http://"+ip)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running nikto:", err)
	}
}

func checkDefaultPasswords(ip string) {
	client := &http.Client{}
	file, err := os.Open("credentials.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		creds := strings.Split(line, ":")
		if len(creds) != 2 {
			continue // Skip lines with incorrect format
		}

		username, password := creds[0], creds[1]
		req, err := http.NewRequest("GET", "http://" + ip,nil)
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Failed with username:", username)
			continue // Try the next combination
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Correct username: %s, password: %s\n", username, password)
			    bodytext,err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			if bodytext == nil {
			  println("empty text ")
			}
			break // Exit the loop when correct credentials are found
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func bruteForceDirectories(ip string) {
	// Implement your logic to brute-force directories here.
	// You can use a tool like DirBuster or custom code for this purpose.
}