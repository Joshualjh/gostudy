package main

import "main/pkg/router"

func main() {
	r := router.Router()

	r.Run(":8081")
}

// gin-gonic serving static files로 html파일을 사용할 수 있도록 서버를 작성
// 추가 개발 필요한 부분은 update, delete 부분을 하면 됩니다.
// 해당하는 부분은 put, delete 부분으로 작성을 진행
