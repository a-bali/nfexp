<script>
  import Table from "$lib/Table.svelte";
  export let report;
  export let format;
</script>

{#if format == "csv"}
  <Table {report} />
{:else}
  <div><pre>{report}</pre></div>
{/if}
