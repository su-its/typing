import createClient from "openapi-fetch";
import { paths } from "./v1";

export const client = createClient<paths>({ baseUrl: process.env.NEXT_PUBLIC_API_URL });
