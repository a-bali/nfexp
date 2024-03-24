import { apiUrl, postBody } from "$lib";
import { timespan } from "$lib/timespan.js";
import { get } from 'svelte/store';

export async function load({ fetch, params }) {
  return {
    overview: {
      promise: fetch(
        apiUrl("cmd"),
        postBody({ command: "-g", from: get(timespan).from, to: get(timespan).to }),
      ).then((r) => r.text()),
    },
  };
}
