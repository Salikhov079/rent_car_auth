package handler

import (
	"github.com/Salikhov079/rent_car/api/token"

	pb "github.com/Salikhov079/rent_car/genprotos"
	"github.com/Salikhov079/rent_car/models"

	"github.com/gin-gonic/gin"
)

// CreateUser 		handles the creation of a new user
// @Summary 		Create User
// @Description 	Create page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     models.Users   true   "Create"
// @Success 		200   {string}   pb.User   "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/user/registr [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	arr := &models.Users{}
	err := ctx.BindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	user := &pb.User{UserName: arr.UserName, PhoneNumber: arr.PhoneNumber, Password: arr.Password, Email: arr.Email, Role: arr.Role}

	_, err = h.User.Create(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	t := token.GenereteJWTToken(user)
	ctx.JSON(200, t)
}

// UpdateUser 		handles the creation of a new user
// @Summary			Update User
// @Description 	Update page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id 		path   string     true   "User ID"
// @Param   		Update  body   models.Users     true   "Update"
// @Success 		200   {string} string    "Update Successful"
// @Failure 		401   {string} string    "Error while created"
// @Router 			/user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	arr := models.Users{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	user := &pb.User{Id: ctx.Param("id"), UserName: arr.UserName, PhoneNumber: arr.PhoneNumber, Password: arr.Password, Email: arr.Email, Role: arr.Role}
	
	_, err = h.User.Update(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteUser 		handles the creation of a new User
// @Summary			Delete User
// @Description 	Delete page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id   path     string   true   "User ID"
// @Success 		200 {string}  string   "Delete Successful"
// @Failure 		401 {string}  string   "Error while Deleted"
// @Router 			/user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.User.Delete(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllUser 		handles the creation of a new User
// @Summary 		GetAll User
// @Description 	GetAll page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			query  query  models.Filter true    "Query parameter"
// @Success 		200 {object}  pb.GetAllUsers  "GetAll Successful"
// @Failure 		401 {string}  string  		  "Error while GetAll"
// @Router 			/user/getall  [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	user := &pb.User{}
	user.Email = ctx.Query("email")
	user.UserName = ctx.Query("user_name")

	res, err := h.User.GetAll(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		GetById User
// @Description 	GetById page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id   path      string   true    "User ID"
// @Success 		200 {object}   pb.User  "GetById Successful"
// @Failure 		401 {string}   string   "Error while GetByIdd"
// @Router 			/user/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.User.GetById(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body  models.Login    true     "Create"
// @Success 		200 {object}  pb.User  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	arr := &models.Login{}
	err := ctx.ShouldBindJSON(arr)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	user := &pb.User{UserName: arr.UserName}

	res, err := h.User.Login(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	t := token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}
