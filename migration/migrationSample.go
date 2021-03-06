package migration

import ( 
	"../models"
)

func UpAddressTable() {
	addressModel := models.NewAddressModel()

	addressModel.CreateTable(models.Address{})
	addressModel.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(models.Address{})
}

func UpUserTable() {
	userModel := models.NewUserModel()

	userModel.CreateTable(userModel.User)
	userModel.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(userModel.User)
}


func DownAddressTable() {
	addressModel := models.NewAddressModel()

	address := models.Address{}
	addressModel.DropTable(address)
}

func DownUserTable() {
	userModel := models.NewUserModel()
	userModel.DropTable(userModel.User)
}

