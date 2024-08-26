module example.com/vulnerable

go 1.18

require (
    github.com/gin-gonic/gin v1.6.0 // 这个版本存在已知的安全漏洞
)
