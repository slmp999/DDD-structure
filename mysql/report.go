package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	//"github.com/labstack/gommon/log"

	log "gitlab.com/satit13/perfect_api/logger"
	logg "gitlab.com/satit13/perfect_api/logger"
	"gitlab.com/satit13/perfect_api/report"

	"upper.io/db.v3/lib/sqlbuilder"

	//"upper.io/db.v3/postgresql"
	"upper.io/db.v3/mysql"
)

// NewAuthRepository creates new auth repository
func NewReportRepository(db *sql.DB) (report.Repository, error) {
	// fmt.Println(1)
	pdb, err := mysql.New(db)
	if err != nil {
		logg.Error(err.Error())
		return nil, err
	}
	dbx := sqlx.NewDb(db, "mysql")
	// fmt.Println(1)
	r := reportRepository{pdb, dbx}
	// fmt.Println(1)
	return &r, nil
}

type reportRepository struct {
	db  sqlbuilder.Database
	dbx *sqlx.DB
}

func (r *reportRepository) SumReportSale(startdate string, enddate string) (float64, float64, error) {
	var amount, count float64 = 0, 0
	sql1 := `select ifnull(sum(sum_cash_amount+sum_credit_amount+sum_bank_amount+sum_coupon_amount),0) as amount,
	count(*) as count from orders where  order_status <> 1 and order_status <> 99 and
	DATE_FORMAT(doc_date, "%Y-%m-%d") BETWEEN ? AND ?  `

	rs, err := r.db.QueryRow(sql1, startdate, enddate)
	if err != nil {
		return 0, 0, err
	}
	err = rs.Scan(&amount, &count)
	if err != nil {
		return 0, 0, err
	}
	return amount, count, nil
}

func (r *reportRepository) ListReportSaleRepo(startdate string, enddate string) (interface{}, error) {
	sql1 := `select a.doc_no,
	a.user_id,
	concat(ifnull(b.user_fname,''),' ',ifnull(b.user_fname,'')) as user_name,
	ifnull(DATE_FORMAT(a.doc_date, "%Y-%m-%d %H:%i:%S"),'1990-01-01 00:00:00') as create_time,
	ifnull((a.sum_cash_amount+a.sum_credit_amount+a.sum_bank_amount+a.sum_coupon_amount),0) as amount,
	case when a.sum_cash_amount != 0 then 'เงินสด' 
		 when a.sum_credit_amount != 0 then 'เครดิต'
		 when a.sum_bank_amount != 0 then 'ธนาคาร'
		 when a.sum_coupon_amount != 0 then 'คูปอง'  end as sale_type 
	from orders a 
	inner join users b on a.user_id = b.user_id
	where  a.order_status <> 1 and a.order_status <> 99 and
	DATE_FORMAT(a.doc_date, "%Y-%m-%d") BETWEEN ? AND ? order by a.doc_date asc`
	models := []report.ModelReportSale{}
	rs, err := r.db.Query(sql1, startdate, enddate)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for rs.Next() {
		model := report.ModelReportSale{}
		err := rs.Scan(&model.DocNo,
			&model.UserID,
			&model.UserName,
			&model.CreateTime,
			&model.Amount,
			&model.SaleType)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		model.CommisionPercent = 10
		model.CommisionAmount = (model.Amount * float64(model.CommisionPercent)) / 100.00
		models = append(models, model)
	}
	return models, nil
}
