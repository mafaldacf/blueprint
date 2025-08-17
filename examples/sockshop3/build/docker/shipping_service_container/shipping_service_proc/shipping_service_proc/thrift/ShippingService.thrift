

struct Shipment {
	1: string ID,
	2: string Name,
	3: string Status,
}

struct ShippingService_GetShipment_Request {
	1: string id,
}

struct ShippingService_GetShipment_Response {
	1: Shipment ret0,
}

struct ShippingService_PostShipping_Request {
	1: Shipment shipment,
}

struct ShippingService_PostShipping_Response {
	1: Shipment ret0,
}

struct ShippingService_UpdateStatus_Request {
	1: string id,
	2: string status,
}

struct ShippingService_UpdateStatus_Response {
}



service ShippingService {
	ShippingService_GetShipment_Response GetShipment (1:ShippingService_GetShipment_Request req),
	ShippingService_PostShipping_Response PostShipping (1:ShippingService_PostShipping_Request req),
	ShippingService_UpdateStatus_Response UpdateStatus (1:ShippingService_UpdateStatus_Request req),
}

