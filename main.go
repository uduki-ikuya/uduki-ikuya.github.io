package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Link represents an external link card
type Link struct {
	ID          string
	URL         string
	Title       string
	Description string
	SVGIcon     template.HTML
}

// Profile represents the author's profile data
type Profile struct {
	Name     string
	Subtitle string
	Bio      string
	Links    []Link
}

func main() {
	// 1. Define data
	profile := Profile{
		Name:     "卯月 郁哉",
		Subtitle: "Novelist Portfolio",
		Bio:      "言葉を通じて、誰かの心に小さな灯火をともすような物語を紡いでいます。ファンタジー、SF、現代ドラマなど様々なジャンルの小説をネット上で公開・執筆しています。",
		Links: []Link{
			{
				ID:          "link-kakuyomu",
				URL:         "https://kakuyomu.jp/",
				Title:       "カクヨム",
				Description: "kakuyomu.jp - 連載作品・短編の掲載ページ",
				SVGIcon:     template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path></svg>`),
			},
			{
				ID:          "link-narou",
				URL:         "https://syosetu.com/",
				Title:       "小説家になろう",
				Description: "syosetu.com - 投稿作品の一覧ページ",
				SVGIcon:     template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5z"></path><line x1="16" y1="8" x2="2" y2="22"></line><line x1="17.5" y1="15" x2="9" y2="15"></line></svg>`),
			},
			{
				ID:          "link-twitter",
				URL:         "https://x.com/",
				Title:       "X (旧 Twitter)",
				Description: "@uduki_ikuya - 近況報告や執筆中のつぶやき",
				SVGIcon:     template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"></path></svg>`),
			},
			{
				ID:          "link-website",
				URL:         "https://uduki-ikuya.github.io",
				Title:       "作品ポートフォリオ (準備中)",
				Description: "本サイトにて独自の作品管理ページを開発予定です",
				SVGIcon:     template.HTML(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="2" y1="12" x2="22" y2="12"></line><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path></svg>`),
			},
		},
	}

	// 2. Ensure output directory exists
	outputDir := "site"
	if err := os.MkdirAll(filepath.Join(outputDir, "css"), 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// 3. Copy CSS file
	if err := copyFile("src/css/style.css", filepath.Join(outputDir, "css/style.css")); err != nil {
		log.Fatalf("Failed to copy CSS: %v", err)
	}

	// 4. Render HTML template
	tmpl, err := template.ParseFiles("src/index.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	outFile, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		log.Fatalf("Failed to create output HTML: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, profile); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Println("Build successful!")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return nil
}
