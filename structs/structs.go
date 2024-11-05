package structs

type ResponseStruct struct{
	Ok bool
	Joke string
}

type ErrorResponse struct{
	Ok bool
	Message string
}