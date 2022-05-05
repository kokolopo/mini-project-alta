package menu

type MenuService interface {
	CreateNewMenu(input InputNewMenu) (Menu, error)
}

type menuService struct {
	repository MenuRepository
}

func NewMenuService(repository MenuRepository) *menuService {
	return &menuService{repository}
}

func (s *menuService) CreateNewMenu(input InputNewMenu) (Menu, error) {
	var menu Menu

	//tangkap nilai dari inputan
	menu.Nama = input.Nama
	menu.Deskripsi = input.Deskripsi
	menu.Harga = input.Harga
	menu.Kategori = input.Kategori
	menu.UrlGambar = "Default.jpg"
	menu.ApakahTersedia = 1

	//save data yang sudah dimapping kedalam struct Mahasiswa
	newMenu, err := s.repository.Save(menu)
	if err != nil {
		return newMenu, err
	}

	return newMenu, nil
}
