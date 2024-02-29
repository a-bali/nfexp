<script>
  import { DataHandler, Th, ThFilter } from "@vincjo/datatables";
  import { csv2map, formatBytes, apiUrl } from "$lib";

  export let report;

  function iplookups(d) {
    d.forEach((row, i) => {
      d[i]["host"] = fetch(
        apiUrl("dns?") + new URLSearchParams({ ip: row["val"] }),
      ).then((r) => r.text());
    });
    return d;
  }
  const handler = new DataHandler(iplookups(csv2map(report)));
  const rows = handler.getRows();
</script>

<table>
  <thead>
    <tr>
      <Th {handler} orderBy="ts">time start</Th>
      <Th {handler} orderBy="val">ip</Th>
      <Th {handler} orderBy={(row) => parseInt(row.byt)}>bytes</Th>
    </tr>
    <tr>
      <ThFilter {handler} filterBy="ts">time start</ThFilter>
      <ThFilter {handler} filterBy="val">ip</ThFilter>
    </tr>
  </thead>
  <tbody>
    {#each $rows as row}
      <tr>
        <td>{row.ts}</td>
        <td>{#await row.host}{row.val}{:then h}{h} ({row.val}){/await}</td>
        <td>{formatBytes(row.byt)}</td>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  table {
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    width: 100%;
  }
  td {
    padding: 4px 20px;
    border-bottom: 1px solid #eee;
  }
</style>
