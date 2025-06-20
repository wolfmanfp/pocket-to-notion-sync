package notion

import (
	"context"
	"fmt"
	"time"

	"github.com/jomei/notionapi"
)

type Client struct {
	client     *notionapi.Client
	databaseID string
}

func NewClient(token, databaseID string) (*Client, error) {
	client := notionapi.NewClient(notionapi.Token(token))
	return &Client{client: client, databaseID: databaseID}, nil
}

func (c *Client) ArticleExists(url string) (bool, error) {
	query := &notionapi.DatabaseQueryRequest{
		Filter: &notionapi.PropertyFilter{
			Property: "URL",
			RichText: &notionapi.TextFilterCondition{
				Equals: url,
			},
		},
	}

	resp, err := c.client.Database.Query(context.Background(), notionapi.DatabaseID(c.databaseID), query)
	if err != nil {
		return false, fmt.Errorf("failed to query Notion database: %v", err)
	}

	return len(resp.Results) > 0, nil
}

func (c *Client) CreatePage(title, url, excerpt string, tags []string, cover string, added time.Time) error {
	addedDate := notionapi.Date(added)
	pageCreateRequest := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: notionapi.DatabaseID(c.databaseID),
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: title,
						},
					},
				},
			},
			"URL": notionapi.URLProperty{
				URL: url,
			},
			"Added": notionapi.DateProperty{
				Type: notionapi.PropertyTypeDate,
				Date: &notionapi.DateObject{
					Start: &addedDate,
					End:   nil,
				},
			},
		},
	}

	if cover != "" {
		pageCreateRequest.Children = append(pageCreateRequest.Children, notionapi.ImageBlock{
			BasicBlock: notionapi.BasicBlock{
				Object: "block",
				Type:   "image",
			},
			Image: notionapi.Image{
				External: &notionapi.FileObject{
					URL: cover,
				},
			},
		})

		pageCreateRequest.Cover = &notionapi.Image{
			External: &notionapi.FileObject{
				URL: cover,
			},
		}
	}

	if excerpt != "" {
		pageCreateRequest.Children = append(pageCreateRequest.Children, notionapi.ParagraphBlock{
			BasicBlock: notionapi.BasicBlock{
				Object: "block",
				Type:   "paragraph",
			},
			Paragraph: notionapi.Paragraph{
				RichText: []notionapi.RichText{
					{
						Type: notionapi.ObjectTypeText,
						Text: &notionapi.Text{
							Content: excerpt,
						},
					},
				},
			},
		})
	}

	if len(tags) > 0 {
		notionTags := make([]notionapi.Option, len(tags))
		for i, tag := range tags {
			notionTags[i] = notionapi.Option{Name: tag}
		}
		pageCreateRequest.Properties["Tags"] = notionapi.MultiSelectProperty{
			MultiSelect: notionTags,
		}
	}

	_, err := c.client.Page.Create(context.Background(), pageCreateRequest)
	if err != nil {
		return fmt.Errorf("failed to create Notion page: %v", err)
	}

	return nil
}
