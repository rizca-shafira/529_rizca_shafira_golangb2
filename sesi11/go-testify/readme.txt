go get github.com/stretchr/testify

agar tdk banyak kondisional if else, perlu mekanisme assertion
caranya menggunakan testify
testify adalah suatu library pada Go untuk melakukan assertion, mocking dan masih banyak lainnya.

untuk menjalankan unit test
go test ./helper

utk menjalankan 1 unit test saja
go test -v ./helper -run=TestTableSum.