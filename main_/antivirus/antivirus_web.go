package antivirus

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func sanitizeFileName(fileName string) string {
	re := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	return re.ReplaceAllString(fileName, "_")
}

func FetchHTMLAndJS(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	htmlFilePath := filepath.Join("data/web", "site.html")
	err = os.WriteFile(htmlFilePath, bodyBytes, 0644)
	if err != nil {
		return err
	}

	doc, err := html.Parse(bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	var jsFiles []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					jsFiles = append(jsFiles, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	for _, jsFile := range jsFiles {
		if strings.HasPrefix(jsFile, "/") {
			jsFile = url + jsFile
		}

		err := DownloadJS(jsFile)
		if err != nil {
			continue
		}
		checkJSFile(sanitizeFileName(filepath.Base(jsFile)))
	}

	return nil
}

func DownloadJS(jsURL string) error {
	resp, err := http.Get(jsURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	jsFilePath := filepath.Join("data/web", sanitizeFileName(filepath.Base(jsURL)))
	jsFile, err := os.Create(jsFilePath)
	if err != nil {
		return err
	}
	defer jsFile.Close()

	_, err = io.Copy(jsFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func checkJSFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	patterns := []string{
		"window.location.href",
		"window.open",
		"requestFullscreen",
	}

	foundPatterns := make(map[string]bool)
	contentStr := string(content)
	for _, pattern := range patterns {
		if strings.Contains(contentStr, pattern) {
			foundPatterns[pattern] = true
		}
	}

	writeResultsToFile(foundPatterns)
}

func writeResultsToFile(foundPatterns map[string]bool) {
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return
	}

	err = os.WriteFile("data/inter.txt", []byte(""), 0644)
	if err != nil {
		return
	}

	file, err := os.OpenFile("data/inter.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	for pattern := range foundPatterns {
		_, err := file.WriteString(pattern + "\n")
		if err != nil {
			return
		}
	}
}

func DeleteFiles() {
	files := []string{filepath.Join("data/web", "site.html")}
	jsFiles, err := filepath.Glob(filepath.Join("data/web", "*.js"))
	if err == nil {
		files = append(files, jsFiles...)
	}

	for _, file := range files {
		os.Remove(file)
	}
}

func CheckUrlInFile(url string) int {
	file, err := os.Open("data/site_virus.txt")
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == url {
			return 1
		}
	}

	if err := scanner.Err(); err != nil {
		return 0
	}

	return 0
}
