package rqProperty

import (
	"fmt"
	"kostjc/model/zCrud"
	"sort"
	"strconv"
	"time"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"
)

func (l *Locations) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Locations) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + l.SqlTableName() + whereAndSql + `
LIMIT 1`
	l.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + l.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	l.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Locations) FindAll() []Locations {
	const comment = `-- Locations) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0`

	var rows = []Locations{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (f *Facilities) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Facilities) FindByPagination`

	validFields := FacilitiesFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + f.SqlTableName() + whereAndSql + `
LIMIT 1`
	f.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + f.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	f.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Buildings) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Buildings) FindByPagination`

	validFields := BuildingsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	b.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + b.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	b.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Locations) FindLocationChoices() map[uint64]string {
	const comment = `-- Locations) FindLocationChoices`

	queryRows := comment + `
SELECT ` + l.SqlId() + `, ` + l.SqlName() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + l.SqlName() + ` ASC`

	out := make(map[uint64]string)
	l.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (b *Bookings) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Bookings) FindByPagination`

	validFields := BookingsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	b.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	if S.Contains(orderBySql, `ORDER BY "id"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "id"`, `ORDER BY "bookings"."id"`)
	}

	if S.Contains(orderBySql, `ORDER BY "createdAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "createdAt"`, `ORDER BY "bookings"."createdAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "updatedAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "updatedAt"`, `ORDER BY "bookings"."updatedAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "deletedAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "deletedAt"`, `ORDER BY "bookings"."deletedAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants" DESC`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants" DESC`, ``)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants" ASC`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants" ASC`, ``)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants"`, ``)
	}

	queryRows := comment + `
SELECT
	"bookings"."id",
	"bookings"."roomId",
	COALESCE(SUM("payments"."paidIDR"), 0) AS "totalPaidIDR",
	"bookings"."dateStart",
	"bookings"."dateEnd",
	"bookings"."tenantId",
	"bookings"."basePriceIDR",
	"bookings"."facilitiesObj",
	"bookings"."totalPriceIDR",
	"bookings"."paidAt",
	"bookings"."extraTenants",
	"bookings"."createdAt",
	"bookings"."updatedAt",
	"bookings"."deletedAt"
FROM ` + b.SqlTableName() + `
LEFT JOIN "payments" ON "payments"."bookingId" = "bookings"."id"
` + whereAndSql + `
GROUP BY "bookings"."id"
` + orderBySql + limitOffsetSql

	b.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (f *Facilities) FindAll() []Facilities {
	const comment = `-- Facilities) FindAll`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName()

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindAllTypeBuilding() []Facilities {
	const comment = `-- Facilities) FindAllTypeBuilding`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Building'
	AND "deletedAt" = 0`

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindAllTypeRoom() []Facilities {
	const comment = `-- Facilities) FindAllTypeRoom`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Room'
	AND "deletedAt" = 0`

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindFacilitiesChoices() map[uint64]string {
	const comment = `-- Locations) FindLocationChoices`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlFacilityName() + `, ` + f.SqlFacilityType() + ` FROM ` + f.SqlTableName() + `
ORDER BY ` + f.SqlFacilityName() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 3 {
			facilityName := X.ToS(row[1])
			facilityType := X.ToS(row[2])
			if facilityType != `` {
				facilityName += ` (` + facilityType + `)`
			}

			out[X.ToU(row[0])] = facilityName
		}
	})

	return out
}

func (f *Facilities) FindFacilitiesBuildingChoices() map[uint64]string {
	const comment = `-- Locations) FindFacilitiesBuildingChoices`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlFacilityName() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Building'
	AND "deletedAt" = 0
ORDER BY ` + f.SqlFacilityName() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (p *Payments) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Payments) FindByPagination`

	validFields := PaymentsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Buildings) FindBuildingChoices() map[uint64]string {
	const comment = `-- Buildings) FindBuildingChoices`

	queryRows := comment + `
SELECT ` + b.SqlId() + `, ` + b.SqlBuildingName() + ` FROM ` + b.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + b.SqlId() + ` ASC`

	out := make(map[uint64]string)
	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (b *Bookings) FindBookingChoices() map[uint64]string {
	const comment = `-- Bookings) FindBookingChoices`

	queryRows := comment + `
SELECT
	"bookings"."id",
	"bookings"."totalPriceIDR",
	COALESCE("rooms"."roomName", ''),
	"bookings"."dateStart",
	COALESCE("tenants"."tenantName", '')
FROM "bookings"
LEFT JOIN "tenants" ON "bookings"."tenantId" = "tenants"."id"
LEFT JOIN "rooms" ON "bookings"."roomId" = "rooms"."id"
ORDER BY "bookings"."id" ASC`

	out := make(map[uint64]string)
	b.Adapter.QuerySql(queryRows, func(row []any) {
		fmt.Println(`Row:`, row)
		if len(row) == 5 {
			id := `#` + X.ToS(row[0])
			totalPriceIdr := formatCurrency(X.ToI(row[1]), `IDR`)
			roomName := X.ToS(row[2])
			dateStart := X.ToS(row[3])
			tenantName := X.ToS(row[4])

			formattedBooking := fmt.Sprintf(
				"%s - %s - %s - %s - %s",
				id, totalPriceIdr, roomName, dateStart, tenantName,
			)
			out[X.ToU(row[0])] = formattedBooking
		}
	})

	return out
}

func formatCurrency(value int64, currency string) string {
	str := strconv.FormatInt(value, 10)

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	var result []rune
	for i, r := range runes {
		if i > 0 && i%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, r)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return currency + ` ` + string(result)
}

func (s *Stocks) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Stocks) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + s.SqlTableName() + whereAndSql + `
LIMIT 1`
	s.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + s.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	s.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (r *Rooms) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Rooms) FindByPagination`

	validFields := RoomsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + r.SqlTableName() + whereAndSql + `
LIMIT 1`
	r.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + r.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	r.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})
	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Bookings) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Bookings) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Bookings) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Buildings) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Buildings) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Buildings) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Facilities) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Facilities) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Facilities) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Locations) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Locations) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Locations) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Payments) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Payments) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Payments) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Rooms) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Rooms) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Rooms) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Stocks) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Stocks) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Stocks) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (r *Rooms) FindRoomChoices() map[uint64]string {
	const comment = `-- Rooms) FindRoomChoices`

	queryRows := comment + `
SELECT ` + r.SqlId() + `, ` + r.SqlRoomName() + ` FROM ` + r.SqlTableName() + `
ORDER BY ` + r.SqlRoomName() + ` ASC`

	out := make(map[uint64]string)
	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (r *Rooms) FindRoomNames() (out []string) {
	const comment = `-- Rooms) FindRoomNames`

	queryRows := comment + `
SELECT ` + r.SqlRoomName() + ` FROM ` + r.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + r.SqlRoomName() + ` ASC`

	extTenantsMap := make(map[string]bool)

	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 1 {
			roomName := X.ToS(row[0])
			extTenantsMap[roomName] = false
		}
	})

	queryExtraTenants := `SELECT
		"rooms"."roomName",
		"bookings"."extraTenants"
	FROM "bookings"
	LEFT JOIN "rooms" ON "bookings"."roomId" = "rooms"."id"
	WHERE "rooms"."currentTenantId" = "bookings"."tenantId"`

	roomExtraTenants := make(map[string]string)
	r.Adapter.QuerySql(queryExtraTenants, func(row []any) {
		if len(row) == 2 {
			roomName := X.ToS(row[0])
			roomExtraTenants[roomName] = X.ToS(row[1])
		}
	})

	for roomName, extTenants := range roomExtraTenants {
		if extTenants != `[]` {
			extTenantsMap[roomName] = true
		}
	}

	for roomName, hasExtraTenants := range extTenantsMap {
		toRoomName := roomName
		if hasExtraTenants {
			toRoomName += ` *`
		}
		out = append(out, toRoomName)
	}

	sort.Strings(out)

	return
}

const DateFormatYYYYMM = "2006-01"

func isValidYearMonth(yearMonth string) bool {
	_, err := time.Parse(DateFormatYYYYMM, yearMonth)
	return err == nil
}

type BookingDetail struct {
	Id           uint64 `json:"id"`
	RoomId       uint64 `json:"roomId"`
	RoomName     string `json:"roomName"`
	TenantId     uint64 `json:"tenantId"`
	TenantName   string `json:"tenantName"`
	DateStart    string `json:"dateStart"`
	DateEnd      string `json:"dateEnd"`
	AmountPaid   int64  `json:"amountPaid"`
	TotalPrice   int64  `json:"totalPrice"`
	DeletedAt    int64  `json:"deletedAt"`
	IsNearEnding bool   `json:"isNearEnding"`
	IsExtended   bool   `json:"isExtended"`
}

type RoomBooking struct {
	RoomName string          `json:"roomName"`
	Bookings []BookingDetail `json:"bookings"`
}

func getEndOfMonth(ym string) string {
	t, _ := time.Parse(DateFormatYYYYMM, ym)
	endOfMonth := t.AddDate(0, 1, 0).AddDate(0, 0, -1)
	return endOfMonth.Format(time.DateOnly)
}

func (b *Bookings) FindBookingsPerQuartal(monthStart, monthEnd string) (out []BookingDetail) {
	const comment = `-- Bookings) FindBookingsPerQuartal`

	if !(isValidYearMonth(monthStart) || isValidYearMonth(monthEnd)) {
		monthStart = time.Now().Format(DateFormatYYYYMM)
		monthEnd = time.Now().AddDate(0, 4, 0).Format(DateFormatYYYYMM)
	}

	monthEnd = getEndOfMonth(monthEnd)
	monthStart += "-01"

	queryRows := comment + `
SELECT 
	"bookings"."id" AS "id",
	"rooms"."id" AS "roomId",
  "rooms"."roomName",
	"bookings"."tenantId",
  COALESCE("tenants"."tenantName", '') AS "tenantName",
  COALESCE("bookings"."dateStart", '') AS "dateStart",
  COALESCE("bookings"."dateEnd", '') AS "dateEnd",
  COALESCE(SUM("payments"."paidIDR"), 0) AS "totalPaidIDR",
  "bookings"."totalPriceIDR",
	"bookings"."deletedAt"
FROM "bookings"
LEFT JOIN "tenants" ON "bookings"."tenantId" = "tenants"."id"
LEFT JOIN "rooms" ON "bookings"."roomId" = "rooms"."id"
LEFT JOIN "payments" ON "bookings"."id" = "payments"."bookingId"
WHERE
	(
    "bookings"."dateStart" BETWEEN ` + S.Z(monthStart) + ` AND ` + S.Z(monthEnd) + `
    OR
    "bookings"."dateEnd" BETWEEN ` + S.Z(monthStart) + ` AND ` + S.Z(monthEnd) + `
  )
GROUP BY "rooms"."roomName", "tenants"."tenantName", "bookings"."dateStart", "bookings"."dateEnd", "bookings"."totalPriceIDR"
ORDER BY "rooms"."roomName" ASC`

	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 10 {
			id := X.ToU(row[0])
			roomId := X.ToU(row[1])
			roomName := X.ToS(row[2])
			tenantId := X.ToU(row[3])
			tenantName := X.ToS(row[4])
			dateStart := X.ToS(row[5])
			dateEnd := X.ToS(row[6])
			totalPaidIdr := X.ToI(row[7])
			totalPriceIdr := X.ToI(row[8])
			deletedAt := X.ToI(row[9])

			out = append(out, BookingDetail{
				Id:         id,
				RoomId:     roomId,
				RoomName:   roomName,
				TenantId:   tenantId,
				TenantName: tenantName,
				DateStart:  dateStart,
				DateEnd:    dateEnd,
				AmountPaid: totalPaidIdr,
				TotalPrice: totalPriceIdr,
				DeletedAt:  deletedAt,
			})
		}
	})

	out = updateFlags(out)

	return
}

func updateFlags(bookings []BookingDetail) []BookingDetail {
	now := time.Now()
	sevenDaysLater := now.AddDate(0, 0, 7)

	for i := range bookings {
		current := &bookings[i]
		if current.DeletedAt != 0 {
			continue
		}

		dateEnd, err := time.Parse(time.DateOnly, current.DateEnd)
		if err != nil {
			continue
		}

		// ============ IS EXTENDED ============
		for _, other := range bookings {
			if other.DeletedAt != 0 || other.Id == current.Id {
				continue
			}
			if other.RoomId == current.RoomId {
				otherStart, err := time.Parse(time.DateOnly, other.DateStart)
				if err != nil {
					continue
				}
				if otherStart.After(dateEnd) {
					current.IsExtended = true
					break
				}
			}
		}

		// ============ IS NEAR ENDING ============
		hasFutureBookingSameRoom := false
		for _, other := range bookings {
			if other.DeletedAt != 0 || other.Id == current.Id {
				continue
			}
			if other.RoomId == current.RoomId {
				otherStart, err := time.Parse(time.DateOnly, other.DateStart)
				if err != nil {
					continue
				}
				if otherStart.After(now) {
					hasFutureBookingSameRoom = true
					break
				}
			}
		}

		if !hasFutureBookingSameRoom {
			if dateEnd.Before(now) {
				current.IsNearEnding = true // Sudah habis dan tidak diperpanjang
			} else if dateEnd.Before(sevenDaysLater.Add(24 * time.Hour)) {
				current.IsNearEnding = true // Akan habis <= 7 hari
			}
		}
	}

	return bookings
}

func parseMonth(monthStr string) (time.Time, error) {
	return time.Parse("2006-01", monthStr)
}

func monthsInRange(start, end time.Time) []time.Time {
	var months []time.Time
	for t := start; !t.After(end); t = t.AddDate(0, 1, 0) {
		months = append(months, t)
	}
	return months
}

func GroupBookingsToQuarter(bookings []BookingDetail, monthStartStr, monthEndStr string) ([]RoomBooking, error) {
	monthStart, err := parseMonth(monthStartStr)
	if err != nil {
		return nil, err
	}
	monthEnd, err := parseMonth(monthEndStr)
	if err != nil {
		return nil, err
	}

	monthSlots := monthsInRange(monthStart, monthEnd)
	roomMap := make(map[string][]BookingDetail)

	for _, b := range bookings {
		startDate, err1 := time.Parse(time.DateOnly, b.DateStart)
		endDate, err2 := time.Parse(time.DateOnly, b.DateEnd)
		if err1 != nil || err2 != nil {
			continue
		}

		for i, m := range monthSlots {
			monthStart := m
			monthEnd := m.AddDate(0, 1, 0).Add(-time.Nanosecond) // end of month
			if (startDate.Before(monthEnd) || startDate.Equal(monthEnd)) &&
				(endDate.After(monthStart) || endDate.Equal(monthStart)) {
				if _, exists := roomMap[b.RoomName]; !exists {
					roomMap[b.RoomName] = make([]BookingDetail, len(monthSlots))
				}
				roomMap[b.RoomName][i] = b
			}
		}
	}

	// Build result
	var result []RoomBooking
	for room, bookings := range roomMap {
		result = append(result, RoomBooking{
			RoomName: room,
			Bookings: bookings,
		})
	}
	return result, nil
}

func (p *Payments) FindByBookingId(bookingId uint64) (rows []Payments) {
	const comment = `-- Payments) FindByBookingId`

	query := `SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() + `
WHERE ` + p.SqlBookingId() + ` = ` + I.UToS(bookingId) + `
AND ` + p.SqlDeletedAt() + ` = 0`

	p.Adapter.QuerySql(query, func(row []any) {
		p.FromArray(row)
		rows = append(rows, *p)
	})

	return
}
