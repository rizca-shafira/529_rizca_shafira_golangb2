Kode Peserta    : 149368582100-529
Nama            : Rizca Shafira S. M.

Tahap Pembuatan Final Project
1. Membuat folder mygram
2. Membuka directory di integrated terminal visual studio code
3. Inisiasi go mod
4. Instalasi package yang dibutuhkan
    "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
    "github.com/asaskevich/govalidator"
    "github.com/asaskevich/govalidator"
5. Menyiapkan database "db_go_mygram"
6. Membuat folder di dalam directory mygram
    assets -> berisi aset foto 
    controllers -> berisi controllers yang dibutuhkan
    database -> berisi setup koneksi ke database
    helpers -> berisi fungsi pendukung
    middlewares -> berisi fungsi middlewares yang digunakan
    models -> berisi model yang terhubung dengan tiap table database
    router -> berisi route endpoint yang terhubung dengan controller
    structs -> berisi struct yang dibutuhkan pada program
    uploaded -> folder static pada server, berisi foto-foto yang akan di upload
7. Membuat main.go berisi setup run router
