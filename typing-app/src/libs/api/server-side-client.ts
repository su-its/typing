import "server-only";

import createClient from "openapi-fetch";
import { paths } from "./v0";

// サーバーで使うときはこっち
export const client = createClient<paths>({ baseUrl: process.env.API_URL });
