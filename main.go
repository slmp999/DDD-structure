package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"

	//"log"
	"net/http"
	"time"

	"github.com/acoshift/middleware"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gitlab.com/satit13/perfect_api/admin"
	"gitlab.com/satit13/perfect_api/auth"
	log "gitlab.com/satit13/perfect_api/logger"
	"gitlab.com/satit13/perfect_api/member"
	"gitlab.com/satit13/perfect_api/mysql"
	"gitlab.com/satit13/perfect_api/order"
	"gitlab.com/satit13/perfect_api/product"
	"gitlab.com/satit13/perfect_api/report"
	"gitlab.com/satit13/perfect_api/sales"
	"gitlab.com/satit13/perfect_api/upload"
	//"log"
)

var (
	dbFile      = "hostdb"
	sqlFile     = "paybox.db"
	mode        = "dev"
	Mode        = "Develop"
	appPort     = "8080"
	Version     = "undefined"
	BuildTime   = "undefined"
	GitHash     = "undefined"
	logFlag     = flag.String("l", "debug", "กำหนดระดับ log -> info, warn, error, fatal, panic")
	proFlag     = flag.Bool("p", false, "รันในโหมดโปรดักชั่น ใช้งานจริง ถ้าไม่ใส่โปรแกรมจะไม่เปิดอุปกรณ์รับเงิน")
	devFlag     = flag.Bool("d", false, "รันในโหมด ไม่ได้ใช้งานจริง")
	versionFlag = flag.Bool("v", false, "show version info")
)

var mysqls *sql.DB
var mysqlx *sqlx.DB
var (
	gEnv     = "development" //default
	gSSLMode = "disable"
	bPort    = "3306"
	myHost   = "perfect_db.extensionsoft.biz"
	myUser   = "perfect"
	myDb     = "perfect"
	myPass   = "P@ssw0rd"
)

func ConnectDBX(user string, dbName string, host string, pass string, port string) (db *sqlx.DB, err error) {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&charset=utf8&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectDB(user string, dbName string, host string, pass string, port string) (db *sql.DB, err error) {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FBangkok&charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql error =", err)
		return nil, err
	} else {
		log.Println("connect sucess", dsn)
	}
	db.Exec("use " + dbName)

	return db, nil
}

func init() {
	flag.Parse()
	log.Printf("#### Version: %s", Version)
	log.Infoln("#### Build Time: %s", BuildTime)
	log.Printf("#### Git Hash: %s", GitHash)

	switch *logFlag {
	case "pro":
		Mode = "Production"
		mode = "Production"
		appPort = "8080"
		bPort = "3306"
		myHost = "perfect_db.extensionsoft.biz"
		myUser = "perfect"
		myDb = "perfect"
		myPass = "P@ssw0rd"
	default:
		Mode = "Develop"
		mode = "Develop"
		appPort = "8081"
		bPort = "3306"
		myHost = "perfect_db.extensionsoft.biz"
		myUser = "perfect"
		myDb = "perfect_demo"
		myPass = "P@ssw0rd"
	}
	mysql_perfect, err := ConnectDB(myUser, myDb, myHost, myPass, bPort)
	if err != nil {
		fmt.Println(err.Error())
	}
	mysqls = mysql_perfect
	fmt.Println("[Mode] : ### APP Mode = " + mode + "  ### ")
}

func main() {

	saleRepo, err := mysql.NewSaleRepository(mysqls)
	must(err)
	saleService, err := sales.NewSales(saleRepo)
	must(err)

	adminRepo, err := mysql.NewAdminRepository(mysqls)
	must(err)
	adminService, err := admin.NewService(adminRepo)
	must(err)
	authRepo, err := mysql.NewAuthRepository(mysqls)
	must(err)
	authService, err := auth.NewService(authRepo, Mode)
	must(err)

	memberRepo, err := mysql.NewMemberRepository(mysqls)
	must(err)
	memberService, err := member.NewService(memberRepo)
	must(err)

	// productRepo, err := mysqls.NewProductRepository(mysql)
	// must(err)
	// productService, err := product.NewService(productRepo)
	// must(err)
	productRepo, err := mysql.NewProductRepository(mysqls)
	must(err)
	productService, err := product.NewService(productRepo)
	must(err)
	orderRepo, err := mysql.NewOrderRepository(mysqls)
	must(err)
	orderService, err := order.NewService(orderRepo)
	must(err)
	uploadRepo, err := mysql.NewUploadRepository(mysqls)
	must(err)
	uploadService, err := upload.NewService(uploadRepo, Mode)
	must(err)
	reprtRepo, err := mysql.NewReportRepository(mysqls)
	must(err)
	reportService, err := report.NewReport(reprtRepo)
	must(err)
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheckHandler)
	mux.HandleFunc("/version", apiVersionHandler)

	fmt.Println("API Running  Port: " + appPort)
	staticDir := http.FileServer(http.Dir("/app/image"))
	staticDemo := http.FileServer(http.Dir("/app/image-demo"))
	DocDir := http.FileServer(http.Dir("/var/www/html/perfect_doc"))

	mux.Handle("/image/", http.StripPrefix("/image/", staticDir))
	mux.Handle("/image-demo/", http.StripPrefix("/image-demo/", staticDemo))
	mux.Handle("/doc/", http.StripPrefix("/doc/v1", DocDir))
	mux.Handle("/auth/", http.StripPrefix("/auth/v1", auth.MakeHandler(authService)))
	mux.Handle("/member/", http.StripPrefix("/member/v1", member.MakeHandler(memberService)))
	mux.Handle("/product/", http.StripPrefix("/product/v1", product.MakeHandler(productService)))
	mux.Handle("/order/", http.StripPrefix("/order/v1", order.MakeHandler(orderService)))
	mux.Handle("/upload/", http.StripPrefix("/upload/v1", upload.MakeHandler(uploadService)))
	mux.Handle("/admin/", http.StripPrefix("/admin/v1", admin.MakeHandler(adminService)))
	mux.Handle("/sales/", http.StripPrefix("/sales/v1", sales.MakeHandler(saleService)))
	mux.Handle("/report/", http.StripPrefix("/report/v1", report.MakeHandler(reportService)))
	corsConfig := middleware.CORSConfig{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
			"Token",
			"Access-Token",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Accept",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
		},
		MaxAge: time.Hour,
	}
	h := auth.MakeMiddleware(authService)(mux)
	h = middleware.CORS(corsConfig)(h)
	h = middleware.AddHeader("Content-Type", "application/json; charset=utf-8")(h)
	// mux.Handle("/product/", http.StripPrefix("/product/v1", product.MakeHandler(productService)))
	http.ListenAndServe(":"+appPort, h)
}

func must(err error) {
	if err != nil {
		log.Println("Error:", err)
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Success bool `json:"api success"`
	}{true})
}

type logEntry struct {
	Time         string `json:"time"`
	RemoteIP     string `json:"remote_ip"`
	Host         string `json:"host"`
	Method       string `json:"method"`
	URI          string `json:"uri"`
	UserAgent    string `json:"user_agent"`
	Status       int    `json:"status"`
	Latency      int64  `json:"latency"`
	LatencyHuman string `json:"latency_human"`
	BytesIn      int64  `json:"bytes_in"`
	BodyIn       string `json:"body_in"`
	BytesOut     int64  `json:"bytes_out"`
	BodyOut      string `json:"body_out"`
}

func apiVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//t := time.Now()
	json.NewEncoder(w).Encode(struct {
		Version     string `json:"version"`
		Description string `json:"description"`
		Creator     string `json:"creator"`
		LastUpdate  string `json:"lastupdate"`
	}{
		"0.1.0 BETA",
		"Perfect Application Service",
		"DevOpt",
		"2019-08-01",
	})
}
