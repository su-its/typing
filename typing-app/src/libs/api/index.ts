import createClient from "openapi-fetch";
import { paths } from "./v0";

// ブラウザから使うときはこっち
export const client = createClient<paths>({ baseUrl: process.env.NEXT_PUBLIC_API_URL });
