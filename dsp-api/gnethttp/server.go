package gnethttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/panjf2000/gnet/v2"
)

// 使用闭包修复 ReadAll 函数避免循环导入

// HTTPHandler 处理 HTTP 请求的函数类型
type HTTPHandler func(req *Request) *Response

// Router HTTP 路由器
type Router struct {
	// 存储路由: method -> path -> handler
	routes map[string]map[string]HTTPHandler
	mu     sync.RWMutex
}

// NewRouter 创建新的路由器
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]HTTPHandler),
	}
}

// HandleFunc 注册路由处理函数
func (r *Router) HandleFunc(method, path string, handler HTTPHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	method = strings.ToUpper(method)
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]HTTPHandler)
	}
	r.routes[method][path] = handler
}

// GET 注册 GET 请求处理函数
func (r *Router) GET(path string, handler HTTPHandler) {
	r.HandleFunc("GET", path, handler)
}

// POST 注册 POST 请求处理函数
func (r *Router) POST(path string, handler HTTPHandler) {
	r.HandleFunc("POST", path, handler)
}
//
//// PUT 注册 PUT 请求处理函数
//func (r *Router) PUT(path string, handler HTTPHandler) {
//	r.HandleFunc("PUT", path, handler)
//}
//
//// DELETE 注册 DELETE 请求处理函数
//func (r *Router) DELETE(path string, handler HTTPHandler) {
//	r.HandleFunc("DELETE", path, handler)
//}
//
//// PATCH 注册 PATCH 请求处理函数
//func (r *Router) PATCH(path string, handler HTTPHandler) {
//	r.HandleFunc("PATCH", path, handler)
//}

// findHandler 查找对应的处理函数
func (r *Router) findHandler(method, path string) (HTTPHandler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	method = strings.ToUpper(method)
	if methodRoutes, ok := r.routes[method]; ok {
		if handler, ok := methodRoutes[path]; ok {
			return handler, true
		}
	}
	return nil, false
}

// Request HTTP 请求
type Request struct {
	Method      string
	Path        string
	Headers     map[string]string
	Body        []byte
	Query       map[string]string
	RemoteAddr  string
	contentType string
}

// PathValue 获取路径参数（用于未来扩展支持路径参数）
func (r *Request) PathValue(key string) string {
	return ""
}

// FormValue 获取表单值
func (r *Request) FormValue(key string) string {
	if r.contentType == "application/x-www-form-urlencoded" {
		values := strings.Split(string(r.Body), "&")
		for _, v := range values {
			parts := strings.SplitN(v, "=", 2)
			if len(parts) == 2 && parts[0] == key {
				return parts[1]
			}
		}
	}
	return ""
}

// JSON 将请求体解析为 JSON 到目标对象
func (r *Request) JSON(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

// Response HTTP 响应
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

// NewResponse 创建新的响应
func NewResponse(statusCode int) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type":   "text/plain; charset=utf-8",
			"Connection":     "close",
			"Server":         "GNet-HTTP/1.0",
			"X-Powered-By":   "gnethttp",
		},
	}
}

// OK 创建 200 响应
func OK() *Response {
	return NewResponse(200)
}

// JSON 创建 JSON 响应
func JSON(data interface{}) *Response {
	resp := NewResponse(200)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	body, err := json.Marshal(data)
	if err != nil {
		return InternalServerError(err.Error())
	}
	resp.Body = body
	return resp
}

//// HTML 创建 HTML 响应
//func HTML(html string) *Response {
//	resp := NewResponse(200)
//	resp.SetHeader("Content-Type", "text/html; charset=utf-8")
//	resp.Body = []byte(html)
//	return resp
//}
//
//// Text 创建文本响应
//func Text(text string) *Response {
//	resp := NewResponse(200)
//	resp.Body = []byte(text)
//	return resp
//}

// Created 创建 201 响应
func Created(data interface{}) *Response {
	resp := NewResponse(201)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return InternalServerError(err.Error())
		}
		resp.Body = body
	}
	return resp
}

// BadRequest 创建 400 响应
func BadRequest(message string) *Response {
	resp := NewResponse(400)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	body, _ := json.Marshal(map[string]string{"error": message})
	resp.Body = body
	return resp
}

// Unauthorized 创建 401 响应
func Unauthorized(message string) *Response {
	resp := NewResponse(401)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	body, _ := json.Marshal(map[string]string{"error": message})
	resp.Body = body
	return resp
}

// NotFound 创建 404 响应
func NotFound(message ...string) *Response {
	resp := NewResponse(404)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	msg := "Not Found"
	if len(message) > 0 {
		msg = message[0]
	}
	body, _ := json.Marshal(map[string]string{"error": msg})
	resp.Body = body
	return resp
}

// MethodNotAllowed 创建 405 响应
func MethodNotAllowed(message string) *Response {
	resp := NewResponse(405)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	body, _ := json.Marshal(map[string]string{"error": message})
	resp.Body = body
	return resp
}

// InternalServerError 创建 500 响应
func InternalServerError(message string) *Response {
	resp := NewResponse(500)
	resp.SetHeader("Content-Type", "application/json; charset=utf-8")
	body, _ := json.Marshal(map[string]string{"error": message})
	resp.Body = body
	return resp
}

// SetHeader 设置响应头
func (r *Response) SetHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[key] = value
}

// SetBody 设置响应体
func (r *Response) SetBody(body []byte) {
	r.Body = body
}

// toBytes 将响应转换为 HTTP 响应字节数组
func (r *Response) toBytes() []byte {
	var buf bytes.Buffer

	// 状态行
	buf.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", r.StatusCode, getStatusText(r.StatusCode)))

	// 响应头
	for key, value := range r.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	// Content-Length
	if len(r.Body) > 0 {
		buf.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(r.Body)))
	}

	// 空行分隔头部和主体
	buf.WriteString("\r\n")

	// 响应体
	buf.Write(r.Body)

	return buf.Bytes()
}

// getStatusText 根据 HTTP 状态码返回对应的状态文本
func getStatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 500:
		return "Internal Server Error"
	default:
		return http.StatusText(code)
	}
}

// Server HTTP 服务器
type Server struct {
	gnet.BuiltinEventEngine
	router      *Router
	addr        string
	middleware  []HTTPHandler
	staticDir   string
}

// ServerOption 服务器配置选项
type ServerOption func(*Server)

// WithAddr 设置监听地址
func WithAddr(addr string) ServerOption {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithMiddleware 添加中间件
func WithMiddleware(middleware ...HTTPHandler) ServerOption {
	return func(s *Server) {
		s.middleware = append(s.middleware, middleware...)
	}
}

// WithStaticDir 设置静态文件目录
func WithStaticDir(dir string) ServerOption {
	return func(s *Server) {
		s.staticDir = dir
	}
}

// NewServer 创建新的 HTTP 服务器
func NewServer(router *Router, opts ...ServerOption) *Server {
	s := &Server{
		router:     router,
		addr:       ":8888",
		middleware: make([]HTTPHandler, 0),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// OnOpen 当新的连接建立时被调用
func (s *Server) OnOpen(conn gnet.Conn) ([]byte, gnet.Action) {
	return nil, gnet.None
}

// OnClose 当连接关闭时被调用
func (s *Server) OnClose(conn gnet.Conn, err error) gnet.Action {
	return gnet.None
}

// OnTraffic 当有数据到达连接时被调用
func (s *Server) OnTraffic(conn gnet.Conn) gnet.Action {
	// 读取所有数据
	data, err := conn.Next(-1)
	if err != nil {
		return gnet.Close
	}

	// 解析 HTTP 请求
	req, err := parseHTTPRequest(data)
	if err != nil {
		conn.Write(NewResponse(400).toBytes())
		return gnet.Close
	}

	req.RemoteAddr = conn.RemoteAddr().String()

	// 查找处理函数
	handler, found := s.router.findHandler(req.Method, req.Path)
	if !found {
		resp := NotFound()
		conn.Write(resp.toBytes())
		return gnet.Close
	}

	// 执行中间件和处理器
	resp := s.executeHandler(req, handler)
	conn.Write(resp.toBytes())

	return gnet.Close
}

// executeHandler 执行处理函数和中间件
func (s *Server) executeHandler(req *Request, handler HTTPHandler) *Response {
	// 创建处理链
	handlers := append(s.middleware, handler)

	var resp *Response
	for i, h := range handlers {
		resp = h(req)
		if resp == nil {
			// 如果没有返回响应，继续执行下一个处理器
			if i == len(handlers)-1 {
				resp = OK()
			}
			continue
		}
	}

	return resp
}

// Start 启动服务器
func (s *Server) Start() error {
	return gnet.Run(s, "tcp://"+s.addr,
		gnet.WithMulticore(true),
		gnet.WithReuseAddr(true),
	)
}

// ListenAndServe 启动服务器（快捷方法）
func ListenAndServe(addr string, router *Router) error {
	return NewServer(router, WithAddr(addr)).Start()
}

// parseHTTPRequest 解析原始 HTTP 请求数据
func parseHTTPRequest(data []byte) (*Request, error) {
	// 查找请求头结束的位置（空行）
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	if headerEnd == -1 {
		headerEnd = bytes.Index(data, []byte("\n\n"))
		if headerEnd == -1 {
			return nil, fmt.Errorf("invalid HTTP request: no header delimiter")
		}
	}

	headerSection := data[:headerEnd]
	bodySection := data[headerEnd+4:] // +4 跳过 \r\n\r\n

	// 解析请求行
	headerLines := bytes.Split(headerSection, []byte("\n"))
	if len(headerLines) == 0 {
		return nil, fmt.Errorf("invalid HTTP request: no request line")
	}

	requestLine := strings.TrimSpace(string(headerLines[0]))
	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid request line format: %s", requestLine)
	}

	method := parts[0]
	path := parts[1]

	req := &Request{
		Method:  method,
		Path:    path,
		Headers: make(map[string]string),
		Body:    bodySection,
		Query:   make(map[string]string),
	}

	// 分离路径和查询参数
	if queryIdx := strings.Index(parts[1], "?"); queryIdx != -1 {
		req.Path = parts[1][:queryIdx]
		queryString := parts[1][queryIdx+1:]
		for _, param := range strings.Split(queryString, "&") {
			kv := strings.SplitN(param, "=", 2)
			if len(kv) == 2 {
				req.Query[kv[0]] = kv[1]
			}
		}
	}

	// 解析请求头
	for i := 1; i < len(headerLines); i++ {
		line := strings.TrimSpace(string(headerLines[i]))
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			req.Headers[key] = value

			if strings.EqualFold(key, "Content-Type") {
				req.contentType = value
			}
		}
	}

	return req, nil
}

// GetLocalIP 获取本机 IP 地址
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// PrintServerInfo 打印服务器启动信息
func PrintServerInfo(addr string) {
	ip := GetLocalIP()
	fmt.Println("========================================")
	fmt.Println("       GNet HTTP 服务器启动成功")
	fmt.Println("========================================")
	if ip != "" {
		fmt.Printf("本地访问: http://localhost%s\n", addr)
		fmt.Printf("局域网访问: http://%s%s\n", ip, addr)
	} else {
		fmt.Printf("服务器地址: http://%s\n", addr)
	}
	fmt.Println("========================================")
	fmt.Println()
}