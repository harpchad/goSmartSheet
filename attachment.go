package gosmartsheet

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type RowAttachments struct {
	Data []struct {
		AttachmentType string `json:"attachmentType"`
		ID             int64  `json:"id"`
		MimeType       string `json:"mimeType"`
		Name           string `json:"name"`
		ParentID       int64  `json:"parentId"`
		ParentType     string `json:"parentType"`
	} `json:"data"`
}

type Attachment struct {
	AttachmentType     string `json:"attachmentType"`
	ID                 int64  `json:"id"`
	MimeType           string `json:"mimeType"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	URLExpiresInMillis int64  `json:"urlExpiresInMillis"`
}

//GetRowAttachments return a list or row attachments
func (c *Client) GetRowAttachments(sheetID, rowID string) (r *RowAttachments, err error) {
	path := "sheets/" + sheetID + "/rows/" + rowID + "/attachments"
	body, statusCode, err := c.Get(path)
	if err != nil {
		err = errors.Wrapf(err, "Failed to sheet and row (SID: %v, RID: %v)", sheetID, rowID)
		return
	}
	defer body.Close()
	dec := json.NewDecoder(body)
	if statusCode == 200 {
		r = &RowAttachments{}
		if err = dec.Decode(r); err != nil {
			err = errors.Wrap(err, "Failed to decode into RowAttachments")
		}
	} else {
		err = ErrorItemDecode(statusCode, dec)
	}
	return
}

func (c *Client) GetAttachment(sheetID, attachmentID string) (r *Attachment, err error) {
	path := "sheets/" + sheetID + "/attachments/" + attachmentID
	body, statusCode, err := c.Get(path)
	if err != nil {
		err = errors.Wrapf(err, "Failed to sheet and row (SID: %v, RID: %v)", sheetID, attachmentID)
		return
	}
	defer body.Close()
	dec := json.NewDecoder(body)
	if statusCode == 200 {
		r = &Attachment{}
		if err = dec.Decode(r); err != nil {
			err = errors.Wrap(err, "Failed to decode into Attachment")
		}
	} else {
		err = ErrorItemDecode(statusCode, dec)
	}
	return
}