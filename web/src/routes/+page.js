import { apiUrl, postBody } from "$lib";

export async function load({ fetch, params }) {
  return {
    overview: {
      promise: fetch(apiUrl("cmd"), postBody({ command: "-g" })).then((r) =>
        r.text(),
      ),
    },
  };
}
