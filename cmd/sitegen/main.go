package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Link represents an external link card.
type Link struct {
	ID          string
	URL         string
	Title       string
	Description string
	SVGIcon     template.HTML
}

type BlogLink struct {
	Title       string
	URL         string
	Description string
}

type DetailItem struct {
	Label string
	Value string
}

type WorkItem struct {
	Title       string
	Platform    string
	URL         string
	Domain      string
	Summary     string
	Description string
}

type WorkSection struct {
	ID          string
	Title       string
	Description string
	Items       []WorkItem
}

// Profile represents the author's profile data.
type Profile struct {
	Name     string
	Subtitle string
	Bio      string
	Avatar   string
	Cover    string
	Details  []DetailItem
	Social   []Link
	Blogs    []BlogLink
}

type PageData struct {
	Profile Profile
	Works   []WorkSection
}

// raw structures for YAML unmarshalling.
type rawLink struct {
	ID          string `yaml:"id"`
	URL         string `yaml:"url"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	SVG         string `yaml:"svg"`
}

type rawDetailItem struct {
	Label string `yaml:"label"`
	Value string `yaml:"value"`
}

type rawBlogLink struct {
	Title       string `yaml:"title"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

type rawProfile struct {
	Name     string          `yaml:"name"`
	Subtitle string          `yaml:"subtitle"`
	Avatar   string          `yaml:"avatar"`
	Cover    string          `yaml:"cover"`
	Bio      string          `yaml:"bio"`
	Details  []rawDetailItem `yaml:"details"`
	Social   []rawLink       `yaml:"social"`
	Blogs    []rawBlogLink   `yaml:"blogs"`
}

type rawWorkItem struct {
	Title       string `yaml:"title"`
	Platform    string `yaml:"platform"`
	URL         string `yaml:"url"`
	Domain      string `yaml:"domain"`
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
}

type rawWorkSection struct {
	ID          string        `yaml:"id"`
	Title       string        `yaml:"title"`
	Description string        `yaml:"description"`
	Items       []rawWorkItem `yaml:"items"`
}

type rawWorks struct {
	Sections []rawWorkSection `yaml:"sections"`
}

func loadProfile(path string) (Profile, error) {
	var rp rawProfile
	b, err := os.ReadFile(path)
	if err != nil {
		return Profile{}, err
	}
	if err := yaml.Unmarshal(b, &rp); err != nil {
		return Profile{}, err
	}

	p := Profile{
		Name:     rp.Name,
		Subtitle: rp.Subtitle,
		Bio:      rp.Bio,
		Avatar:   rp.Avatar,
		Cover:    rp.Cover,
	}

	for _, detail := range rp.Details {
		p.Details = append(p.Details, DetailItem{
			Label: detail.Label,
			Value: detail.Value,
		})
	}

	for _, rl := range rp.Social {
		p.Social = append(p.Social, Link{
			ID:          rl.ID,
			URL:         rl.URL,
			Title:       rl.Title,
			Description: rl.Description,
			SVGIcon:     template.HTML(rl.SVG),
		})
	}

	for _, blog := range rp.Blogs {
		p.Blogs = append(p.Blogs, BlogLink{
			Title:       blog.Title,
			URL:         blog.URL,
			Description: blog.Description,
		})
	}

	return p, nil
}

func loadWorks(path string) ([]WorkSection, error) {
	var rw rawWorks
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &rw); err != nil {
		return nil, err
	}

	var works []WorkSection
	for _, section := range rw.Sections {
		ws := WorkSection{
			ID:          section.ID,
			Title:       section.Title,
			Description: section.Description,
		}
		for _, item := range section.Items {
			ws.Items = append(ws.Items, WorkItem{
				Title:       item.Title,
				Platform:    item.Platform,
				URL:         item.URL,
				Domain:      item.Domain,
				Summary:     item.Summary,
				Description: item.Description,
			})
		}
		works = append(works, ws)
	}
	return works, nil
}

func main() {
	profile, err := loadProfile("data/profile.yaml")
	if err != nil {
		log.Fatalf("Failed to load profile: %v", err)
	}

	works, err := loadWorks("data/works.yaml")
	if err != nil {
		log.Fatalf("Failed to load works: %v", err)
	}

	data := PageData{
		Profile: profile,
		Works:   works,
	}

	outputDir := "site"

	staticDirs := []string{"css", "js", "image"}
	for _, dir := range staticDirs {
		src := filepath.Join("src", dir)
		dst := filepath.Join(outputDir, dir)
		if err := copyDir(src, dst); err != nil {
			log.Fatalf("Failed to copy %s: %v", dir, err)
		}
	}

	tmpl, err := template.ParseFiles("src/index.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	outFile, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		log.Fatalf("Failed to create output HTML: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
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

func copyDir(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
			continue
		}
		if err := copyFile(srcPath, dstPath); err != nil {
			return err
		}
	}
	return nil
}
