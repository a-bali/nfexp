import { writable } from "svelte/store";
import { formatDate } from "$lib";
import * as chrono from "chrono-node";

export const timespan = writable({ from: formatDate(chrono.parseDate("7d ago")), to: "" });
