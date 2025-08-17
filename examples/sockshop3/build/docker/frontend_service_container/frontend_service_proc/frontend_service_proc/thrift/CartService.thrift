

struct CartService_AddItem_Request {
	1: string customerID,
	2: Item item,
}

struct CartService_AddItem_Response {
	1: Item ret0,
}

struct CartService_DeleteCart_Request {
	1: string customerID,
}

struct CartService_DeleteCart_Response {
}

struct CartService_GetCart_Request {
	1: string customerID,
}

struct CartService_GetCart_Response {
	1: list<Item> ret0,
}

struct CartService_GetItem_Request {
	1: string customerID,
	2: string itemID,
}

struct CartService_GetItem_Response {
	1: Item ret0,
}

struct CartService_MergeCarts_Request {
	1: string customerID,
	2: string sessionID,
}

struct CartService_MergeCarts_Response {
}

struct CartService_RemoveItem_Request {
	1: string customerID,
	2: string itemID,
}

struct CartService_RemoveItem_Response {
}

struct CartService_UpdateItem_Request {
	1: string customerID,
	2: Item item,
}

struct CartService_UpdateItem_Response {
}

struct Item {
	1: string ID,
	2: i32 Quantity,
	3: double UnitPrice,
}



service CartService {
	CartService_AddItem_Response AddItem (1:CartService_AddItem_Request req),
	CartService_DeleteCart_Response DeleteCart (1:CartService_DeleteCart_Request req),
	CartService_GetCart_Response GetCart (1:CartService_GetCart_Request req),
	CartService_GetItem_Response GetItem (1:CartService_GetItem_Request req),
	CartService_MergeCarts_Response MergeCarts (1:CartService_MergeCarts_Request req),
	CartService_RemoveItem_Response RemoveItem (1:CartService_RemoveItem_Request req),
	CartService_UpdateItem_Response UpdateItem (1:CartService_UpdateItem_Request req),
}

