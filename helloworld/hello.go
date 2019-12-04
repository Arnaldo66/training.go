package main

import "fmt"

type Post struct {
	Title string
	Text string
	published bool
}

func (p Post)  Headline() string{
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:20])
}

func (p *Post) Published() bool  {
	return p.published
}

func (p *Post) Publish() {
	p.published = true
}

func UpdateTitle(p *Post)  {
	p.Title = "Change"
}

func main() {
	p := &Post{
		Title:     "Go go gadjet",
		Text:      `il est une fois dans 
la ville de foix
un vendeur de foie
`,
		published: false,
	}

	fmt.Println(p.Headline())
	fmt.Println(p.published)
	p.Publish()
	fmt.Println(p.published)
	UpdateTitle(p)
	fmt.Println(p.Headline())
}





