import { fetchApi } from "$lib";

export async function load({ fetch, params }) {
  return {
    reports: {
      promise: fetchApi("reports", "GET").then((r) => r.json()),
    },
  };
}
