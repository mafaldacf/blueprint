

struct Address {
	1: string Street,
	2: string Number,
	3: string Country,
	4: string City,
	5: string PostCode,
	6: string ID,
}

struct Card {
	1: string LongNum,
	2: string Expires,
	3: string CCV,
	4: string ID,
}

struct Item {
	1: string ID,
	2: i32 Quantity,
	3: double UnitPrice,
}

struct Order {
	1: string ID,
	2: string CustomerID,
	3: User Customer,
	4: Address Address,
	5: Card Card,
	6: list<Item> Items,
	7: Shipment Shipment,
	8: string Date,
	9: double Total,
}

struct OrderService_GetOrder_Request {
	1: string orderID,
}

struct OrderService_GetOrder_Response {
	1: Order ret0,
}

struct OrderService_GetOrders_Request {
	1: string customerID,
}

struct OrderService_GetOrders_Response {
	1: list<Order> ret0,
}

struct OrderService_NewOrder_Request {
	1: string customerID,
	2: string addressID,
	3: string cardID,
	4: string cartID,
}

struct OrderService_NewOrder_Response {
	1: Order ret0,
}

struct Shipment {
	1: string ID,
	2: string Name,
	3: string Status,
}

struct User {
	1: string FirstName,
	2: string LastName,
	3: string Email,
	4: string Username,
	5: string Password,
	6: Address Addresses,
	7: Card Cards,
	8: string UserID,
	9: string Salt,
}



service OrderService {
	OrderService_GetOrder_Response GetOrder (1:OrderService_GetOrder_Request req),
	OrderService_GetOrders_Response GetOrders (1:OrderService_GetOrders_Request req),
	OrderService_NewOrder_Response NewOrder (1:OrderService_NewOrder_Request req),
}

