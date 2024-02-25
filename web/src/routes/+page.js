import { fetchApi } from "$lib";

export async function load({ fetch, params }) {
  return {
    overview: {
      promise: fetchApi("cmd", "POST", { command: "-g" }).then((r) => r.text()),
    },
  };
}
