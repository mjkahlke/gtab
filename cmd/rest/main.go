// Copyright 2022 Michael Kahlke.

package main

// Given a known chord name, display the guitar tab.

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	gtab "gtab/impl"
)

type formula struct {
	Suffix	string `json:"suffix"`
	Halfsteps	string `json:"halfsteps"`
}

type tablature struct {
	Chord	string `json:"chord"`
	Tabs	[]string `json:"tabs"`
}

func main() {
	router := gin.Default()
	router.GET("/list", getList)
	router.GET("/tabs/:chord", getTabsByChord)
	router.Run("localhost:7777")
}

func getList(c *gin.Context) {
	var formulas []formula
	for key, halfsteps := range gtab.Chords {
		formulas = append(formulas, formula{key, gtab.GenTab(halfsteps)})
	}
	c.IndentedJSON(http.StatusOK, formulas)
}

func getTabsByChord(c *gin.Context) {
	chord := c.Param("chord")
	chord = strings.Replace(chord, "H", "#", 1)
	root, suffix := gtab.ChordComponents(chord)
	if root == "" {
		message := chord + ": invalid chord name specified"
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
		return
	}
	if suffix == gtab.BADSUFFIX {
		message := chord + ": unsupported chord suffix specified"
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": message})
		return
	}
	tabs := gtab.Match(root, suffix)
	var result tablature
	result.Chord = chord
	result.Tabs = []string{}
	for _, tab := range tabs {
		result.Tabs = append(result.Tabs, tab)
	}
	c.IndentedJSON(http.StatusOK, result)
}

