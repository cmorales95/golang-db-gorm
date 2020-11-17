# GORM 

#### Create
```
	obs := "testing con go"
	product1 := models.Product{
		Name: "Curso de Go",
		Price: 120,
	}
	product2 := models.Product{
		Name: "Curso de Testing con Go",
		Observations: &obs, 
		Price: 150,
	}
	product3 := models.Product{
		Name: "Curso de Python",
		Price: 92,
	}
	
	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
```

#### Get All
```
        products := make([]models.Product, 0)
	storage.DB().Find(&products)
	for _, product := range products {
		fmt.Printf("%d - %s\n",product.ID, product.Name)
	}
}
```

#### Get by ID
```
    product := models.Product{}
	storage.DB().First(&product, 3)
	fmt.Println(product)
```

#### Update
```
    product := models.Product{}
	product.ID = 3
	storage.DB().Model(&product).Updates(
		models.Product{Name: "Curso de Java", Price: 80},
	)
```
 
 #### Delete (Soft)
```
    product := models.Product{}
	product.ID = 3
	storage.DB().Delete(&product)
```

### Delete (Permanently)
```
    product := models.Product{}
	product.ID = 3
	storage.DB().Unscoped().Delete(&product)
```