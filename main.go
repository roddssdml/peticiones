package main

import ("github.com/gin-gonic/gin"
		"fmt")

func main(){

	router:= gin.Default()

	//router
	router.GET("/ping", func(c *gin.Context){

	//request

	//process

	//response

		c.String(200, "pong")
	})



	

	//person:= Persona{"Andrea", "Rivas"}
	//jsonData, err:= json.Marshal(person)

	/*if err!=nil{
		panic(err)
	}*/

	//fmt.Println("verrrr", string(jsonData))
		
	router.POST("/saludo", func(c *gin.Context){

		var body Persona
		err:= c.BindJSON(&body)
			if err!= nil{
				panic(err)
			}
			
		msg:= fmt.Sprintf("Hola %v %v", body.Nombre, body.Apellido)
		c.JSON(200, gin.H{"message": msg})

	})

	err:= router.Run(":8080")

	if err!= nil{
		panic(err)
	}


}


type Persona struct{
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
}



