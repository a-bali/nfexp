export async function load({ fetch, params }) {
  const base = import.meta.env.VITE_API_BASEURL || "";
  const res = fetch(base + "/api/overview").then((r) => r.text());
  return { overview: { promise: res } };
}
