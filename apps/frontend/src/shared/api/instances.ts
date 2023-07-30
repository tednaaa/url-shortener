import axios from "axios";

import { API_URL, PASSWORD, USERNAME } from "../config";

export const apiInstance = axios.create({
	baseURL: API_URL,
	auth: {
		username: USERNAME,
		password: PASSWORD,
	},
});
