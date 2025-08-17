

struct Authorisation {
	1: bool Authorised,
	2: string Message,
}

struct PaymentService_Authorise_Request {
	1: double amount,
}

struct PaymentService_Authorise_Response {
	1: Authorisation ret0,
}



service PaymentService {
	PaymentService_Authorise_Response Authorise (1:PaymentService_Authorise_Request req),
}

