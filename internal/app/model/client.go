package model

import "container/list"

type Client struct {
	ID              int
	ClientsSegments list.List
}
