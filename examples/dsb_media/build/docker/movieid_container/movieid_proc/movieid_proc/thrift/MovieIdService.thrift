

struct MovieId {
	1: string MovieID,
	2: string Title,
}

struct MovieIdService_RegisterMovieId_Request {
	1: i64 reqID,
	2: string movieID,
	3: string title,
}

struct MovieIdService_RegisterMovieId_Response {
	1: MovieId ret0,
}



service MovieIdService {
	MovieIdService_RegisterMovieId_Response RegisterMovieId (1:MovieIdService_RegisterMovieId_Request req),
}

