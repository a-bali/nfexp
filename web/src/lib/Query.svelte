<script>
  export let params;
  export let onClick;

  if (!params) params = { command: "", filter: "", format: "" };
  function runCommand() {
    onClick(params);
  }
</script>

<table>
  <tr>
    <td>command</td>
    <td><input type="text" bind:value={params.command} size="100" /></td>
  </tr>
  <tr>
    <td>filter</td>
    <td><input type="text" bind:value={params.filter} size="100" /></td>
  </tr>
  <tr>
    <td>format</td>
    <td><input type="format" bind:value={params.format} size="10" /></td>
  </tr>
</table>

<button on:click={runCommand}>run</button>
