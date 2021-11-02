package repository

import "golang-unit-test/entity"

/*	jgn langsung bikin function mengambil property ke database
	Bikin function kontrak, yg return value-nya bisa nge-check,
	apakah ada/tidak ada property yg dipanggil di dalam  database
*/
type CategoryRepository interface {
	FindById(id string) *entity.Category
}
