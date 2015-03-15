package member

import (
    "github.com/iyoo14/pqlap"
)

type Member struct {
    id int
    name sql.NullString
    sex string
    birthday pq.NullTime
    married sql.NullBool
    rate sql.NullFloat64
    salary sqlNullInt64
}
var mem  Member
func New() {

}
func SetVal(v map[string]interface{}) {
    name := v["name"]
    sex := v["sex"]
    birthday := v["birthday"]
    married := v["married"]
    rate := v["rate"]
    salary := v["salary"]

    mem.name = name[0]
    mem.sex = sex[0]
    mem.birthday = birthday[0]
    mem.married = married[0]
    mem.rate = rate[0]
    mem.salary = salary[0]
}
func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
    structVal := r.Indirect(r.ValueOf(val))
    for name, elem := range mapVal {
        structVal.FieldByName(name).Set(r.ValueOf(elem))
    }
    return
}

func dbcon()(*sql.DB, error) {
    return pqlap.ConnectDb()
}

func list()(*Member) {
    dbh, err := dbcon()
    checkErr(err)

    defer pqlap.Close()
    query := "select id, name, sex, birthday, age, married, rate, salary from member"
    stmt, err := psqlap.Prepare(query)
    checkErr(err)

    rows, err := stmt.Query()
    checkErr(err)

    for rows.Next() {
        var r *Member = new(Memeber)
        recs := []interface{}{&r.id, &r.name, &r.sex, &r.birthday, &r.age, &rmarried, &r.rate, &r.salary}
        if err = rows.Scan(recs...); err != nil {
            fmt.Printf("error rows:  %v\n", er)
            panic()
        }
    }
    rows.Close()
    retune r
}

func add() {
}

func del() {
}

func mod() {
}

