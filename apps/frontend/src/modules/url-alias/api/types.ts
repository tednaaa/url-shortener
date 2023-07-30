export interface CreateUrlError {
	error: "Error";
	status: string;
}

export interface Url {
	url: string;
	alias: string;
}

export interface GetUrlsOkResponse {
	status: "OK";
	urls: Url[];
}
