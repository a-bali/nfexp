<script>
  import { apiUrl, postBody } from "$lib";
  import Report from "$lib/Report.svelte";
  import Query from "$lib/Query.svelte";
  let result;
  let format;

  async function runCommand(params) {
    localStorage.query = JSON.stringify(params);
    format = params.format;
    result = fetch(apiUrl("cmd"), postBody(params)).then((v) => v.text());
  }
</script>

<h1>Query</h1>

<Query
  params={localStorage.query ? JSON.parse(localStorage.query) : {}}
  onClick={runCommand}
/>

{#if result}
  {#await result}
    ...
  {:then report}
    <Report {report} {format} />
  {/await}
{/if}
