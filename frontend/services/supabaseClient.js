import { createClient } from "@supabase/supabase-js";
import "dotenv/config";

export const supabase = createClient(
  process.env.REACT_APP_SUPABASE_URL,
  process.env.REACT_APP_ANON_KEY,
);
