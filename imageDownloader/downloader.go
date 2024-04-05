package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// Function to download and save an image from a URL
func downloadImage(url string, imageNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading image %d...\n", imageNumber)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download image %d: %s\n", imageNumber, err)
		return
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create(fmt.Sprintf("image_%d.png", imageNumber))
	if err != nil {
		fmt.Printf("Failed to create file for image %d: %s\n", imageNumber, err)
		return
	}
	defer file.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Failed to save image %d: %s\n", imageNumber, err)
		return
	}

	fmt.Printf("Image %d downloaded.\n", imageNumber)
}

func main() {
	var wg sync.WaitGroup
	imageUrls := []string{
		"https://imgs.xkcd.com/comics/family_circus.jpg",
		"https://imgs.xkcd.com/comics/bill_nye.png",
		// Add more image URLs here
	}

	for i, url := range imageUrls {
		wg.Add(1)
		go downloadImage(url, i+1, &wg)
	}

	wg.Wait()
	fmt.Println("All images downloaded.")
}
