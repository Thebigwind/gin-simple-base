应用中间件

需要访问http://127.0.0.1:8000/auth?username=test&password=test123456，得到token

{
"code": 200,
"data": {
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE1MTg3MjQ2OTMsImlzcyI6Imdpbi1ibG9nIn0.KSBY6TeavV_30kfmP7HWLRYKP5TPEDgHtABe9HCsic4"
},
"msg": "ok"
}


再用包含token的 URL 参数去访问我们的应用 API，

访问http://127.0.0.1:8000/api/v1/articles?token=eyJhbGci.


