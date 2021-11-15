package main

func main() {

}

/*
http.Redirect(w ResponseWriter, r *Request, url string, code int)：这个函数会让浏览器重定向到 url（是请求的 url 的相对路径）以及状态码。
http.NotFound(w ResponseWriter, r *Request)：这个函数将返回网页没有找到，HTTP 404 错误。
http.Error(w ResponseWriter, error string, code int)：这个函数返回特定的错误信息和 HTTP 代码。
另 http.Request 对象的一个重要属性 req：req.Method，这是一个包含 GET 或 POST 字符串，用来描述网页是以何种方式被请求的


你可以使用 `w.header ().Set ("Content-Type","../..") 设置头信息
比如在网页应用发送 html 字符串的时候，在输出之前执行 w.Header().Set("Content-Type", "text/html")。

*/
