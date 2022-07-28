package mocks

//go:generate mockgen -destination=generated.go -package=mocks github.com/pjb7687/phabulous/app/interfaces Bot,Message,HandlerTuple,Module,Connector
