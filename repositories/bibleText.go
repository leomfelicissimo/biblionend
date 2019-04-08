package repositories

import (
	"github.com/leomfelicissimo/biblionend/dbutil"
)

// BibleText represets a biblical text
type BibleText struct {
	Book    string
	Chapter int32
	Verse   int32
	Text    string
}

// BibleTextRepository represents a data repository of biblical texts
type BibleTextRepository struct{}

func documentToBibleText(document map[string]interface{}) *BibleText {
	return &BibleText{
		Book:    document["book"].(string),
		Chapter: document["chapter"].(int32),
		Verse:   document["verse"].(int32),
		Text:    document["text"].(string),
	}
}

// FindByReference method returns a array of BibleText using a biblical reference filter
// Ex: FindByReference("ap210"), returns the text of Book of Apocalipse, chapter 2, verse 10
func (r BibleTextRepository) FindByReference(ref string) (*BibleText, error) {
	repository := &dbutil.Repository{CollectionName: "nvi"}
	document, err := repository.FindBy("reference", ref)
	if err != nil {
		return nil, err
	}

	return documentToBibleText(document), nil
}
