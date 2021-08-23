package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
	"golang.org/x/xerrors"

	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/entity"
	"github.com/gotd/td/telegram/message/styling"
)

type issueType string

const (
	featureRequest issueType = "feature request"
	bugReport      issueType = "bug report"
	plain          issueType = "issue"
)

func getIssueType(issue *github.Issue) issueType {
	for _, label := range issue.Labels {
		switch label.GetName() {
		case "enhancement":
			return featureRequest
		case "bug":
			return bugReport
		}
	}

	return plain
}

func formatIssue(e *github.IssuesEvent) message.StyledTextOption {
	issue := e.GetIssue()
	sender := e.GetSender()
	formatter := func(eb *entity.Builder) error {
		eb.Plain("New ")
		eb.Plain(string(getIssueType(issue)))

		urlName := fmt.Sprintf(" %s#%d",
			e.GetRepo().GetFullName(),
			issue.GetNumber(),
		)
		eb.TextURL(urlName, issue.GetHTMLURL())
		eb.Plain(" by ")
		eb.TextURL(sender.GetLogin(), sender.GetHTMLURL())
		eb.Plain("\n\n")

		eb.Italic(issue.GetTitle())
		eb.Plain("\n\n")

		length := len(issue.Labels)
		if length > 0 {
			eb.Italic("Labels: ")

			for idx, label := range issue.Labels {
				switch label.GetName() {
				case "":
					continue
				case "bug":
					eb.Bold(label.GetName())
				default:
					eb.Italic(label.GetName())
				}

				if idx != length-1 {
					eb.Plain(", ")
				}
			}
		}
		return nil
	}

	return styling.Custom(formatter)
}

func (h Webhook) handleIssue(ctx context.Context, e *github.IssuesEvent) error {
	if e.GetAction() != "opened" {
		return nil
	}

	p, err := h.notifyPeer(ctx)
	if err != nil {
		return xerrors.Errorf("peer: %w", err)
	}

	if _, err := h.sender.To(p).NoWebpage().StyledText(ctx, formatIssue(e)); err != nil {
		return xerrors.Errorf("send: %w", err)
	}
	return nil
}
