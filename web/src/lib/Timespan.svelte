<script>
  import { timespan } from "$lib/timespan.js";
  import { invalidateAll } from "$app/navigation";
  import { formatDate } from "$lib";
  import * as chrono from "chrono-node";

  let timespanSelect = "7d";

  let fromError = false;
  let toError = false;

  function updateTimespan() {
    $timespan.from = timespanSelect
      ? formatDate(chrono.parseDate(timespanSelect + " ago"))
      : "";
    $timespan.to = "";
    toError = fromError = false;
    invalidateAll();
  }

  function updateFromTo() {
    let from = chrono.parseDate($timespan.from);
    let to = chrono.parseDate($timespan.to);
    timespanSelect = "";
    if (!$timespan.from || from) {
      $timespan.from = from ? formatDate(from) : "";
      fromError = false;
    } else {
      fromError = true;
    }
    if (!$timespan.to || to) {
      $timespan.to = to ? formatDate(to) : "";
      toError = false;
    } else {
      toError = true;
    }
    if (!fromError && !toError) {
      invalidateAll();
    }
  }
</script>

<div>
  from:
  <input
    size="18"
    class={fromError ? "error" : ""}
    bind:value={$timespan.from}
    on:change={updateFromTo}
  />
  to:
  <input
    size="18"
    class={toError ? "error" : ""}
    bind:value={$timespan.to}
    on:change={updateFromTo}
  />
  <select bind:value={timespanSelect} on:change={updateTimespan}>
    <option value="">(select)</option>
    <option value="1h">last hour</option>
    <option value="12h">last 12h</option>
    <option value="24h">last 24h</option>
    <option value="3d">last 3d</option>
    <option value="7d">last 7d</option>
  </select>
</div>

<style>
  input.error {
    background-color: red;
  }
</style>
