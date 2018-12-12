package server

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/controller"
	"github.com/hundredwz/GBlog/dao"
	"github.com/hundredwz/GBlog/service"
	"github.com/hundredwz/GBlog/util"
	"html/template"
	"path/filepath"
)

var (
	us              service.UserService
	ac              controller.ArticleController
	mc              controller.MetaController
	pc              controller.PageController
	cc              controller.CommentController
	ic              controller.InstallController
	sc              controller.SettingController
	uc              controller.UserController
	adminController controller.AdminController
)

func initController() {
	db := &dao.DataBase{}
	db.Init()
	contentService := service.ContentService{DB: db}
	metaService := service.MetaService{DB: db}
	commentService := service.CommentService{DB: db}
	us := service.UserService{DB: db}
	installService := service.InstallService{DB: db}
	ac = controller.ArticleController{ContentService: contentService}
	mc = controller.MetaController{MetaService: metaService, ContentService: contentService}
	pc = controller.PageController{ContentService: contentService, MetaService: metaService, CommentService: commentService}
	cc = controller.CommentController{CommentService: commentService}
	ic = controller.InstallController{InstallService: installService, UserService: us}
	sc = controller.SettingController{UserService: us}
	uc = controller.UserController{UserService: us}
	adminController = controller.AdminController{ContentService: contentService, MetaService: metaService, CommentService: commentService, UserService: us}
}

func InitServer(server *gin.Engine) {
	config.InitConfig()
	initController()

	server.Static("/admin/static", "web/admin/static")
	server.Static("/static", "web/user/static")

	server.GET("/install", ic.InstallPage)
	server.POST("/api/install", ic.Install)
	server.POST("/api/install/dbtest", ic.DBConnection)

	server.Use(Installed())

	server.HTMLRender = addTemplate("web")

	server.GET("/", pc.Index)
	server.GET("/index", pc.Index)
	server.GET("/article", pc.Article)
	server.GET("/category/:slug", pc.Category)
	server.GET("/tag/:slug", pc.Tag)
	server.GET("/page", pc.Page)

	//admin
	server.GET("/admin/login", adminController.Login)
	admin := server.Group("/admin", Authorized())
	{
		admin.GET("", adminController.Index)
		admin.GET("/index", adminController.Index)
		admin.GET("/article/edit", adminController.ArticleEdit)
		admin.GET("/article/list", adminController.ArticleList)
		admin.GET("/page/edit", adminController.PageEdit)
		admin.GET("/page/list", adminController.PageList)
		admin.GET("/category/edit", adminController.CategoryEdit)
		admin.GET("/category/list", adminController.CategoryList)
		admin.GET("/tag/edit", adminController.TagEdit)
		admin.GET("/tag/list", adminController.TagList)
		admin.GET("/comment/list", adminController.CommentList)
		admin.GET("/setting/blog", adminController.BlogSetting)
		admin.GET("/setting/user", adminController.UserSetting)
	}

	api := server.Group("/api")
	{
		api.GET("/articles", ac.GetArticles)
		api.GET("/article", ac.GetArticle)
		api.POST("/article/edit", Authorized(), ac.EditArticle)
		api.POST("/article/status", Authorized(), ac.UpdateContentStatus)
		api.GET("/article/delete", Authorized(), ac.DelContent)
		api.POST("/page/edit", Authorized(), ac.EditPage)
		api.GET("/page/delete", Authorized(), ac.DelContent)
		api.POST("/page/status", Authorized(), ac.UpdateContentStatus)

		api.POST("/meta/status", Authorized(), mc.UpdateMetaStatus)
		api.GET("/categoryList", mc.CategoryList)
		api.GET("/category", mc.CategoryInfo)
		api.POST("/category/edit", Authorized(), mc.EditCategory)
		api.GET("/category/delete", Authorized(), mc.DelMeta)

		api.GET("/tag", mc.TagInfo)
		api.POST("/tag/edit", mc.EditTag)
		api.GET("/tagList", mc.TagList)
		api.GET("/tag/delete", Authorized(), mc.DelMeta)

		api.GET("/comment", cc.GetComment)
		api.POST("/comment/edit", Authorized(), cc.EditComment)
		api.POST("/comment/status", Authorized(), cc.UpdateCommentStatus)
		api.GET("/comment/delete", Authorized(), cc.DelComment)

		api.POST("/user/login", uc.Login)
		api.GET("/user/logout", uc.Logout)
		api.POST("/setting/blog", Authorized(), sc.BlogSetting)
		api.POST("/setting/user", Authorized(), sc.UserSetting)

	}

}

func addTemplate(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("install.html", "web/admin/install/install.html")

	funcMap := template.FuncMap{
		"formatTime": util.FormatTime,
		"ifMetaIn":   util.IfMetaIn,
	}

	admin, err := filepath.Glob(templateDir + "/admin/views/*.html")
	if err != nil {
		return r
	}
	adminLayouts, _ := filepath.Glob(templateDir + "/admin/views/layouts/*.html")

	for _, adminHtml := range admin {
		layoutCopy := make([]string, len(adminLayouts))
		copy(layoutCopy, adminLayouts)
		files := append(layoutCopy, adminHtml)
		if filepath.Base(adminHtml) == "login.html" {
			continue
		}
		r.AddFromFilesFuncs("admin-"+filepath.Base(adminHtml), funcMap, files...)
	}
	r.AddFromFiles("admin-login.html", "web/admin/views/login.html")

	user, err := filepath.Glob(templateDir + "/user/views/*.html")
	if err != nil {
		return r
	}
	userLayouts, _ := filepath.Glob(templateDir + "/user/views/layouts/*.html")
	for _, userHtml := range user {
		layoutCopy := make([]string, len(userLayouts))
		copy(layoutCopy, userLayouts)
		files := append(layoutCopy, userHtml)
		r.AddFromFiles(filepath.Base(userHtml), files...)
	}
	return r
}
