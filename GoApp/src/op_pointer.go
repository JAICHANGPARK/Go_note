package main

func main() {

	var k int = 10
	var p = &k  //k의 주소를 할당
	println(*p) //p가 가리키는 주소에 있는 실제 내용을 출력

	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	println(sum)

}
