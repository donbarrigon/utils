package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/donbarrigon/utils/herror"
	"github.com/donbarrigon/utils/logs"
)

type ControllerFun func(ctx *Context)
type MiddlewareFun func(ControllerFun) ControllerFun
type IRouter interface {
	Name(name string)
	Namef(name string)
}

type Router struct {
	// http.Handler
	// prefix/path/method -> controller
	Routes *RouteNode
	// default method.prefix.path -> route
	RoutesNames map[string]string
	middlewares []MiddlewareFun
	prefixes    []string
	lastRoute   string
	lastName    string
}

func NewRouter() *Router {
	return &Router{
		Routes:      NewRouteNode(),
		RoutesNames: map[string]string{},
		middlewares: []MiddlewareFun{},
		prefixes:    []string{},
	}
}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r, h)
	defer func() {
		if er := recover(); er != nil {
			e := herror.InternalServerError(er)
			msg, er2 := json.Marshal(e)
			if er2 != nil {
				msg = []byte{}
			}
			go logs.Critical("Panic: %s", string(msg))
			c.ResponseError(e)
			return
		}
	}()

	route := strings.ToLower(strings.Trim(r.URL.Path, "/")) + "/" + r.Method
	params, ctrl := h.Routes.Find(route)
	if ctrl != nil {
		c.Params = params
		ctrl(c)
		return
	} else {
		c.ResponseNotFound()
		return
	}

}

func (h *Router) Get(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodGet, path, ctrl, middlewares...)
	return h
}

func (h *Router) Post(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodPost, path, ctrl, middlewares...)
	return h
}

func (h *Router) Put(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodPut, path, ctrl, middlewares...)
	return h
}

func (h *Router) Patch(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodPatch, path, ctrl, middlewares...)
	return h
}

func (h *Router) Delete(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodDelete, path, ctrl, middlewares...)
	return h
}

func (h *Router) Head(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodHead, path, ctrl, middlewares...)
	return h
}

func (h *Router) Options(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodOptions, path, ctrl, middlewares...)
	return h
}

func (h *Router) Connect(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodConnect, path, ctrl, middlewares...)
	return h
}

func (h *Router) Trace(path string, ctrl ControllerFun, middlewares ...MiddlewareFun) IRouter {
	h.add(http.MethodTrace, path, ctrl, middlewares...)
	return h
}

func (h *Router) Use(mws ...MiddlewareFun) {
	h.middlewares = append(mws, h.middlewares...)
}

func (h *Router) Prefix(prefix string, callback func(), mws ...MiddlewareFun) {
	pfs := strings.Split(strings.Trim(prefix, "/"), "/")

	h.prefixes = append(h.prefixes, pfs...)
	h.middlewares = append(h.middlewares, mws...)
	callback()
	h.prefixes = h.prefixes[:len(h.prefixes)-len(pfs)]
	h.middlewares = h.middlewares[:len(h.middlewares)-len(mws)]
}

func (h *Router) Group(callback func(), mws ...MiddlewareFun) {
	h.middlewares = append(h.middlewares, mws...)
	callback()
	h.middlewares = h.middlewares[:len(h.middlewares)-len(mws)]
}

// el name se concatena asi prefix.name
func (h *Router) Name(name string) {
	delete(h.RoutesNames, h.lastName)
	name = strings.Trim(strings.Join(h.prefixes, ".")+"."+name, ".")
	h.RoutesNames[name] = h.lastRoute
}

// name es el mismo que ingrse ingresoesaste
func (h *Router) Namef(name string) {
	delete(h.RoutesNames, h.lastName)
	h.RoutesNames[name] = h.lastRoute
}

func (h *Router) add(method string, path string, ctrl ControllerFun, mws ...MiddlewareFun) {
	prefix := strings.Join(h.prefixes, "/")
	route := strings.ToLower(prefix+"/"+strings.Trim(path, "/")) + "/" + method

	mws = append(h.middlewares, mws...)
	mwCopy := make([]MiddlewareFun, len(mws))
	copy(mwCopy, mws)
	for i := len(mwCopy) - 1; i >= 0; i-- {
		ctrl = mwCopy[i](ctrl)
	}

	parts := strings.Split(strings.Trim(path, "/"), "/")
	for i, p := range parts {
		if strings.HasPrefix(p, ":") {
			parts[i] = "*"
		}
	}
	name := method + "." + strings.Join(h.prefixes, ".") + "." + strings.Join(parts, "/")
	name = strings.ToLower(name)
	routeForName := strings.ToLower(prefix+"/"+strings.Join(parts, "/")) + "/" + method

	h.Routes.Add(route, ctrl)
	h.RoutesNames[name] = routeForName

	h.lastRoute = routeForName
	h.lastName = name
}

func (h *Router) PrintRoutes() {
	printNode(h.Routes, "")
}

func printNode(node *RouteNode, indent string) {
	for _, child := range node.Children {
		if child.Controller != nil {
			fmt.Printf("%s/%s *\n", indent, child.Part)
		} else {
			fmt.Printf("%s/%s\n", indent, child.Part)
		}
		printNode(&child, indent+"  ")
	}
}
