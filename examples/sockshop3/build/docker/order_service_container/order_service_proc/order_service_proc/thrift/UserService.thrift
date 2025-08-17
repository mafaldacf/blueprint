

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

struct UserService_Delete_Request {
	1: string entity,
	2: string id,
}

struct UserService_Delete_Response {
}

struct UserService_GetAddresses_Request {
	1: string id,
}

struct UserService_GetAddresses_Response {
	1: list<Address> ret0,
}

struct UserService_GetCards_Request {
	1: string cardid,
}

struct UserService_GetCards_Response {
	1: list<Card> ret0,
}

struct UserService_GetUsers_Request {
	1: string id,
}

struct UserService_GetUsers_Response {
	1: list<User> ret0,
}

struct UserService_Login_Request {
	1: string username,
	2: string password,
}

struct UserService_Login_Response {
	1: User ret0,
}

struct UserService_PostAddress_Request {
	1: string userid,
	2: Address address,
}

struct UserService_PostAddress_Response {
	1: string ret0,
}

struct UserService_PostCard_Request {
	1: string userid,
	2: Card card,
}

struct UserService_PostCard_Response {
	1: string ret0,
}

struct UserService_PostUser_Request {
	1: User user,
}

struct UserService_PostUser_Response {
	1: string ret0,
}

struct UserService_Register_Request {
	1: string username,
	2: string password,
	3: string email,
	4: string first,
	5: string last,
}

struct UserService_Register_Response {
	1: string ret0,
}



service UserService {
	UserService_Delete_Response Delete (1:UserService_Delete_Request req),
	UserService_GetAddresses_Response GetAddresses (1:UserService_GetAddresses_Request req),
	UserService_GetCards_Response GetCards (1:UserService_GetCards_Request req),
	UserService_GetUsers_Response GetUsers (1:UserService_GetUsers_Request req),
	UserService_Login_Response Login (1:UserService_Login_Request req),
	UserService_PostAddress_Response PostAddress (1:UserService_PostAddress_Request req),
	UserService_PostCard_Response PostCard (1:UserService_PostCard_Request req),
	UserService_PostUser_Response PostUser (1:UserService_PostUser_Request req),
	UserService_Register_Response Register (1:UserService_Register_Request req),
}

