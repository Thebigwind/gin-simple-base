1. 日志级别

2. 修改配置，自动加载

3. swagger 生成doc文档

4. pprof 性能分析

5. makefile

6.Dockerfile 参考 https://eddycjy.com/posts/go/gin/2018-03-24-golang-docker/
为什么 gin-blog-docker 占用空间这么大？（可用 docker ps -as | grep gin-blog-docker 查看）
Mysql 容器直接这么使用，数据存储到哪里去了？

Q：第一个问题，为什么这么镜像体积这么大？

A：FROM golang:latest 拉取的是官方 golang 镜像，包含 Golang 的编译和运行环境，外加一堆 GCC、build 工具，相当齐全

这是有问题的，我们可以不在 Golang 容器中现场编译的，压根用不到那些东西，我们只需要一个能够运行可执行文件的环境即可

构建 Scratch 镜像
Scratch 镜像，简洁、小巧，基本是个空镜像

