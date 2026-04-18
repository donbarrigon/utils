package handler

import (
	"strings"
)

type RouteNode struct {
	Part       string
	Controller ControllerFun
	Children   []RouteNode
	Keys       []string
	IsDynamic  bool
}

type Param struct {
	Key   string
	Value string
}

func NewRouteNode() *RouteNode {
	return &RouteNode{
		Children: []RouteNode{},
		Keys:     []string{},
	}
}

func (r *RouteNode) Add(path string, ctrl ControllerFun) {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	keys := []string{}

	node := r
	for i, part := range parts {
		isLast := i == len(parts)-1
		isDynamic := false
		if strings.Contains(part, ":") {
			isDynamic = true
			part = strings.Replace(part, ":", "", 1)
			keys = append(keys, part)
			part = "*"
		}

		childIndex := -1
		for j := range node.Children {
			if node.Children[j].Part == part && node.Children[j].IsDynamic == isDynamic {
				childIndex = j
				if isLast {
					node.Children[j].Controller = ctrl
				}
				break
			}
		}

		// si el nodo ya existe usarlo
		if childIndex >= 0 {
			node = &node.Children[childIndex]
			continue
		}

		// si no existe se crea el nuevo nodo

		newNode := RouteNode{
			Part:      part,
			Children:  []RouteNode{},
			Keys:      []string{},
			IsDynamic: isDynamic,
		}
		if isLast {
			newNode.Controller = ctrl
			newNode.Keys = keys
		}

		isInserted := false
		if !isDynamic {
			for j := range node.Children {
				if node.Children[j].IsDynamic {
					node.Children = append(node.Children[:j],
						append([]RouteNode{newNode}, node.Children[j:]...)...)
					node = &node.Children[j]
					isInserted = true
					break
				}
			}
		}

		if !isInserted {
			node.Children = append(node.Children, newNode)
			node = &node.Children[len(node.Children)-1]
		}
	}
}

func (r *RouteNode) Find(path string) ([]Param, ControllerFun) {
	// path = strings.Trim(path, "/") // ya se hace en el handler
	parts := strings.Split(path, "/")
	params := []Param{}

	node := r
	lastPartIndex := len(parts) - 1
	// i := 0 // index del children
	j := 0 // index del parts
	for i := 0; i < len(node.Children); i++ {
		if node.Children[i].Part == parts[j] || node.Children[i].IsDynamic {
			if node.Children[i].IsDynamic {
				params = append(params, Param{
					Value: parts[j],
				})
			}
			if lastPartIndex == j {
				for k := range params {
					params[k].Key = node.Children[i].Keys[k]
				}
				return params, node.Children[i].Controller
			}
			node = &node.Children[i]
			i = -1
			j++
		}
	}
	return nil, nil
}
