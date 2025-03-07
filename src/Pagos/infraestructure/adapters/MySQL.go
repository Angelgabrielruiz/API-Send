package adapters

import (
	"Send/src/core"
	"log"
)

type MySQLPago struct {
    conn *core.Conn_MySQL
}

func NewMySQLPago() *MySQLPago {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }
    return &MySQLPago{conn: conn}
}

func (mysql *MySQLPago) Save(monto int32, pago int32, cambio int32, fecha string) error {
    query := "INSERT INTO ticket (monto, pago, cambio, fecha) VALUES (?, ?, ?, ?)"
    result, err := mysql.conn.ExecutePreparedQuery(query, monto, pago, cambio, fecha)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Pago guardado: %d", rowsAffected)
    }

    return nil
}

func (mysql *MySQLPago) GetAll() ([]map[string]interface{}, error) {
    query := "SELECT * FROM ticket"
    rows := mysql.conn.FetchRows(query)
    defer rows.Close()

    var pagos []map[string]interface{}
    for rows.Next() {
        var id int32
        var monto int32
        var pago int32
        var cambio int32
		var fecha string
        if err := rows.Scan(&id, &monto, &pago, &cambio, &fecha); err != nil {
            return nil, err
        }
        Pago := map[string]interface{}{
            "id":     id,
            "monto": monto,
            "pago": pago,
            "cambio": cambio,
			"fecha": fecha,
        }
        pagos = append(pagos, Pago)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return pagos, nil
}


func (mysql *MySQLPago) Update(id int, monto int32, pago int32, cambio int32, fecha string) error {
    query := "UPDATE ticket SET monto = ?, pago = ?, cambio = ?, fecha = ? WHERE id = ?"
    result, err := mysql.conn.ExecutePreparedQuery(query, monto, pago, cambio, fecha, id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Pago actualizado: %d", rowsAffected)
    }

    return nil
}

func (mysql *MySQLPago) Delete(id int) error {
    query := "DELETE FROM ticket WHERE id = ?"
    result, err := mysql.conn.ExecutePreparedQuery(query, id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 1 {
        log.Printf("[MySQL] - Pago eliminado: %d", rowsAffected)
    }

    return nil
}

