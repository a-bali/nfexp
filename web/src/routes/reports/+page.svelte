<script>
  import Report from "$lib/Report.svelte";
  import { fetchApi } from "$lib";
  export let data;
  let selected;
  let output;
  let format;

  $: data.reports.promise.then((r) => runCommand(r[selected]));

  async function runCommand(r) {
    if (!r) {
      output = "";
      return false;
    }
    format = r.format;
    output = fetchApi("cmd", "POST", {
      command: r.command,
      filter: r.filter,
      format: r.format,
    }).then((v) => v.text());
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

{#if output}
  {#await output}
    ...
  {:then output}
    <Report report={output} {format} />
  {/await}
{/if}
