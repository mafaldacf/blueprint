

struct CatalogueService_AddSock_Request {
	1: Sock sock,
}

struct CatalogueService_AddSock_Response {
	1: string ret0,
}

struct CatalogueService_AddTags_Request {
	1: list<string> tags,
}

struct CatalogueService_AddTags_Response {
}

struct CatalogueService_Count_Request {
	1: list<string> tags,
}

struct CatalogueService_Count_Response {
	1: i32 ret0,
}

struct CatalogueService_DeleteSock_Request {
	1: string id,
}

struct CatalogueService_DeleteSock_Response {
}

struct CatalogueService_Get_Request {
	1: string id,
}

struct CatalogueService_Get_Response {
	1: Sock ret0,
}

struct CatalogueService_List_Request {
	1: list<string> tags,
	2: string order,
	3: i32 pageNum,
	4: i32 pageSize,
}

struct CatalogueService_List_Response {
	1: list<Sock> ret0,
}

struct CatalogueService_Tags_Request {
}

struct CatalogueService_Tags_Response {
	1: list<string> ret0,
}

struct Sock {
	1: string ID,
	2: string Name,
	3: string Description,
	4: list<string> ImageURL,
	5: string ImageURL_1,
	6: string ImageURL_2,
	7: double Price,
	8: i32 Quantity,
	9: list<string> Tags,
	10: string TagString,
}



service CatalogueService {
	CatalogueService_AddSock_Response AddSock (1:CatalogueService_AddSock_Request req),
	CatalogueService_AddTags_Response AddTags (1:CatalogueService_AddTags_Request req),
	CatalogueService_Count_Response Count (1:CatalogueService_Count_Request req),
	CatalogueService_DeleteSock_Response DeleteSock (1:CatalogueService_DeleteSock_Request req),
	CatalogueService_Get_Response Get (1:CatalogueService_Get_Request req),
	CatalogueService_List_Response List (1:CatalogueService_List_Request req),
	CatalogueService_Tags_Response Tags (1:CatalogueService_Tags_Request req),
}

