package handlers

import (
	"meetspace_backend/client/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"

	"github.com/gin-gonic/gin"
)

// RegisterClientHandler godoc
//	@Summary		UserLogin User account
//	@Description	UserLogin User account
//	@Tags			Client
//	@Produce		json
// @Param user body types.ClientCreateData true "User login details"
//	@Router			/v1/clients [post]
func RegisterClientHandler(c *gin.Context){

	var reqData types.ClientCreateData
	if err := utils.BindJsonData(c, &reqData); err != nil {
		utils.HandleError(c, err)
		return
	}
	reqData.Password = "Indal@123"
	client, err := config.ClientService.CreateClient(reqData)
	if err != nil {  
		utils.HandleError(c, err)
		return 
	}
	resp := utils.SuccessResponse("success", client)
	
	c.JSON(resp.StatusCode, resp)
	return
}

// GetClientById godoc
//	@Summary		GetClientById User account
//	@Description	GetClientById User account
//	@Tags			Client
//	@Produce		json
// @Param id path int true "Client ID"
//	@Router			/v1/clients/{id} [get]
func GetClientById(c *gin.Context){
	clientId := c.Param("clientId")

	client, _ := config.ClientService.GetClientById(clientId)

	resp := utils.SuccessResponse("success", client)
	
	c.JSON(resp.StatusCode, resp)
	return
}

// GetAllClients godoc
//	@Summary		GetAllClients User account
//	@Description	GetAllClients User account
//	@Tags			Client
//	@Produce		json
// @Param company_name query string false "Client's company name"
// @Param user body types.ClientCreateData true "GetAllClients login details"
//	@Router			/v1/clients [get]
func GetAllClients(c *gin.Context){
	_, isUser:= utils.GetUserFromContext(c)
	if !isUser{
		c.JSON(400, gin.H{"message": "user not found"})
		return
	}

	companyName := c.Query("company_name")

	clients, _ := config.ClientService.GetAllClients(companyName)

	resp := utils.SuccessResponse("success", clients)
	c.JSON(resp.StatusCode, resp)
	return
}
