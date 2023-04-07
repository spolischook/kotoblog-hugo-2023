package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conn := sqlite.Open("../../data.db")
	db, err := gorm.Open(conn, &gorm.Config{})
	checkErr(err)
	converter := md.NewConverter("", true, nil)

	articles := make([]Article, 0)
	result := db.Preload("Tags").Find(&articles)

	_ = result
	fmt.Println(len(articles))

	for _, article := range articles {
		markdown := article.TextSource
		if article.TextSource == article.Text {
			markdown, err = converter.ConvertString(article.TextSource)
			checkErr(err)
		}

		tags := make([]string, 0)
		for _, tag := range article.Tags {
			tags = append(tags, fmt.Sprintf(`"%s"`, tag.Title))
		}
		title := strings.Replace(article.Title, `"`, `\"`, -1)
		text := fmt.Sprintf(`---
author: "Serhii Polishchuk"
title: "%s"
date: %s
tags: [%s]
draft: %v
---
<!--more-->
%s
`, title, article.CreatedAt.Format("2006-01-02"), strings.Join(tags, ", "), !article.Published, markdown)

		path := "/Users/spolischook/www/hugo/www/content/articles/" + article.Slug
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			checkErr(err)
		}

		err = os.WriteFile(path+"/index.md", []byte(text), 0644)
		checkErr(err)
	}

	fmt.Println("The end!")
}

type Article struct {
	ID         uint
	Title      string
	Slug       string
	Text       string
	TextSource string
	Image      string
	Published  bool
	CreatedAt  time.Time
	Tags       []Tag `gorm:"many2many:article_tag;"`
}
type Tag struct {
	ID    uint
	Title string
	Slug  string
}

func (Tag) TableName() string {
	return "tag"
}

func (Article) TableName() string {
	return "article"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
