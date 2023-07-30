import { apiInstance } from "@/shared/api/instances";
import { GetUrlsOkResponse } from "./types";

export const routes = {
	url: "/url",
	urls: "/urls",
};

export const createUrl = (url: string, alias: string) => {
	return apiInstance.post(routes.url, { url, alias });
};

export const getUrls = () => {
	return apiInstance.get<GetUrlsOkResponse>(routes.urls);
};
