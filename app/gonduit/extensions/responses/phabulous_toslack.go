package responses

import "github.com/pjb7687/phabulous/app/gonduit/extensions/entities"

// PhabulousToSlackResponse is the response to calling phabulous.toslack.
type PhabulousToSlackResponse map[string]*entities.SlackUser
