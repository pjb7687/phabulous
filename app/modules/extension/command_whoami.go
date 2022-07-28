package extension

import (
	"fmt"

	gonduitRequests "github.com/etcinit/gonduit/requests"
	"github.com/pjb7687/phabulous/app/gonduit/extensions"
	"github.com/pjb7687/phabulous/app/gonduit/extensions/requests"
	"github.com/pjb7687/phabulous/app/interfaces"
	"github.com/pjb7687/phabulous/app/messages"
)

// WhoamiCommand allows one to send test messages to the feed channel.
type WhoamiCommand struct{}

// GetUsage returns the usage of this command.
func (t *WhoamiCommand) GetUsage() string {
	return "whoami"
}

// GetDescription returns the description of this command.
func (t *WhoamiCommand) GetDescription() string {
	return "Gets the name of your Phabricator user."
}

// GetMatchers returns the matchers for this command.
func (t *WhoamiCommand) GetMatchers() []string {
	return []string{}
}

// GetIMMatchers returns IM matchers for this command.
func (t *WhoamiCommand) GetIMMatchers() []string {
	return []string{"whoami"}
}

// GetMentionMatchers returns the channel mention matchers for this command.
func (t *WhoamiCommand) GetMentionMatchers() []string {
	return []string{"whoami"}
}

// GetHandler returns the handler for this command.
func (t *WhoamiCommand) GetHandler() interfaces.Handler {
	return func(s interfaces.Bot, m interfaces.Message, matches []string) {
		if _, ok := s.(interfaces.SlackBot); !ok {
			s.Post(
				m.GetChannel(),
				"This command only works over Slack",
				messages.IconTasks,
				true,
			)

			return
		}

		conn, err := s.GetGonduit()
		if err != nil {
			s.Excuse(m, err)
			return
		}

		res, err := extensions.PhabulousFromSlack(
			conn,
			requests.PhabulousFromSlackRequest{
				AccountIDs: []string{m.GetUserID()},
			},
		)
		if err != nil {
			s.Excuse(m, err)
			return
		}

		if len(*res) == 0 {
			s.Post(
				m.GetChannel(),
				"I was unable to find a Phabricator user linked with your "+
					"Slack account. Make sure they are linked under "+
					"_External Accounts_ in your Phabricator user settings.",
				messages.IconTasks,
				true,
			)
			return
		}

		userPHIDs := []string{}
		for _, userInfo := range *res {
			userPHIDs = append(userPHIDs, userInfo.UserPHID)
		}

		res2, err := conn.UserQuery(gonduitRequests.UserQueryRequest{
			PHIDs: userPHIDs,
		})

		if err != nil {
			s.Excuse(m, err)
			return
		}

		for _, user := range *res2 {
			s.Post(
				m.GetChannel(),
				fmt.Sprintf(
					"You are known as %s on Phabricator.",
					user.UserName,
				),
				messages.IconTasks,
				true,
			)
		}
	}
}
