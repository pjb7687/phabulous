package app

import (
	"github.com/pjb7687/phabulous/app/interfaces"
	"github.com/pjb7687/phabulous/app/messages"
)

// A ConnectorManager can be seen as an aggregator of Connector objects. It
// implements the Connector interface by delegating every method to its
// registered connectors. This allows clients to use multiple connectors at
// once with ease.
//
// While one could technically register a ConnectorManager as part of another
// manager, it is not extremely useful at this point.
//
type ConnectorManager struct {
	connectors []interfaces.Connector
}

// RegisterConnector adds a Connector to the list of connectors managed by this
// manager.
func (c *ConnectorManager) RegisterConnector(connector interfaces.Connector) {
	c.connectors = append(c.connectors, connector)
}

// LoadModules loads the provided modules on all the registered connectors.
func (c *ConnectorManager) LoadModules(modules []interfaces.Module) {
	for _, connector := range c.connectors {
		connector.LoadModules(modules)
	}
}

// Boot boots all the registered connectors. If one of them fails, the
// operation is interrupted and the first error is returned.
func (c *ConnectorManager) Boot() error {
	for _, connector := range c.connectors {
		result := connector.Boot()

		if result != nil {
			return result
		}
	}

	return nil
}

// Post posts a text message.
func (c *ConnectorManager) Post(
	channelName string,
	storyText string,
	icon messages.Icon,
	asUser bool,
) {
	for _, connector := range c.connectors {
		connector.Post(channelName, storyText, icon, asUser)
	}
}

// PostImage posts a message with an attached image.
func (c *ConnectorManager) PostImage(
	channelName string,
	storyText string,
	imageURL string,
	icon messages.Icon,
	asUser bool,
) {
	for _, connector := range c.connectors {
		connector.PostImage(channelName, storyText, imageURL, icon, asUser)
	}
}

// PostOnFeed posts a message on the bot's "feed" channel.
func (c *ConnectorManager) PostOnFeed(storyText string) error {
	for _, connector := range c.connectors {
		result := connector.PostOnFeed(storyText)

		if result != nil {
			return result
		}
	}

	return nil
}
