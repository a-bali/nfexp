export async function load({ fetch, params }) {
    const response = await fetch("/api/overview");
    const value = await response.text();
    return { response: value };
}
