import { apiUrl } from "$lib";

export async function load({ fetch, params }) {
  return {
    reports: {
      promise: fetch(apiUrl("reports")).then((r) => r.json()),
    },
  };
}
