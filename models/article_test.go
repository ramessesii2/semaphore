package models

import "testing"

func TestGetAllArtciles(t *testing.T) {
	allArticles := GetAllArtciles()
	if len(allArticles) != len(Articles) {
		t.Fail()
		return
	}
	for i, v := range allArticles {
		if v.ID != Articles[i].ID || v.Title != Articles[i].Title || v.Content != Articles[i].Content {
			t.Fail()
			return
		}
	}
}
