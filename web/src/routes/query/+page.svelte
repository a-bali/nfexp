<script>
  import { fetchApi } from "$lib";
  import Report from "$lib/Report.svelte";
  let command = localStorage.getItem("command") || "";
  let filter = localStorage.getItem("filter") || "";
  let format = localStorage.getItem("format") || "";
  let result;
  async function runCommand() {
    result = fetchApi("cmd", "POST", {
      command: command,
      filter: filter,
      format: format,
    }).then((v) => v.text());
  }
  $: localStorage.setItem("command", command);
  $: localStorage.setItem("filter", filter);
  $: localStorage.setItem("format", format);
</script>

<h2>Query</h2>

<table>
  <tr>
    <td>command</td>
    <td><input type="text" bind:value={command} size="100" /></td>
  </tr>
  <tr>
    <td>filter</td>
    <td><input type="text" bind:value={filter} size="100" /></td>
  </tr>
  <tr>
    <td>format</td>
    <td><input type="format" bind:value={format} size="10" /></td>
  </tr>
</table>
<button on:click={runCommand}>run</button>

{#if result}
  {#await result}
    ...
  {:then report}
    <Report {report} {format} />
  {/await}
{/if}
