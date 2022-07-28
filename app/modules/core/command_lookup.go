package core

import (
	"fmt"
	"github.com/pjb7687/phabulous/app/interfaces"
	"github.com/pjb7687/phabulous/app/messages"
	"github.com/pjb7687/phabulous/app/utilities"
)

// LookupCommand allows users to lookup objects from Phabricator.
type LookupCommand struct {
	CustomMatchers []string
}

// GetDefaultMatchers returns the default matchers of the LookupCommand.
func (c *LookupCommand) GetDefaultMatchers() []string {
	defaultMatchers := []string{"^([T|D][0-9]{1,16})$"}
	return defaultMatchers
}

// GetUsage returns the usage of this command.
func (c *LookupCommand) GetUsage() string {
	return "(Txxx|Dxxx)"
}

// GetDescription returns the description of this command.
func (c *LookupCommand) GetDescription() string {
	return "Looks up a task or revision by its number."
}

// GetMatchers returns the matchers for this command.
func (c *LookupCommand) GetMatchers() []string {
	if c.CustomMatchers != nil {
		return c.CustomMatchers
	}

	return c.GetDefaultMatchers()
}

// GetIMMatchers returns IM matchers for this command.
func (c *LookupCommand) GetIMMatchers() []string {
	return []string{
		"^([T|D][0-9]{1,16})$",
	}
}

// GetMentionMatchers returns the channel mention matchers for this command.
func (c *LookupCommand) GetMentionMatchers() []string {
	return []string{
		"([T|D][0-9]{1,16})",
	}
}

// GetHandler returns the handler for this command.
func (c *LookupCommand) GetHandler() interfaces.Handler {
	return func(s interfaces.Bot, m interfaces.Message, matches []string) {
		s.StartTyping(m.GetChannel())

		conn, err := s.GetGonduit()
		if err != nil {
			s.Excuse(m, err)
			return
		}

		uniqueMatches := utilities.UniqueItemsOf(matches)

		for _, match := range uniqueMatches {
			res, err := conn.PHIDLookupSingle(match)
			if err != nil {
				s.Excuse(m, err)
			} else if res == nil {
				s.Post(
					m.GetChannel(),
					fmt.Sprintf("I couldn't find %s", match),
					messages.IconDefault,
					true,
				)
			} else {
				s.Post(
					m.GetChannel(),
					fmt.Sprintf("*%s* (%s): %s", res.FullName, res.Status, res.URI),
					messages.IconTasks,
					true,
				)
			}
		}
	}
}
