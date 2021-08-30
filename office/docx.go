package office

import (
	"baliance.com/gooxml/document"
)

// ReadDocx function reads text from docx document.
func ReadDocx(p string) (string, error) {
	doc, err := document.Open(p)
	if err != nil {
		return "", err
	}
	t := ""
	paras := doc.Paragraphs()
	for _, para := range paras {
		runs := para.Runs()
		for _, run := range runs {
			t += run.Text()
		}
	}
	return t, nil
}
