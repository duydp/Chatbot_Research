package main

import (
	"soap"
	"fmt"
	"encoding/xml"
	"BVGI/config"
)

func main()  {
	result, _ :=  soap.GetInfoByBIEN_SOXE("30Z0043")
	fmt.Print("Result: " + result.SearchCommonInfoResult)

	// Read xml
	byteValue := []byte(result.SearchCommonInfoResult)

	var rowSets config.RowSet
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &rowSets)
	//var raw map[string]string

	if (rowSets.Rows[0].DATA == "No DataFound" || rowSets.Rows[0].DATA == "False Authentication") {
		fmt.Print("Result: " + result.SearchCommonInfoResult)
	}

	var msgText string
	for i := 0; i < len(rowSets.Rows); i++ {

		msgText = `Thông tin hồ sơ ContactID: ` + rowSets.Rows[i].ContactID +
			`- Chu hop dong:`+ rowSets.Rows[i].ChuHopDong+
			`-So Hop dong: `+ rowSets.Rows[i].SoHopDong+
			`- Bien So Xe: ` + rowSets.Rows[i].BienSoXe+
			`- Trang thai hop dong: ` + rowSets.Rows[i].TrangThaiHopDong+
			`- Ngay Hieu Luc: ` + rowSets.Rows[i].NgayHieuLuc+
			`- Ngay Het Hieu Luc: ` + rowSets.Rows[i].NgayHetHieuLuc+
			`- Ten San Pham: ` + rowSets.Rows[i].TenSanPham+
			`- Ma DVKD: ` + rowSets.Rows[i].MaDVKD+
			`- Don vi kinh doanh: ` + rowSets.Rows[i].DONVIKINHDOANH+
			`- Chi nhanh don vi kinh doanh: ` + rowSets.Rows[i].CHINHANHDONVIKINHDOANH+
			`- PHone: ` + rowSets.Rows[i].PHONE+
			`- So Khung: ` + rowSets.Rows[i].SOKHUNG+
			`- So May: ` + rowSets.Rows[i].SOMAY+
			`- Nhom Xe: ` + rowSets.Rows[i].CAR_GROUP+
			`- Chung Loai: ` + rowSets.Rows[i].MAKE+
			`- Model: ` + rowSets.Rows[i].MODEL+``

	}
	fmt.Print(msgText)
}


