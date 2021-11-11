&http.Server 和 ListenAndServe？
1、http.Server：

type Server struct {
    Addr    string
    Handler Handler
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
Addr：监听的 TCP 地址，格式为:8000
Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
TLSConfig：安全传输层协议（TLS）的配置
ReadTimeout：允许读取的最大时间
ReadHeaderTimeout：允许读取请求头的最大时间
WriteTimeout：允许写入的最大时间
IdleTimeout：等待的最大时间
MaxHeaderBytes：请求头的最大字节数
ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）

2、 ListenAndServe：

func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求。

我们在源码中看到Addr是调用我们在&http.Server中设置的参数，因此我们在设置时要用&，我们要改变参数的值，
因为我们ListenAndServe和其他一些方法需要用到&http.Server中的参数，他们是相互影响的。




http.ListenAndServe和 连载一 的r.Run()有区别吗？

我们看看r.Run的实现：

func (engine *Engine) Run(addr ...string) (err error) {
    defer func() { debugPrintError(err) }()

    address := resolveAddress(addr)
    debugPrint("Listening and serving HTTP on %s\n", address)
    err = http.ListenAndServe(address, engine)
    return
}
通过分析源码，得知本质上没有区别，同时也得知了启动gin时的监听 debug 信息在这里输出。