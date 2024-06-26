package handler

import (
	"github.com/Salikhov079/rent_car/api/token"

	pb "github.com/Salikhov079/rent_car/genprotos"

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
	arr := pb.User{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.User.Create(ctx, &arr)
	if err != nil {
		panic(err)
	}
	t := token.GenereteJWTToken(&arr)
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
	arr := pb.User{}
	arr.Id = ctx.Param("id")
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.User.Update(ctx, &arr)
	if err != nil {
		panic(err)
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
		panic(err)
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
// @Param 			query  query  models.UsersFilter true    "Query parameter"
// @Success 		200 {object}  pb.GetAllUsers  "GetAll Successful"
// @Failure 		401 {string}  string  		  "Error while GetAll"
// @Router 			/user/getall  [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	user := &pb.User{}
	user.Email = ctx.Query("email")
	user.UserName = ctx.Query("user_name")

	res, err := h.User.GetAll(ctx, user)
	if err != nil {
		panic(err)
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
		panic(err)
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
// @Param   		Create  body  models.Users    true     "Create"
// @Success 		200 {object}  pb.User  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := &pb.User{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		panic(err)
	}
	_, err = h.User.Login(ctx, user)
	if err != nil {
		panic(err)
	}
	t := token.GenereteJWTToken(user)
	ctx.JSON(200, t)
}
