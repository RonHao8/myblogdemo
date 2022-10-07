package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"myblogdemo/utils"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

//---------数据处理-----------
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

//-----------数据库操作---------------
//插入一篇文章
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article (title,tags,short,content,author,createtime) "+
		"values(?,?,?,?,?,?)", article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("page----->", page)
	return QueryArticleWithPage(page, num)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面咩有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d,%d", num*page, num)
	return QueryArticleWithCon(sql)
}

func QueryArticleWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	fmt.Println("queryarticlesql==", sql)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
		//fmt.Println("artList_model===", artList)
	}
	return artList, nil
}

//------翻页------
//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

//设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}

//----------查询文章-------------
func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id = " +
		strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{
		Id:         id,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     content,
		Createtime: createtime,
	}
	return art
}

//----------修改数据----------
func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id = ?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

//----------删除文章---------
func DeleteArticle(artID int) (int64, error) {
	i, err := DeleteArticleWithId(artID)
	SetArticleRowsNum()
	return i, err
}

func DeleteArticleWithId(artID int) (int64, error) {
	return utils.ModifyDB("delete from article where id = ?", artID)
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}
	return paramList
}

//--------------按照标签查询--------------
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := " where tags like '%& " + tag + "%&'"
	sql += "or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	return QueryArticleWithCon(sql)
}
