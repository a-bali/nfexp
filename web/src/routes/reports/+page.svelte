<script>
  import Report from "$lib/Report.svelte";
  import { apiUrl, postBody } from "$lib";
  import { invalidateAll } from "$app/navigation";
  export let data;
  let selected;
  let result;
  let format;

  $: data.reports.promise.then((r) => runCommand(r[selected]));

  async function runCommand(r) {
    if (!r) {
      result = "";
      return false;
    }
    format = r.format;
    result = fetch(apiUrl("cmd"), postBody(r)).then((v) => v.text());
  }

  async function delReport() {
    fetch(apiUrl("reports?uuid=" + selected), { method: "DELETE" }).then(() => {
      selected = undefined;
      invalidateAll();
    });
  }
</script>

<h1>Reports</h1>

<main>
  {#await data.reports.promise}
    (loading)
  {:then reports}
    <div>
      <select bind:value={selected}>
        <option>(select a report)</option>
        {#each Object.keys(reports) as report}
          <option value={report}>{reports[report].name}</option>
        {/each}
      </select>
    </div>
  {:catch error}
    Error: {error.message}
  {/await}
</main>

{#if result}
  {#await result}
    ...
  {:then report}
    <Report {report} {format} />
    <p />
    <button on:click={delReport}>(delete report)</button>
  {/await}
{/if}
