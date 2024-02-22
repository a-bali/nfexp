<script>
  let command = "";
  let result;
  async function runCommand() {
    const res = await fetch("http://localhost:8080/api/query", {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: new URLSearchParams({ command: command }),
    });
    result = await res.text();
  }
</script>

<h2>Query</h2>

<div>
  <input type="text" bind:value={command} size="100" />
  <button on:click={runCommand}>run</button>
</div>

{#if result}
  <div><pre>{result}</pre></div>
{/if}
