

struct MovieInfo {
	1: string MovieID,
	2: string Title,
	3: string Casts,
}

struct MovieInfoService_WriteMovieInfo_Request {
	1: i64 reqID,
	2: string movieID,
	3: string title,
	4: string casts,
}

struct MovieInfoService_WriteMovieInfo_Response {
	1: MovieInfo ret0,
}



service MovieInfoService {
	MovieInfoService_WriteMovieInfo_Response WriteMovieInfo (1:MovieInfoService_WriteMovieInfo_Request req),
}

