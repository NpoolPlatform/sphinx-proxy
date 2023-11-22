// nolint
package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	servicename "github.com/NpoolPlatform/sphinx-proxy/pkg/servicename"

	"github.com/google/uuid"
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 5
	maxIdle      = 2
	MaxLife      = 0
	keyServiceID = "serviceid"
)

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Warnw("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

func tables(ctx context.Context, dbName string, tx *sql.DB) ([]string, error) {
	tables := []string{}
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select table_name from information_schema.tables where table_schema = '%v'", dbName),
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		table := []byte{}
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, string(table))
	}
	logger.Sugar().Warnw(
		"tables",
		"Tables", tables,
	)
	return tables, nil
}

func existIDInt(ctx context.Context, dbName, table string, tx *sql.DB) (bool, bool, bool, error) {
	_type := []byte{}
	rc := 0

	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select column_type,1 from information_schema.columns where table_name='%v' and column_name='id' and table_schema='%v'", table, dbName),
	)
	if err != nil {
		return false, false, false, err
	}
	for rows.Next() {
		if err := rows.Scan(&_type, &rc); err != nil {
			return false, false, false, err
		}
	}
	return rc == 1, strings.Contains(string(_type), "int"), strings.Contains(string(_type), "unsigned"), nil
}

func setIDUnsigned(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"setIDUnsigned",
		"db", dbName,
		"table", table,
		"State", "INT UNSIGNED",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v change id id int unsigned not null auto_increment", dbName, table),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"setIDUnsigned",
		"db", dbName,
		"table", table,
		"State", "INT UNSIGNED",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func existEntIDUnique(ctx context.Context, dbName, table string, tx *sql.DB) (bool, bool, error) {
	key := ""
	rc := 0
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select column_key,1 from information_schema.columns where table_schema='%v' and table_name='%v' and column_name='ent_id'", dbName, table),
	)
	if err != nil {
		return false, false, err
	}
	for rows.Next() {
		if err := rows.Scan(&key, &rc); err != nil {
			return false, false, err
		}
	}
	return rc == 1, key == "UNI", nil
}

func setEmptyEntID(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"setEmptyEntID",
		"db", dbName,
		"table", table,
		"State", "ENT_ID UUID",
	)
	rows, err := tx.QueryContext(
		ctx,
		fmt.Sprintf("select id, ent_id from %v.%v", dbName, table),
	)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id uint32
		var entID string
		if err := rows.Scan(&id, &entID); err != nil {
			return err
		}
		if _, err := uuid.Parse(entID); err == nil {
			continue
		}
		result, err := tx.ExecContext(
			ctx,
			fmt.Sprintf("update %v.%v set ent_id='%v' where id=%v", dbName, table, uuid.New(), id),
		)
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		logger.Sugar().Warnw(
			"setEmptyEntID",
			"db", dbName,
			"table", table,
			"State", "ENT_ID UUID",
			"rowsAffected", rowsAffected,
		)
	}
	return nil
}

func setEntIDUnique(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"setEntIDUnique",
		"db", dbName,
		"table", table,
		"State", "ENT_ID UNIQUE",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v change column ent_id ent_id char(36) unique", dbName, table),
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"setEntIDUnique",
		"db", dbName,
		"table", table,
		"State", "ENT_ID UNIQUE",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func id2EntID(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"id2EntID",
		"db", dbName,
		"table", table,
		"State", "ID -> EntID",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v change column id ent_id char(36) unique", dbName, table),
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"id2EntID",
		"db", dbName,
		"table", table,
		"State", "ID -> EntID",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func addIDColumn(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"addIDColumn",
		"db", dbName,
		"table", table,
		"State", "ID INT",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v add id int unsigned not null auto_increment, add primary key(id)", dbName, table),
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"addIDColumn",
		"db", dbName,
		"table", table,
		"State", "ID INT",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func dropPrimaryKey(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"dropPrimaryKey",
		"db", dbName,
		"table", table,
		"State", "DROP PRIMARY",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v drop primary key", dbName, table),
	)
	if err != nil {
		logger.Sugar().Warnw(
			"dropPrimaryKey",
			"db", dbName,
			"table", table,
			"Error", err,
		)
		return nil
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"dropPrimaryKey",
		"db", dbName,
		"table", table,
		"State", "DROP PRIMARY",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func addEntIDColumn(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"addEntIDColumn",
		"db", dbName,
		"table", table,
		"State", "ADD ENT_ID",
	)
	result, err := tx.ExecContext(
		ctx,
		fmt.Sprintf("alter table %v.%v add ent_id char(36) not null", dbName, table),
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"addEntIDColumn",
		"db", dbName,
		"table", table,
		"State", "ADD ENT_ID",
		"rowsAffected", rowsAffected,
	)
	return nil
}

func migrateEntID(ctx context.Context, dbName, table string, tx *sql.DB) error {
	logger.Sugar().Warnw(
		"migrateEntID",
		"db", dbName,
		"table", table,
	)

	idExist, idInt, idUnsigned, err := existIDInt(ctx, dbName, table, tx)
	if err != nil {
		return err
	}

	if idInt && !idUnsigned {
		if err := setIDUnsigned(ctx, dbName, table, tx); err != nil {
			return err
		}
	}

	entIDExist, entIDUnique, err := existEntIDUnique(ctx, dbName, table, tx)
	if err != nil {
		return err
	}

	if entIDExist && idInt {
		if err := setEmptyEntID(ctx, dbName, table, tx); err != nil {
			return err
		}
		if !entIDUnique {
			if err := setEntIDUnique(ctx, dbName, table, tx); err != nil {
				return err
			}
		}
		return nil
	}
	if idExist && !idInt {
		if err := id2EntID(ctx, dbName, table, tx); err != nil {
			return err
		}
	}
	if !idInt {
		if err := dropPrimaryKey(ctx, dbName, table, tx); err != nil {
			return err
		}
		if err := addIDColumn(ctx, dbName, table, tx); err != nil {
			return err
		}
	}

	entIDExist, _, err = existEntIDUnique(ctx, dbName, table, tx)
	if err != nil {
		return err
	}
	if !entIDExist {
		if err := addEntIDColumn(ctx, dbName, table, tx); err != nil {
			return err
		}
	}
	if err := setEmptyEntID(ctx, dbName, table, tx); err != nil {
		return err
	}
	if err := setEntIDUnique(ctx, dbName, table, tx); err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"migrateEntID",
		"db", dbName,
		"table", table,
		"State", "Migrated",
	)
	return err
}

func Migrate(ctx context.Context) error {
	var err error
	var conn *sql.DB

	logger.Sugar().Warnw("Migrate", "Start", "...")
	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}
	defer func(err *error) {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Warnw("Migrate", "Done", "...", "error", *err)
	}(&err)

	conn, err = open(servicename.ServiceDomain)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Sugar().Errorw("Close", "Error", err)
		}
	}()

	dbname := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyDBName)
	_tables, err := tables(ctx, dbname, conn)
	if err != nil {
		return err
	}

	logger.Sugar().Warnw(
		"Migrate",
		"Round", 1,
	)
	for _, table := range _tables {
		if err = migrateEntID(ctx, dbname, table, conn); err != nil {
			return err
		}
	}

	logger.Sugar().Warnw(
		"Migrate",
		"Round", 2,
	)
	for _, table := range _tables {
		if err = migrateEntID(ctx, dbname, table, conn); err != nil {
			return err
		}
	}
	return nil
}
