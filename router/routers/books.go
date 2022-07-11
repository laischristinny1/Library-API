package routers

import(
	"net/http"
    "Library-API/controllers"
) 
var BooksRouters = []Router{
	{
		URI:                "/books",
		Method:             http.MethodPost,
		Function:             controllers.CreateBook,
	
	},
	{
		URI:                "/books/{bookID}",
		Method:             http.MethodGet,
		Function:             controllers.SearchBook,
	
	},
	{
		URI:                "/books",
		Method:             http.MethodGet,
		Function:             controllers.GetAllBooks,
	
	},
	{
		URI:                "/books/checkout/{bookID}",
		Method:             http.MethodPatch,
		Function:             controllers.CheckoutBook,
		
	},
	{
		URI:                "/books/return/{bookID}",
		Method:             http.MethodPatch,
		Function:             controllers.ReturnBook,
		
	},
}